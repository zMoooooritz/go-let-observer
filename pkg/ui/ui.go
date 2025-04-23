package ui

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/zMoooooritz/go-let-loose/pkg/rconv2"
	"github.com/zMoooooritz/go-let-observer/pkg/rcndata"
	"github.com/zMoooooritz/go-let-observer/pkg/record"
	"github.com/zMoooooritz/go-let-observer/pkg/util"
)

const (
	RCON_WORKER_COUNT = 3

	MIN_ZOOM_LEVEL       = 1.0
	MAX_ZOOM_LEVEL       = 10.0
	ZOOM_STEP_MULTIPLIER = 0.1

	MIN_SCREEN_SIZE   = 500
	ROOT_SCALING_SIZE = 1000
	MAX_SCREEN_SIZE   = 2500

	PLAYER_SIZE_MODIFIER          = 1.0
	SELECTED_PLAYER_SIZE_MODIFIER = 1.1
	TANK_SIZE_MODIFIER            = 1.2
	SPAWN_SIZE_MODIFIER           = 1.2
	TANK_ICON_SIZE_MODIFIER       = 1.8

	INITIAL_FETCH_STEP = 2
)

var (
	CLR_AXIS_LIGHT     = color.RGBA{255, 120, 120, 255}
	CLR_AXIS           = color.RGBA{255, 0, 0, 255}
	CLR_AXIS_DARK      = color.RGBA{180, 0, 30, 255}
	CLR_ALLIES_LIGHT   = color.RGBA{120, 120, 255, 255}
	CLR_ALLIES         = color.RGBA{0, 0, 255, 255}
	CLR_ALLIES_DARK    = color.RGBA{0, 0, 180, 255}
	CLR_SELECTED_LIGHT = color.RGBA{120, 255, 120, 255}
	CLR_SELECTED       = color.RGBA{0, 255, 0, 255}
	CLR_BLACK          = color.RGBA{0, 0, 0, 255}
	CLR_WHITE          = color.RGBA{255, 255, 255, 255}
	CLR_OVERLAY        = color.RGBA{0, 0, 0, 200}

	FALLBACK_BACKGROUND = color.RGBA{31, 31, 31, 255}
)

type State interface {
	Update() error
	Draw(screen *ebiten.Image)
	Layout(outsideWidth, outsideHeight int) (int, int)
}

type UI struct {
	state State
}

func getFetcherAndRecorder(rcon *rconv2.Rcon, recordPath string, replayPath string) (rcndata.DataFetcher, record.DataRecorder) {
	var dataFetcher rcndata.DataFetcher
	if replayPath == "" {
		dataFetcher = rcndata.NewRconDataFetcher(rcon)
	} else {
		var err error
		dataFetcher, err = record.NewMatchReplayer(replayPath)
		if err != nil {
			log.Fatal(err)
		}
	}

	var dataRecorder record.DataRecorder
	if recordPath != "" {
		currMap, err := rcon.GetCurrentMap()
		if err != nil {
			log.Fatal(err)
		}
		dataRecorder, err = record.NewMatchRecorder(recordPath, currMap)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		dataRecorder = &record.NoRecorder{}
	}

	return dataFetcher, dataRecorder
}

func NewUI(size int, rcon *rconv2.Rcon, recordPath string, replayPath string) *UI {
	util.ScaleFactor = float32(size) / float32(ROOT_SCALING_SIZE)

	ui := &UI{}

	dataFetcher, dataRecorder := getFetcherAndRecorder(rcon, recordPath, replayPath)
	showLoginView := rcon == nil && replayPath == ""

	bv := NewBaseViewer(size)

	if showLoginView {
		ui.state = NewLoginView(bv, ui.openMapView, recordPath)
	} else {
		ui.state = NewMapView(bv, dataFetcher, dataRecorder)
	}

	return ui
}

func (ui *UI) Update() error {

	return ui.state.Update()
}

func (ui *UI) Draw(screen *ebiten.Image) {
	ui.state.Draw(screen)
}

func (ui *UI) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ui.state.Layout(outsideWidth, outsideHeight)
}

func (ui *UI) openMapView(size int, rcon *rconv2.Rcon, recordPath string) {
	dataFetcher, dataRecorder := getFetcherAndRecorder(rcon, recordPath, "")
	bv := NewBaseViewer(size)
	ui.state = NewMapView(bv, dataFetcher, dataRecorder)
}
