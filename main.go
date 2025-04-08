package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/zMoooooritz/go-let-loose/pkg/rconv2"
	"github.com/zMoooooritz/go-let-observer/pkg/game"
	"github.com/zMoooooritz/go-let-observer/pkg/util"
)

var (
	// values set via ldflags
	Version   = ""
	CommitSHA = ""
	Date      = ""

	host     = flag.String("host", "", "RCON server host")
	port     = flag.String("port", "", "RCON server port")
	password = flag.String("password", "", "RCON server password")
	size     = flag.Int("size", game.ROOT_SCALING_SIZE, "Screen size")
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
		rcon, err = rconv2.NewRcon(cfg, 3)
		if err != nil {
			log.Fatal("Invalid CLI credentials or connection error")
		}
	}
	size := util.Clamp(*size, game.MIN_SCREEN_SIZE, game.MAX_SCREEN_SIZE)
	gm := game.NewGame(size, rcon)

	ebiten.SetWindowSize(size, size)
	ebiten.SetWindowTitle("HLL Observer")
	if err := ebiten.RunGame(gm); err != nil {
		log.Fatal(err)
	}
}
