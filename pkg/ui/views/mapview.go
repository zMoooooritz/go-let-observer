package views

import (
	"fmt"
	"log"
	"sort"
	"sync"
	"time"
	"unicode"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/zMoooooritz/go-let-loose/pkg/hll"
	"github.com/zMoooooritz/go-let-observer/assets"
	"github.com/zMoooooritz/go-let-observer/pkg/rcndata"
	"github.com/zMoooooritz/go-let-observer/pkg/record"
	"github.com/zMoooooritz/go-let-observer/pkg/ui/components"
	"github.com/zMoooooritz/go-let-observer/pkg/ui/shared"
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
	showVehicles      bool
	showHelp          bool
	showScoreboard    bool
	initialDataLoaded bool
	selectedPlayerID  string
}

type InputState struct {
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
	currentMapName  string
	currentMapID    string
	serverName      string
	playerCurrCount int
	playerMaxCount  int
	playerMap       map[string]hll.DetailedPlayerInfo
	playerList      []hll.DetailedPlayerInfo
	serverView      *hll.ServerView
	spawnTracker    *rcndata.SpawnTracker
	gameScore       hll.TeamData
	remainingTime   time.Duration
}

type MapView struct {
	*BaseViewer
	MapViewState
	InputState
	FetchState
	RconData
	roleImages  map[string]*ebiten.Image
	spawnImages map[string]*ebiten.Image

	dataFetcher  rcndata.DataFetcher
	dataRecorder record.DataRecorder

	notifications *components.NotificationManager
}

func NewMapView(bv *BaseViewer, dataFetcher rcndata.DataFetcher, dataRecorder record.DataRecorder) *MapView {
	mv := &MapView{
		BaseViewer: bv,
		MapViewState: MapViewState{
			showServerInfo: util.Config.UIOptions.ShowServerInfoOverlay,
			showGrid:       util.Config.UIOptions.ShowGridOverlay,
			showPlayers:    util.Config.UIOptions.ShowPlayers,
			showPlayerInfo: util.Config.UIOptions.ShowPlayerInfo,
			showSpawns:     util.Config.UIOptions.ShowSpawns,
			showVehicles:   util.Config.UIOptions.ShowVehicles,
		},
		FetchState: FetchState{
			intervalIndex: shared.INITIAL_FETCH_STEP,
		},
		RconData: RconData{
			spawnTracker: rcndata.NewSpawnTracker(),
		},
		roleImages:    util.LoadRoleImages(),
		spawnImages:   util.LoadSpawnImages(),
		dataFetcher:   dataFetcher,
		dataRecorder:  dataRecorder,
		notifications: components.NewNotificationManager(),
	}
	mv.backgroundImage = util.LoadGreeterImage()
	return mv
}

func (mv *MapView) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyControl) && ebiten.IsKeyPressed(ebiten.KeyC) || ebiten.IsKeyPressed(ebiten.KeyEscape) || ebiten.IsKeyPressed(ebiten.KeyQ) {
		mv.dataRecorder.Stop()
		return ebiten.Termination
	}

	if time.Since(mv.lastUpdateTime) >= fetchIntervalSteps[mv.intervalIndex] {
		mv.fetchMutex.Lock()
		if !mv.isFetching {
			mv.isFetching = true
			go func() {
				snapshot, err := mv.dataFetcher.FetchRconDataSnapshot()
				if err == nil {
					mv.processRconData(snapshot)
					mv.dataRecorder.RecordSnapshot(snapshot)
				} else {
					fmt.Println(err)
				}
				mv.fetchMutex.Lock()
				mv.lastUpdateTime = time.Now()
				mv.isFetching = false

				if !mv.dataFetcher.IsPaused() {
					mv.dataFetcher.Seek(10 * fetchIntervalSteps[mv.intervalIndex])
				}

				mv.fetchMutex.Unlock()
			}()
		}
		mv.fetchMutex.Unlock()
	}

	mv.handleKeyboardInput()
	mv.handleMouseInput()

	mv.notifications.Update()

	return nil
}

