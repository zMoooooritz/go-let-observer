package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/zMoooooritz/go-let-loose/pkg/logger"
	"github.com/zMoooooritz/go-let-loose/pkg/rconv2"
	"github.com/zMoooooritz/go-let-observer/pkg/ui"
	"github.com/zMoooooritz/go-let-observer/pkg/util"
)

var (
	// values set via ldflags
	Version   = ""
	CommitSHA = ""

	config     = flag.String("config", "", "Path to configuration file")
	host       = flag.String("host", "", "RCON server host")
	port       = flag.String("port", "", "RCON server port")
	password   = flag.String("password", "", "RCON server password")
	size       = flag.Int("size", ui.ROOT_SCALING_SIZE, "Screen size")
	recordPath = flag.String("record", "", "Path to the recording directory")
	replayPath = flag.String("replay", "", "Path to match data JSON file")
	version    = flag.Bool("version", false, "Display version information")
)

func showVersion() {
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

func extractConfigParams(configLoadErr error) (rconv2.ServerConfig, int) {
	cfg := rconv2.ServerConfig{}

	screenSize := 0

	if configLoadErr == nil {
		if util.Config.ServerCredentials.Host != "" {
			cfg.Host = util.Config.ServerCredentials.Host
		}
		if util.Config.ServerCredentials.Port != "" {
			cfg.Port = util.Config.ServerCredentials.Port
		}
		if util.Config.ServerCredentials.Password != "" {
			cfg.Password = util.Config.ServerCredentials.Password
		}
		if util.Config.UIStartupOptions.ScreenSize != 0 {
			screenSize = util.Config.UIStartupOptions.ScreenSize
		}
	}

	if *host != "" {
		cfg.Host = *host
	}
	if *port != "" {
		cfg.Port = *port
	}
	if *password != "" {
		cfg.Password = *password
	}
	if *size != ui.ROOT_SCALING_SIZE {
		screenSize = *size
	}

	return cfg, screenSize
}

func initRcon(cfg rconv2.ServerConfig) *rconv2.Rcon {
	var rcon *rconv2.Rcon
	if cfg.Host != "" && cfg.Port != "" && cfg.Password != "" {
		var err error
		rcon, err = rconv2.NewRcon(cfg, ui.RCON_WORKER_COUNT)
		if err != nil {
			log.Fatal("Invalid CLI credentials or connection error")
		}
	}
	return rcon
}

func main() {
	logger.DefaultLogger()

	flag.Parse()

	err := util.InitConfig(*config)

	if *version {
		showVersion()
	}

	if *recordPath != "" && *replayPath != "" {
		log.Fatal("Cannot record and replay at the same time")
	}

	cfg, screenSize := extractConfigParams(err)

	var rcon *rconv2.Rcon
	if *replayPath == "" {
		rcon = initRcon(cfg)
	}

	screenSize = util.Clamp(screenSize, ui.MIN_SCREEN_SIZE, ui.MAX_SCREEN_SIZE)
	util.InitializeFonts(screenSize)
	userInterface := ui.NewUI(screenSize, rcon, *recordPath, *replayPath)

	ebiten.SetWindowSize(screenSize, screenSize)
	ebiten.SetWindowTitle("HLL Observer")
	if err := ebiten.RunGame(userInterface); err != nil {
		log.Fatal(err)
	}
}
