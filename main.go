package main

import (
	"flag"
	"log"
	"time"

	"image/color"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/zMoooooritz/go-let-loose/pkg/rconv2"
)

const (
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
		game.uiState = StateMap
	}

	ebiten.SetWindowSize(game.dim.sizeX, game.dim.sizeY)
	ebiten.SetWindowTitle("HLL Observer")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
