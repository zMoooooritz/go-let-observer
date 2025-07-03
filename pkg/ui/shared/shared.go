package shared

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
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
	CLR_AXIS_LIGHT            = color.RGBA{255, 120, 120, 255}
	CLR_AXIS                  = color.RGBA{255, 0, 0, 255}
	CLR_AXIS_DARK             = color.RGBA{180, 0, 30, 255}
	CLR_ALLIES_LIGHT          = color.RGBA{120, 120, 255, 255}
	CLR_ALLIES                = color.RGBA{0, 0, 255, 255}
	CLR_ALLIES_DARK           = color.RGBA{0, 0, 180, 255}
	CLR_SELECTED_LIGHT        = color.RGBA{120, 255, 120, 255}
	CLR_SELECTED              = color.RGBA{0, 255, 0, 255}
	CLR_BLACK                 = color.RGBA{0, 0, 0, 255}
	CLR_WHITE                 = color.RGBA{255, 255, 255, 255}
	CLR_OVERLAY               = color.RGBA{0, 0, 0, 200}
	CLR_ALLIES_OVERLAY        = color.RGBA{30, 30, 60, 50}
	CLR_AXIS_OVERLAY          = color.RGBA{60, 30, 30, 50}
	CLR_ACTIVE_SECTOR_OVERLAY = color.RGBA{30, 30, 30, 90}

	FALLBACK_BACKGROUND = color.RGBA{31, 31, 31, 255}
)

type PresentationMode string

const (
	MODE_VIEWER PresentationMode = "viewer"
	MODE_REPLAY PresentationMode = "replay"
	MODE_RECORD PresentationMode = "record"
	MODE_NONE   PresentationMode = ""
)

type ViewDimension struct {
	SizeX     int
	SizeY     int
	ZoomLevel float64
	PanX      float64
	PanY      float64
}

func (vd *ViewDimension) FrustumSize() (float64, float64) {
	screenSizeX := float64(vd.SizeX) * vd.ZoomLevel
	screenSizeY := float64(vd.SizeY) * vd.ZoomLevel
	return screenSizeX, screenSizeY
}

type State interface {
	Update() error
	Draw(screen *ebiten.Image)
	Layout(outsideWidth, outsideHeight int) (int, int)
}

type StateContext interface {
	TransitionTo(newState State)
}
