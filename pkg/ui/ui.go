package ui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/zMoooooritz/go-let-loose/pkg/rconv2"
	"github.com/zMoooooritz/go-let-observer/pkg/util"
)

const (
	RCON_WORKER_COUNT = 5

	MIN_ZOOM_LEVEL       = 1.0
	MAX_ZOOM_LEVEL       = 10.0
	ZOOM_STEP_MULTIPLIER = 0.1

	MIN_SCREEN_SIZE   = 500
	ROOT_SCALING_SIZE = 1000
	MAX_SCREEN_SIZE   = 2500

	PLAYER_SIZE_MODIFIER          = 1.0
	SELECTED_PLAYER_SIZE_MODIFIER = 1.1
	SPAWN_SIZE_MODIFIER           = 1.2

	INITIAL_FETCH_STEP = 2
)

var (
	CLR_AXIS         = color.RGBA{255, 0, 0, 255}
	CLR_ALLIES       = color.RGBA{0, 0, 255, 255}
	CLR_AXIS_SPAWN   = color.RGBA{180, 0, 30, 255}
	CLR_ALLIES_SPAWN = color.RGBA{0, 0, 180, 255}
	CLR_SELECTED     = color.RGBA{0, 255, 0, 255}
	CLR_BLACK        = color.RGBA{0, 0, 0, 255}
	CLR_WHITE        = color.RGBA{255, 255, 255, 255}
	CLR_OVERLAY      = color.RGBA{0, 0, 0, 200}
)

type ViewDimension struct {
	sizeX int
	sizeY int
}

func (vd *ViewDimension) getDims() (int, int) {
	return vd.sizeX, vd.sizeY
}

type State interface {
	Update() error
	Draw(screen *ebiten.Image)
}

type UI struct {
	dim   *ViewDimension
	state State
}

func NewUI(size int, rcon *rconv2.Rcon) *UI {
	util.ScaleFactor = float32(size) / float32(ROOT_SCALING_SIZE)

	dim := &ViewDimension{
		sizeX: size,
		sizeY: size,
	}

	ui := &UI{
		dim: dim,
	}

	if rcon == nil {
		ui.state = NewLoginView(ui.openMapView)
	} else {
		ui.state = NewMapView(rcon, dim.getDims)
	}

	return ui
}

func (ui *UI) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyControl) && ebiten.IsKeyPressed(ebiten.KeyC) || ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return ebiten.Termination
	}

	return ui.state.Update()
}

func (ui *UI) Draw(screen *ebiten.Image) {
	ui.state.Draw(screen)
}

func (ui *UI) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ui.dim.sizeX, ui.dim.sizeX
}

func (ui *UI) openMapView(rcon *rconv2.Rcon) {
	ui.state = NewMapView(rcon, ui.dim.getDims)
}
