package ui

import (
	"image/color"
	"log"
	"sort"
	"strings"
	"sync"
	"time"
	"unicode"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/zMoooooritz/go-let-loose/pkg/hll"
	"github.com/zMoooooritz/go-let-loose/pkg/rconv2"
	"github.com/zMoooooritz/go-let-observer/assets"
	"github.com/zMoooooritz/go-let-observer/pkg/rcndata"
	"github.com/zMoooooritz/go-let-observer/pkg/util"
)

var fetchIntervalSteps = []time.Duration{
	100 * time.Millisecond,
	500 * time.Millisecond,
	time.Second,
	2 * time.Second,
	5 * time.Second,
}

type MapViewState struct {
	showServerInfo    bool
	showGrid          bool
	showPlayers       bool
	showPlayerInfo    bool
	showSpawns        bool
	showHelp          bool
	showScoreboard    bool
	initialDataLoaded bool
	selectedPlayerID  string
}

type CameraState struct {
	panX       float64
	panY       float64
	zoomLevel  float64
	isDragging bool
	lastMouseX int
	lastMouseY int
}

type FetchState struct {
	intervalIndex  int
	lastUpdateTime time.Time
	isFetching     bool
	fetchMutex     sync.Mutex
}

type RconData struct {
	currentMapName        string
	currentMapOrientation hll.Orientation
	serverName            string
	playerCurrCount       int
	playerMaxCount        int
	playerMap             map[string]hll.DetailedPlayerInfo
	playerList            []hll.DetailedPlayerInfo
	spawnTracker          *rcndata.SpawnTracker
}

type MapView struct {
	MapViewState
	CameraState
	FetchState
	RconData
	roleImages  map[string]*ebiten.Image
	spawnImages map[string]*ebiten.Image

	backgroundImage *ebiten.Image
	rcon            *rconv2.Rcon
	getDims         func() (int, int)
}

func NewMapView(rcon *rconv2.Rcon, getDims func() (int, int)) *MapView {
	return &MapView{
		MapViewState: MapViewState{
			showServerInfo: util.Config.UIStartupOptions.ShowServerInfoOverlay,
			showGrid:       util.Config.UIStartupOptions.ShowGridOverlay,
			showPlayers:    util.Config.UIStartupOptions.ShowPlayers,
			showPlayerInfo: util.Config.UIStartupOptions.ShowPlayerInfo,
			showSpawns:     util.Config.UIStartupOptions.ShowSpawns,
		},
		CameraState: CameraState{
			zoomLevel: MIN_ZOOM_LEVEL,
		},
		FetchState: FetchState{
			intervalIndex: INITIAL_FETCH_STEP,
		},
		RconData: RconData{
			spawnTracker: rcndata.NewSpawnTracker(),
		},
		roleImages:  util.LoadRoleImages(),
		spawnImages: util.LoadSpawnImages(),
		rcon:        rcon,
		getDims:     getDims,
	}
}

func (mv *MapView) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		return ebiten.Termination
	}

	if time.Since(mv.lastUpdateTime) >= fetchIntervalSteps[mv.intervalIndex] {
		mv.fetchMutex.Lock()
		if !mv.isFetching {
			mv.isFetching = true
			go func() {
				snapshot, err := rcndata.FetchRconDataSnapshot(mv.rcon)
				if err == nil {
					mv.processRconData(snapshot)
				}
				mv.fetchMutex.Lock()
				mv.isFetching = false
				mv.fetchMutex.Unlock()
			}()
		}
		mv.fetchMutex.Unlock()
	}

	mv.handleKeyboardInput()
	mv.handleMouseInput()

	return nil
}

func (mv *MapView) Draw(screen *ebiten.Image) {
	mv.drawBackground(screen)

	if !mv.initialDataLoaded {
		return
	}

	mv.drawMapView(screen)

	if mv.showHelp {
		mv.drawHelp(screen)
	} else if mv.showScoreboard {
		mv.drawScoreboard(screen)
	} else {
		if mv.showServerInfo {
			mv.drawHeader(screen)
		}
		if mv.showPlayerInfo {
			if player, ok := mv.playerMap[mv.selectedPlayerID]; ok {
				mv.drawPlayerOverlay(screen, player)
			}
		}
	}
}

