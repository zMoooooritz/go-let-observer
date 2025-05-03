package ui

import (
	"image/color"

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

type PresentationMode string

const (
	MODE_VIEWER PresentationMode = "viewer"
	MODE_REPLAY PresentationMode = "replay"
	MODE_RECORD PresentationMode = "record"
	MODE_NONE   PresentationMode = ""
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

type StateContext interface {
	TransitionTo(newState State)
}

type UI struct {
	state State
}

func NewUI(mode PresentationMode) *UI {
	util.ScaleFactor = float32(util.Config.UIOptions.ScreenSize) / float32(ROOT_SCALING_SIZE)

	ui := &UI{}
	bv := NewBaseViewer(ui)

	switch mode {
	case MODE_NONE:
		ui.state = NewMainMenu(bv)
	case MODE_VIEWER:
		ui.state, _ = createState(bv, MODE_VIEWER, nil)
	case MODE_RECORD:
		ui.state, _ = createState(bv, MODE_RECORD, nil)
	case MODE_REPLAY:
		ui.state = NewReplayView(bv)
	}

	return ui
}

func (ui *UI) TransitionTo(newState State) {
	ui.state = newState
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

func createState(bv *BaseViewer, targetMode PresentationMode, rconCreds *rconv2.ServerConfig) (State, error) {
	if rconCreds == nil {
		creds := util.Config.GetServerCredentials()
		rconCreds = &creds
	}
	rcn, rcnErr := rconv2.NewRcon(*rconCreds, RCON_WORKER_COUNT)
	if rcnErr == nil {
		dataFetcher := rcndata.NewRconDataFetcher(rcn)

		var dataRecorder record.DataRecorder
		if targetMode == MODE_VIEWER {
			dataRecorder = record.NewNoRecorder()
		} else {
			currMap, _ := rcn.GetCurrentMap()
			dataRecorder, _ = record.NewMatchRecorder(util.Config.ReplaysDirectory, currMap)
		}

		return NewMapView(bv, dataFetcher, dataRecorder), nil
	}
	return NewLoginView(bv, targetMode), rcnErr
}
