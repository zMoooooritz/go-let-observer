package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/zMoooooritz/go-let-loose/pkg/logger"
	"github.com/zMoooooritz/go-let-observer/pkg/ui"
	"github.com/zMoooooritz/go-let-observer/pkg/ui/shared"
	"github.com/zMoooooritz/go-let-observer/pkg/util"
)

var (
	// values set via ldflags
	Version   = ""
	CommitSHA = ""

	config    = flag.String("config", "", "Path to configuration file")
	host      = flag.String("host", "", "RCON server host")
	port      = flag.String("port", "", "RCON server port")
	password  = flag.String("password", "", "RCON server password")
	size      = flag.Int("size", shared.ROOT_SCALING_SIZE, "Screen size")
	directory = flag.String("directory", "", "Path to the replay directory")
	mode      = flag.String("mode", string(shared.MODE_NONE), "Mode to run on startup (viewer, replay, record)")
	version   = flag.Bool("version", false, "Display version information")
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

func extractCliParams() {
	if *host != "" {
		util.Config.ServerCredentials.Host = *host
	}
	if *port != "" {
		util.Config.ServerCredentials.Port = *port
	}
	if *password != "" {
		util.Config.ServerCredentials.Password = *password
	}
	if *size != shared.ROOT_SCALING_SIZE {
		util.Config.UIOptions.ScreenSize = *size
	}
	if *directory != "" {
		util.Config.ReplaysDirectory = *directory
	}
}

func main() {
	logger.DefaultLogger()

	flag.Parse()

	util.InitConfig(*config)

	if *version {
		showVersion()
	}

	viewerMode := shared.PresentationMode(*mode)
	if viewerMode != shared.MODE_NONE && viewerMode != shared.MODE_VIEWER && viewerMode != shared.MODE_REPLAY && viewerMode != shared.MODE_RECORD {
		fmt.Printf("Invalid startup mode. Allowed values are: %s, %s, %s.\n", shared.MODE_VIEWER, shared.MODE_REPLAY, shared.MODE_RECORD)
		os.Exit(1)
	}

	extractCliParams()

	if (viewerMode == shared.MODE_RECORD || viewerMode == shared.MODE_REPLAY) && util.Config.ReplaysDirectory == "" {
		fmt.Println("Replay directory is required in replay and record mode.")
		os.Exit(1)
	}

	screenSize := util.Config.UIOptions.ScreenSize
	screenSize = util.Clamp(screenSize, shared.MIN_SCREEN_SIZE, shared.MAX_SCREEN_SIZE)
	util.Config.UIOptions.ScreenSize = screenSize
	util.InitializeFonts(screenSize)
	userInterface := ui.NewUI(viewerMode)

	ebiten.SetWindowSize(screenSize, screenSize)
	ebiten.SetWindowTitle("HLL Observer")
	if err := ebiten.RunGame(userInterface); err != nil {
		log.Fatal(err)
	}
}
