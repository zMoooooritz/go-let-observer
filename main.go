package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/zMoooooritz/go-let-loose/pkg/rconv2"
	ui "github.com/zMoooooritz/go-let-observer/pkg/game"
	"github.com/zMoooooritz/go-let-observer/pkg/util"
)

var (
	// values set via ldflags
	Version   = ""
	CommitSHA = ""

	host     = flag.String("host", "", "RCON server host")
	port     = flag.String("port", "", "RCON server port")
	password = flag.String("password", "", "RCON server password")
	size     = flag.Int("size", ui.ROOT_SCALING_SIZE, "Screen size")
	version  = flag.Bool("version", false, "Display version")
)

func main() {
	flag.Parse()

	if *version {
		if len(CommitSHA) > 7 {
			CommitSHA = CommitSHA[:7]
		}
		if Version == "" {
			Version = "(built from source)"
		}

		fmt.Printf("go-let-observer %s", Version)
		if len(CommitSHA) > 0 {
			fmt.Printf(" (%s)", CommitSHA)
		}

		fmt.Println()
		os.Exit(0)
	}

	var rcon *rconv2.Rcon
	if *host != "" && *port != "" && *password != "" {
		cfg := rconv2.ServerConfig{
			Host:     *host,
			Port:     *port,
			Password: *password,
		}

		// Attempt to connect with the provided credentials
		var err error
		rcon, err = rconv2.NewRcon(cfg, ui.RCON_WORKER_COUNT)
		if err != nil {
			log.Fatal("Invalid CLI credentials or connection error")
		}
	}
	size := util.Clamp(*size, ui.MIN_SCREEN_SIZE, ui.MAX_SCREEN_SIZE)
	util.InitializeFonts(size)
	userInterface := ui.NewUI(size, rcon)

	ebiten.SetWindowSize(size, size)
	ebiten.SetWindowTitle("HLL Observer")
	if err := ebiten.RunGame(userInterface); err != nil {
		log.Fatal(err)
	}
}
