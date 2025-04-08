package game

import (
	"fmt"
	"image/color"
	"sync"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/zMoooooritz/go-let-loose/pkg/hll"
	"github.com/zMoooooritz/go-let-loose/pkg/rconv2"
	"github.com/zMoooooritz/go-let-observer/pkg/util"
)

const (
	MIN_ZOOM_LEVEL       = 1.0
	MAX_ZOOM_LEVEL       = 10.0
	ZOOM_STEP_MULTIPLIER = 0.1

	MIN_SCREEN_SIZE   = 500
	ROOT_SCALING_SIZE = 1000
	MAX_SCREEN_SIZE   = 2500
)

var (
	CLR_AXIS     = color.RGBA{255, 0, 0, 255}
	CLR_ALLIES   = color.RGBA{0, 0, 255, 255}
	CLR_SELECTED = color.RGBA{0, 255, 0, 255}
	CLR_BLACK    = color.RGBA{0, 0, 0, 255}
	CLR_WHITE    = color.RGBA{255, 255, 255, 255}
	CLR_OVERLAY  = color.RGBA{0, 0, 0, 200}
)

type UIState int

const (
	StateLogin UIState = iota
	StateMap
)

var fetchIntervalSteps = []time.Duration{
	100 * time.Millisecond,
	500 * time.Millisecond,
	time.Second,
	2 * time.Second,
	5 * time.Second,
	10 * time.Second,
}

type ViewDimension struct {
	sizeX int
	sizeY int
}

type LoginView struct {
	activeField   int
	hostInput     string
	portInput     string
	passwordInput string
	errorMessage  string
}

type MapViewState struct {
	showHeader        bool
	showGrid          bool
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
}

type MapView struct {
	MapViewState
	CameraState
	FetchState
	RconData
	roleImages map[string]*ebiten.Image
}

type Game struct {
	fnt             util.Font
	uiState         UIState
	rcon            *rconv2.Rcon
	backgroundImage *ebiten.Image

	dim       ViewDimension
	loginView *LoginView
	mapView   *MapView
}

func NewGame(size int, rcon *rconv2.Rcon) *Game {
	util.ScaleFactor = float32(size) / float32(ROOT_SCALING_SIZE)
	fmt.Println(util.ScaleFactor)

	dim := ViewDimension{
		sizeX: size,
		sizeY: size,
	}
	loginView := &LoginView{}
	mapView := &MapView{
		MapViewState: MapViewState{
			showHeader: true,
			showGrid:   true,
		},
		CameraState: CameraState{
			zoomLevel: MIN_ZOOM_LEVEL,
		},
		FetchState: FetchState{
			intervalIndex: 2,
		},
		roleImages: util.LoadRoleImages(),
	}
	uiState := StateLogin
	if rcon != nil {
		uiState = StateMap
	}

	return &Game{
		fnt:             util.LoadFonts(size),
		uiState:         uiState,
		rcon:            rcon,
		backgroundImage: util.LoadGreeterImage(),
		dim:             dim,
		loginView:       loginView,
		mapView:         mapView,
	}
}

func (g *Game) Update() error {
	// Handle Ctrl+C to exit the application
	if ebiten.IsKeyPressed(ebiten.KeyControl) && ebiten.IsKeyPressed(ebiten.KeyC) || ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return ebiten.Termination
	}

	switch g.uiState {
	case StateLogin:
		return g.updateLogin()
	case StateMap:
		return g.updateMap()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch g.uiState {
	case StateLogin:
		g.drawLogin(screen)
	case StateMap:
		g.drawMap(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.dim.sizeX, g.dim.sizeX
}
