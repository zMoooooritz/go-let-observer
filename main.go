package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/zMoooooritz/go-let-loose/pkg/rconv2"
	"github.com/zMoooooritz/go-let-observer/pkg/ui"
	"github.com/zMoooooritz/go-let-observer/pkg/util"
)

var (
	// values set via ldflags
	Version   = ""
	CommitSHA = ""

	config   = flag.String("config", "", "Path to config file")
	host     = flag.String("host", "", "RCON server host")
	port     = flag.String("port", "", "RCON server port")
	password = flag.String("password", "", "RCON server password")
	size     = flag.Int("size", ui.ROOT_SCALING_SIZE, "Screen size")
	version  = flag.Bool("version", false, "Display version")
)

func main() {
	flag.Parse()

	err := util.InitConfig(*config)

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

	serverHost := ""
	serverPort := ""
	serverPassword := ""
	screenSize := 0

	if err == nil {
		if util.Config.ServerCredentials.Host != "" {
			serverHost = util.Config.ServerCredentials.Host
		}
		if util.Config.ServerCredentials.Port != "" {
			serverPort = util.Config.ServerCredentials.Port
		}
		if util.Config.ServerCredentials.Password != "" {
			serverPassword = util.Config.ServerCredentials.Password
		}
		if util.Config.UIStartupOptions.ScreenSize != 0 {
			screenSize = util.Config.UIStartupOptions.ScreenSize
		}
	}

	if *host != "" {
		serverHost = *host
	}
	if *port != "" {
		serverPort = *port
	}
	if *password != "" {
		serverPassword = *password
	}
	if *size != ui.ROOT_SCALING_SIZE {
		screenSize = *size
	}

	var rcon *rconv2.Rcon
	if serverHost != "" && serverPort != "" && serverPassword != "" {
		cfg := rconv2.ServerConfig{
			Host:     serverHost,
			Port:     serverPort,
			Password: serverPassword,
		}

		// Attempt to connect with the provided credentials
		var err error
		rcon, err = rconv2.NewRcon(cfg, ui.RCON_WORKER_COUNT)
		if err != nil {
			log.Fatal("Invalid CLI credentials or connection error")
		}
	}

	screenSize = util.Clamp(screenSize, ui.MIN_SCREEN_SIZE, ui.MAX_SCREEN_SIZE)
	util.InitializeFonts(screenSize)
	userInterface := ui.NewUI(screenSize, rcon)

	ebiten.SetWindowSize(screenSize, screenSize)
	ebiten.SetWindowTitle("HLL Observer")
	if err := ebiten.RunGame(userInterface); err != nil {
		log.Fatal(err)
	}
}