func (mv *MapView) processRconData(snapshot *rcndata.RconDataSnapshot) {
	oldPlayerMap := mv.playerMap
	playerMap := map[string]hll.DetailedPlayerInfo{}
	for _, player := range snapshot.Players {
		playerMap[player.ID] = player
		if oldPlayer, ok := oldPlayerMap[player.ID]; ok {
			mv.spawnTracker.TrackPlayerPosition(oldPlayer, player)
		}
	}
	sort.Slice(snapshot.Players, func(i, j int) bool {
		return snapshot.Players[i].ID > snapshot.Players[j].ID
	})
	mv.spawnTracker.CleanExpiredSpawns()
	mv.playerMap = playerMap
	mv.playerList = snapshot.Players

	currMapName := assets.ToFileName(snapshot.CurrentMap.ID)
	mv.currentMapOrientation = snapshot.CurrentMap.Orientation
	if currMapName != mv.currentMapName {
		mv.spawnTracker.ResetSpawns()
		mv.currentMapName = currMapName
		img, err := util.LoadMapImage(currMapName)
		if err == nil {
			mv.backgroundImage = img
		} else {
			log.Println("Error loading background image:", err)
		}
	}

	mv.serverName = snapshot.SessionInfo.ServerName
	mv.playerCurrCount = snapshot.SessionInfo.PlayerCount
	mv.playerMaxCount = snapshot.SessionInfo.MaxPlayerCount

	mv.lastUpdateTime = snapshot.FetchTime
	mv.initialDataLoaded = true
}

func (mv *MapView) drawBackground(screen *ebiten.Image) {
	if mv.backgroundImage != nil {
		screenSize := screen.Bounds().Size()
		imageSize := mv.backgroundImage.Bounds().Size()
		scale := (float64(screenSize.X) / float64(imageSize.X)) * mv.zoomLevel

		options := &ebiten.DrawImageOptions{}
		options.GeoM.Scale(scale, scale)
		options.GeoM.Translate(mv.panX, mv.panY)
		screen.DrawImage(mv.backgroundImage, options)
	}
}

func (mv *MapView) drawMapView(screen *ebiten.Image) {
	screenSize := screen.Bounds().Size()

	if mv.showGrid {
		mv.drawGrid(screen, screenSize.X, screenSize.Y)
	}

	if mv.showSpawns {
		mv.drawSpawns(screen)
	}

	if mv.showPlayers {
		mv.drawPlayers(screen)
	}
}

func (mv *MapView) drawPlayers(screen *ebiten.Image) {
	var selectedPlayer *hll.DetailedPlayerInfo

	for _, player := range mv.playerList {
		if !player.IsSpawned() {
			continue
		}

		if mv.selectedPlayerID != "" && player.ID == mv.selectedPlayerID {
			selectedPlayer = &player
			continue
		}

		mv.drawPlayer(screen, player)
	}

	if selectedPlayer != nil {
		mv.drawPlayer(screen, *selectedPlayer)
	}
}

func (mv *MapView) drawPlayer(screen *ebiten.Image, player hll.DetailedPlayerInfo) {
	sizeX, sizeY := mv.getDims()
	x, y := util.TranslateCoords(sizeX, sizeY, player.Position)
	x = x*mv.zoomLevel + mv.panX
	y = y*mv.zoomLevel + mv.panY

	sizeModifier := PLAYER_SIZE_MODIFIER
	clr := CLR_ALLIES
	if player.Team == hll.TmAxis {
		clr = CLR_AXIS
	}
	if mv.selectedPlayerID == player.ID {
		sizeModifier = SELECTED_PLAYER_SIZE_MODIFIER
		clr = CLR_SELECTED
	}

	vector.DrawFilledCircle(screen, float32(x), float32(y), float32(util.IconCircleRadius(mv.zoomLevel, sizeModifier)), clr, false)

	roleImage, ok := mv.roleImages[strings.ToLower(string(player.Role))]
	if ok {
		targetSize := util.IconSize(mv.zoomLevel, sizeModifier)
		iconScale := targetSize / float64(roleImage.Bounds().Dx())

		options := &ebiten.DrawImageOptions{}
		options.GeoM.Scale(iconScale, iconScale)
		options.GeoM.Translate(x-targetSize/2, y-targetSize/2)
		screen.DrawImage(roleImage, options)
	}
}

