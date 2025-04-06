package main

import (
	"flag"
	"log"
	"sync"
	"time"

	"image/color"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/zMoooooritz/go-let-loose/pkg/hll"
	"github.com/zMoooooritz/go-let-loose/pkg/rconv2"
)

const (
	roleCount = 14

	MIN_ZOOM_LEVEL       = 1.0
	MAX_ZOOM_LEVEL       = 10.0
	ZOOM_STEP_MULTIPLIER = 0.1

	RCON_FETCH_INTERVAL = 500 * time.Millisecond

	SCEEN_SIZE = 1000
)

var (
	RED   = color.RGBA{255, 0, 0, 255}
	GREEN = color.RGBA{0, 255, 0, 255}
	BLUE  = color.RGBA{0, 0, 255, 255}
	WHITE = color.RGBA{255, 255, 255, 255}
)

type GameStage int

const (
	StageLogin GameStage = iota
	StageMap
)

type Game struct {
	stage                 GameStage
	rcon                  *rconv2.Rcon
	currentMapName        string
	currentMapOrientation hll.Orientation
	serverName            string
	playerCurrCount       int
	playerMaxCount        int
	backgroundImage       *ebiten.Image
	sizeX                 int
	sizeY                 int
	players               map[string]hll.DetailedPlayerInfo
	lastUpdateTime        time.Time
	isFetching            bool
	fetchMutex            sync.Mutex
	zoomLevel             float64
	panX                  float64
	panY                  float64
	lastMouseX            int
	lastMouseY            int
	isDragging            bool
	roleImages            map[string]*ebiten.Image
	selectedPlayerID      string
	initialDataLoaded     bool
	showScoreboard        bool
	showGrid              bool
	gridKeyPressed        bool
	fnt                   Font

	// Login fields
	hostInput     string
	portInput     string
	passwordInput string
	activeField   int // 0: host, 1: port, 2: password
	errorMessage  string
}

func (g *Game) Update() error {
	// Handle Ctrl+C to exit the application
	if ebiten.IsKeyPressed(ebiten.KeyControl) && ebiten.IsKeyPressed(ebiten.KeyC) {
		return ebiten.Termination
	}

	switch g.stage {
	case StageLogin:
		return g.updateLogin()
	case StageMap:
		return g.updateMap()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch g.stage {
	case StageLogin:
		g.drawLogin(screen)
	case StageMap:
		g.drawMap(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.sizeX, g.sizeY
}

func NewGame() *Game {
	return &Game{
		stage:           StageLogin,
		sizeX:           SCEEN_SIZE,
		sizeY:           SCEEN_SIZE,
		backgroundImage: loadGreeterImage(),
		players:         make(map[string]hll.DetailedPlayerInfo),
		lastUpdateTime:  time.Now(),
		zoomLevel:       MIN_ZOOM_LEVEL,
		roleImages:      loadRoleImages(),
		fnt:             loadFonts(),
		showGrid:        true,
	}
}

func main() {
	// Define CLI flags
	host := flag.String("host", "", "RCON server host")
	port := flag.String("port", "", "RCON server port")
	password := flag.String("password", "", "RCON server password")
	flag.Parse()

	game := NewGame()
	// Check if CLI arguments are provided
	if *host != "" && *port != "" && *password != "" {
		cfg := rconv2.ServerConfig{
			Host:     *host,
			Port:     *port,
			Password: *password,
		}

		// Attempt to connect with the provided credentials
		rcon, err := rconv2.NewRcon(cfg, 3)
		if err != nil {
			log.Fatal("Invalid CLI credentials or connection error")
		}

		game.rcon = rcon
		game.stage = StageMap
	}

	ebiten.SetWindowSize(game.sizeX, game.sizeY)
	ebiten.SetWindowTitle("HLL Observer")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
