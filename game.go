package main

import (
	"sync"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/zMoooooritz/go-let-loose/pkg/hll"
	"github.com/zMoooooritz/go-let-loose/pkg/rconv2"
)

type UIState int

const (
	StateLogin UIState = iota
	StateMap
)

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
	showScoreboard    bool
	showGrid          bool
	gridKeyPressed    bool
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
	players               map[string]hll.DetailedPlayerInfo
}

type MapView struct {
	MapViewState
	CameraState
	FetchState
	RconData
	roleImages map[string]*ebiten.Image
}

type Game struct {
	fnt             Font
	uiState         UIState
	rcon            *rconv2.Rcon
	backgroundImage *ebiten.Image

	dim       ViewDimension
	loginView *LoginView
	mapView   *MapView
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

func NewGame() *Game {
	dim := ViewDimension{
		sizeX: SCEEN_SIZE,
		sizeY: SCEEN_SIZE,
	}
	mapView := &MapView{
		CameraState: CameraState{
			zoomLevel: MIN_ZOOM_LEVEL,
		},
		MapViewState: MapViewState{
			showGrid: true,
		},
		roleImages: loadRoleImages(),
	}

	return &Game{
		fnt:             loadFonts(),
		uiState:         StateLogin,
		backgroundImage: loadGreeterImage(),
		dim:             dim,
		mapView:         mapView,
	}
}