func (mv *MapView) drawSpawns(screen *ebiten.Image) {
	sizeX, sizeY := mv.getDims()
	spawns := mv.spawnTracker.GetSpawns()
	for _, spawn := range spawns {
		if spawn.SpawnType == rcndata.SpawnTypeNone {
			continue
		}

		x, y := util.TranslateCoords(sizeX, sizeY, spawn.Position)
		x = x*mv.zoomLevel + mv.panX
		y = y*mv.zoomLevel + mv.panY

		clr := CLR_ALLIES_SPAWN
		if spawn.Team == hll.TmAxis {
			clr = CLR_AXIS_SPAWN
		}

		rectSize := int(2 * util.IconCircleRadius(mv.zoomLevel, SPAWN_SIZE_MODIFIER))
		util.DrawScaledRect(screen, int(x)-rectSize/2, int(y)-rectSize/2, rectSize, rectSize, clr)

		spawnImage, ok := mv.spawnImages[string(spawn.SpawnType)]
		if ok {
			targetSize := util.IconSize(mv.zoomLevel, SPAWN_SIZE_MODIFIER)
			iconScale := targetSize / float64(spawnImage.Bounds().Dx())

			options := &ebiten.DrawImageOptions{}
			options.GeoM.Scale(iconScale, iconScale)
			options.GeoM.Translate(x-targetSize/2, y-targetSize/2)
			screen.DrawImage(spawnImage, options)
		}
	}
}

func (mv *MapView) drawGrid(screen *ebiten.Image, width, height int) {
	sizeX, sizeY := mv.getDims()
	bgWidth := float64(sizeX) * mv.zoomLevel
	bgHeight := float64(sizeY) * mv.zoomLevel

	cellWidth := bgWidth / 5
	cellHeight := bgHeight / 5

	active := []int{1, 2, 1, 0, 0}

	gridColor := color.RGBA{100, 100, 100, 255}
	// fillColor := color.RGBA{50, 50, 50, 100}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			x := float64(i)*cellWidth + mv.panX
			y := float64(j)*cellHeight + mv.panY

			if mv.currentMapOrientation == hll.OriHorizontal {
				if j == 0 || j == 4 {
					continue
				}

				if active[i]+1 == j {
					// vector.DrawFilledRect(screen, float32(x), float32(y), float32(cellWidth), float32(cellHeight), fillColor, false)
				}
			}

			if mv.currentMapOrientation == hll.OriVertical {
				if i == 0 || i == 4 {
					continue
				}

				if active[j]+1 == i {
					// vector.DrawFilledRect(screen, float32(x), float32(y), float32(cellWidth), float32(cellHeight), fillColor, false)
				}
			}

			vector.StrokeLine(screen, float32(x), float32(y), float32(x+cellWidth), float32(y), 3, gridColor, false)
			vector.StrokeLine(screen, float32(x), float32(y+cellHeight), float32(x+cellWidth), float32(y+cellHeight), 3, gridColor, false)
			vector.StrokeLine(screen, float32(x), float32(y), float32(x), float32(y+cellHeight), 3, gridColor, false)
			vector.StrokeLine(screen, float32(x+cellWidth), float32(y), float32(x+cellWidth), float32(y+cellHeight), 3, gridColor, false)
		}
	}
}