func (mv *MapView) Draw(screen *ebiten.Image) {
	mv.DrawBackground(screen)

	if !mv.initialDataLoaded {
		return
	}

	if mv.showGrid {
		components.DrawGrid(screen, mv.dim, mv.currentMapID, mv.gameScore)
	}

	if mv.showSpawns && !mv.dataFetcher.IsUserSeekable() {
		components.DrawSpawns(screen, mv.spawnTracker.GetSpawns(), mv.spawnImages, mv.dim)
	}

	if mv.showPlayers {
		components.DrawPlayers(screen, mv.dim, mv.roleImages, mv.playerList, mv.selectedPlayerID)
	}

	if mv.showVehicles {
		components.DrawVehicleSquads(screen, mv.dim, mv.roleImages, mv.serverView, mv.selectedPlayerID)
	}

	if mv.showHelp {
		components.DrawHelp(screen)
	} else if mv.showScoreboard {
		components.DrawScoreboard(screen, mv.playerList)
	} else {
		if mv.showServerInfo && !mv.dataFetcher.IsUserSeekable() {
			components.DrawServerName(screen, mv.serverName)
			components.DrawPlayerCount(screen, mv.playerCurrCount, mv.playerMaxCount)
		}
		if mv.showPlayerInfo {
			if player, ok := mv.playerMap[mv.selectedPlayerID]; ok {
				components.DrawPlayerInfoOverlay(screen, player)
			}
		}
	}

	if mv.dataFetcher.IsUserSeekable() {
		start, current, end := mv.dataFetcher.StartCurrentEndTime()
		progress := float64(current.Sub(start)) / float64(end.Sub(start))

		components.DrawProgressBar(screen, progress)
	}

	mv.notifications.Draw(screen)
}

func (mv *MapView) processRconData(snapshot *rcndata.RconDataSnapshot) {
	oldPlayerMap := mv.playerMap
	playerMap := map[string]hll.DetailedPlayerInfo{}
	for _, player := range snapshot.Players {
		playerMap[player.ID] = player
		if !mv.dataFetcher.IsUserSeekable() {
			if oldPlayer, ok := oldPlayerMap[player.ID]; ok {
				mv.spawnTracker.TrackPlayerPosition(oldPlayer, player)
			}
		}
	}

	serverView := hll.PlayersToServerView(snapshot.Players)
	mv.serverView = serverView

	sort.Slice(snapshot.Players, func(i, j int) bool {
		return snapshot.Players[i].ID > snapshot.Players[j].ID
	})

	if !mv.dataFetcher.IsUserSeekable() {
		mv.spawnTracker.CleanExpiredSpawns()
	}

	mv.playerMap = playerMap
	mv.playerList = snapshot.Players

	currMapName := assets.ToFileName(snapshot.CurrentMap.ID)
	mv.currentMapID = string(snapshot.CurrentMap.ID)
	if currMapName != mv.currentMapName {
		if !mv.dataFetcher.IsUserSeekable() {
			mv.spawnTracker.ResetSpawns()
		}

		if mv.currentMapName != "" && currMapName != "" {
			mv.dataRecorder.MapChanged(snapshot.CurrentMap)
		}

		mv.currentMapName = currMapName
		img, err := util.LoadMapImage(currMapName)
		if err == nil {
			mv.backgroundImage = img
		} else {
			log.Println("Error loading tacmap image:", err)
		}
	}

	mv.serverName = snapshot.SessionInfo.ServerName
	mv.playerCurrCount = snapshot.SessionInfo.PlayerCount
	mv.playerMaxCount = snapshot.SessionInfo.MaxPlayerCount
	mv.gameScore = hll.TeamData{
		Allies: snapshot.SessionInfo.AlliedScore,
		Axis:   snapshot.SessionInfo.AxisScore,
	}
	mv.remainingTime = snapshot.SessionInfo.RemainingMatchTime

	mv.initialDataLoaded = true
}

