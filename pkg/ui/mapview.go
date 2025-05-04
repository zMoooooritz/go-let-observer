package ui

import (
	"fmt"
	"log"
	"sort"
	"sync"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/zMoooooritz/go-let-loose/pkg/hll"
	"github.com/zMoooooritz/go-let-observer/assets"
	"github.com/zMoooooritz/go-let-observer/pkg/rcndata"
	"github.com/zMoooooritz/go-let-observer/pkg/record"
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
	showTanks         bool
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
	currentMapName        string
	currentMapOrientation hll.Orientation
	serverName            string
	playerCurrCount       int
	playerMaxCount        int
	playerMap             map[string]hll.DetailedPlayerInfo
	playerList            []hll.DetailedPlayerInfo
	serverView            *hll.ServerView
	spawnTracker          *rcndata.SpawnTracker
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

	notifications *NotificationManager
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
			showTanks:      util.Config.UIOptions.ShowTanks,
		},
		FetchState: FetchState{
			intervalIndex: INITIAL_FETCH_STEP,
		},
		RconData: RconData{
			spawnTracker: rcndata.NewSpawnTracker(),
		},
		roleImages:    util.LoadRoleImages(),
		spawnImages:   util.LoadSpawnImages(),
		dataFetcher:   dataFetcher,
		dataRecorder:  dataRecorder,
		notifications: NewNotificationManager(),
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
		drawGrid(screen, mv.dim, mv.currentMapOrientation)
	}

	if mv.showSpawns && !mv.dataFetcher.IsUserSeekable() {
		drawSpawns(screen, mv.spawnTracker.GetSpawns(), mv.spawnImages, mv.dim)
	}

	if mv.showPlayers {
		drawPlayers(screen, mv.dim, mv.roleImages, mv.playerList, mv.selectedPlayerID)
	}

	if mv.showTanks {
		drawTankSquads(screen, mv.dim, mv.roleImages, mv.serverView, mv.selectedPlayerID)
	}

	if mv.showHelp {
		drawHelp(screen)
	} else if mv.showScoreboard {
		drawScoreboard(screen, mv.playerList)
	} else {
		if mv.showServerInfo {
			drawServerName(screen, mv.serverName)
			drawPlayerCount(screen, mv.playerCurrCount, mv.playerMaxCount)
		}
		if mv.showPlayerInfo {
			if player, ok := mv.playerMap[mv.selectedPlayerID]; ok {
				drawPlayerOverlay(screen, player)
			}
		}
	}

	if mv.dataFetcher.IsUserSeekable() {
		start, current, end := mv.dataFetcher.StartCurrentEndTime()
		progress := float64(current.Sub(start)) / float64(end.Sub(start))

		drawProgressBar(screen, progress)
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
	mv.currentMapOrientation = snapshot.CurrentMap.Orientation
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

	mv.initialDataLoaded = true
}