func (mv *MapView) handleMouseInput() {
	sizeX, sizeY := mv.getDims()
	mouseX, mouseY := ebiten.CursorPosition()
	_, wheelY := ebiten.Wheel()
	if wheelY != 0 {
		oldZoom := mv.zoomLevel
		mv.zoomLevel += float64(wheelY * ZOOM_STEP_MULTIPLIER)
		if mv.zoomLevel < MIN_ZOOM_LEVEL {
			mv.zoomLevel = MIN_ZOOM_LEVEL
		} else if mv.zoomLevel > MAX_ZOOM_LEVEL {
			mv.zoomLevel = MAX_ZOOM_LEVEL
		}

		mouseWorldX := (float64(mouseX) - mv.panX) / oldZoom
		mouseWorldY := (float64(mouseY) - mv.panY) / oldZoom
		mv.panX -= mouseWorldX * (mv.zoomLevel - oldZoom)
		mv.panY -= mouseWorldY * (mv.zoomLevel - oldZoom)
	}

	if mv.zoomLevel == MIN_ZOOM_LEVEL {
		mv.panX = 0
		mv.panY = 0
	} else {
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
			x, y := ebiten.CursorPosition()
			if mv.isDragging {
				mv.panX += float64(x - mv.lastMouseX)
				mv.panY += float64(y - mv.lastMouseY)
			}
			mv.lastMouseX = x
			mv.lastMouseY = y
			mv.isDragging = true
		} else {
			mv.isDragging = false
		}

		mv.panX = util.Clamp(mv.panX, float64(sizeX)*(MIN_ZOOM_LEVEL-mv.zoomLevel), 0)
		mv.panY = util.Clamp(mv.panY, float64(sizeY)*(MIN_ZOOM_LEVEL-mv.zoomLevel), 0)
	}

	if mv.showPlayers && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		foundPlayer := false
		mouseX, mouseY := ebiten.CursorPosition()
		for _, player := range mv.playerMap {
			if !player.IsSpawned() {
				continue
			}

			x, y := util.TranslateCoords(sizeX, sizeY, player.Position)
			x = x*mv.zoomLevel + mv.panX
			y = y*mv.zoomLevel + mv.panY

			radius := util.IconCircleRadius(mv.zoomLevel, PLAYER_SIZE_MODIFIER)
			if (float64(mouseX)-x)*(float64(mouseX)-x)+(float64(mouseY)-y)*(float64(mouseY)-y) <= radius*radius {
				mv.selectedPlayerID = player.ID
				foundPlayer = true
				break
			}
		}
		if !foundPlayer {
			mv.selectedPlayerID = ""
		}
	}
}

func (mv *MapView) handleKeyboardInput() {
	typed := ebiten.AppendInputChars(nil)
	for _, r := range typed {
		typedKey := string(unicode.ToLower(r))

		for _, key := range util.Config.Keys.ToggleGridOverlay {
			if typedKey == strings.ToLower(key) {
				mv.showGrid = !mv.showGrid
				break
			}
		}

		for _, key := range util.Config.Keys.TogglePlayers {
			if typedKey == strings.ToLower(key) {
				mv.showPlayers = !mv.showPlayers
				break
			}
		}

		for _, key := range util.Config.Keys.TogglePlayerInfo {
			if typedKey == strings.ToLower(key) {
				mv.showPlayerInfo = !mv.showPlayerInfo
				break
			}
		}

		for _, key := range util.Config.Keys.ToggleSpawns {
			if typedKey == strings.ToLower(key) {
				mv.showSpawns = !mv.showSpawns
				break
			}
		}

		for _, key := range util.Config.Keys.ToggleServerInfoOverlay {
			if typedKey == strings.ToLower(key) {
				mv.showServerInfo = !mv.showServerInfo
				break
			}
		}

		for _, key := range util.Config.Keys.IncreaseInterval {
			if typedKey == strings.ToLower(key) {
				if mv.intervalIndex < len(fetchIntervalSteps)-1 {
					mv.intervalIndex++
				}
				break
			}
		}

		for _, key := range util.Config.Keys.DecreaseInterval {
			if typedKey == strings.ToLower(key) {
				if mv.intervalIndex > 0 {
					mv.intervalIndex--
				}
				break
			}
		}

		for _, key := range util.Config.Keys.Help {
			if typedKey == strings.ToLower(key) {
				mv.showHelp = !mv.showHelp
				break
			}
		}
	}

	mv.showScoreboard = false
	for _, key := range util.Config.Keys.ShowScoreboard {
		if ebiten.IsKeyPressed(util.MapKey(key)) {
			mv.showScoreboard = true
			break
		}
	}
}