func (mv *MapView) handleMouseInput() {
	mouseX, mouseY := ebiten.CursorPosition()
	_, wheelY := ebiten.Wheel()
	if wheelY != 0 {
		oldZoom := mv.dim.ZoomLevel
		mv.dim.ZoomLevel += float64(wheelY * shared.ZOOM_STEP_MULTIPLIER)
		if mv.dim.ZoomLevel < shared.MIN_ZOOM_LEVEL {
			mv.dim.ZoomLevel = shared.MIN_ZOOM_LEVEL
		} else if mv.dim.ZoomLevel > shared.MAX_ZOOM_LEVEL {
			mv.dim.ZoomLevel = shared.MAX_ZOOM_LEVEL
		}

		mouseWorldX := (float64(mouseX) - mv.dim.PanX) / oldZoom
		mouseWorldY := (float64(mouseY) - mv.dim.PanY) / oldZoom
		mv.dim.PanX -= mouseWorldX * (mv.dim.ZoomLevel - oldZoom)
		mv.dim.PanY -= mouseWorldY * (mv.dim.ZoomLevel - oldZoom)
	}

	if mv.dim.ZoomLevel == shared.MIN_ZOOM_LEVEL {
		mv.dim.PanX = 0
		mv.dim.PanY = 0
	} else {
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
			x, y := ebiten.CursorPosition()
			if mv.isDragging {
				mv.dim.PanX += float64(x - mv.lastMouseX)
				mv.dim.PanY += float64(y - mv.lastMouseY)
			}
			mv.lastMouseX = x
			mv.lastMouseY = y
			mv.isDragging = true
		} else {
			mv.isDragging = false
		}

		mv.dim.PanX = util.Clamp(mv.dim.PanX, float64(mv.dim.SizeX)*(shared.MIN_ZOOM_LEVEL-mv.dim.ZoomLevel), 0)
		mv.dim.PanY = util.Clamp(mv.dim.PanY, float64(mv.dim.SizeY)*(shared.MIN_ZOOM_LEVEL-mv.dim.ZoomLevel), 0)
	}

	if mv.showPlayers && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		foundPlayer := false
		mouseX, mouseY := ebiten.CursorPosition()
		for _, player := range mv.playerMap {
			if !player.IsSpawned() {
				continue
			}

			x, y := util.TranslateCoords(mv.dim.SizeX, mv.dim.SizeY, player.Position)
			x = x*mv.dim.ZoomLevel + mv.dim.PanX
			y = y*mv.dim.ZoomLevel + mv.dim.PanY

			radius := util.IconCircleRadius(mv.dim.ZoomLevel, shared.PLAYER_SIZE_MODIFIER)
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

		if typedKey == "g" {
			mv.showGrid = !mv.showGrid
			mv.notifications.Push(fmt.Sprintf("Show Grid: %t", mv.showGrid))
		}

		if typedKey == "p" {
			mv.showPlayers = !mv.showPlayers
			mv.notifications.Push(fmt.Sprintf("Show Players: %t", mv.showPlayers))
		}

		if typedKey == "i" {
			mv.showPlayerInfo = !mv.showPlayerInfo
			mv.notifications.Push(fmt.Sprintf("Show Player Info: %t", mv.showPlayerInfo))
		}

		if typedKey == "s" {
			mv.showSpawns = !mv.showSpawns
			mv.notifications.Push(fmt.Sprintf("Show Spawns: %t", mv.showSpawns))
		}

		if typedKey == "t" {
			mv.showVehicles = !mv.showVehicles
			mv.notifications.Push(fmt.Sprintf("Show Vehicles: %t", mv.showVehicles))
		}

		if typedKey == "h" {
			mv.showServerInfo = !mv.showServerInfo
			mv.notifications.Push(fmt.Sprintf("Show Server Info: %t", mv.showServerInfo))
		}

		if typedKey == "+" {
			if mv.intervalIndex < len(fetchIntervalSteps)-1 {
				mv.intervalIndex++
			}
			mv.notifications.Push(fmt.Sprintf("Fetch-Interval: %s", fetchIntervalSteps[mv.intervalIndex]))
		}

		if typedKey == "-" {
			if mv.intervalIndex > 0 {
				mv.intervalIndex--
			}
			mv.notifications.Push(fmt.Sprintf("Fetch-Interval: %s", fetchIntervalSteps[mv.intervalIndex]))
		}

		if typedKey == "?" {
			mv.showHelp = !mv.showHelp
		}

	}

	mv.showScoreboard = false
	if ebiten.IsKeyPressed(ebiten.KeyTab) {
		mv.showScoreboard = true
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		if mv.dataFetcher.IsPaused() {
			mv.dataFetcher.Continue()
		} else {
			mv.dataFetcher.Pause()
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		duration := -time.Second
		if ebiten.IsKeyPressed(ebiten.KeyShift) {
			duration = -time.Minute
		}
		mv.dataFetcher.Seek(duration)
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		mv.dataFetcher.Seek(-2 * time.Hour)
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		duration := time.Second
		if ebiten.IsKeyPressed(ebiten.KeyShift) {
			duration = time.Minute
		}
		mv.dataFetcher.Seek(duration)
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		mv.dataFetcher.Seek(2 * time.Hour)
	}
}
