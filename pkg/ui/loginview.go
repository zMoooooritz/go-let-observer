package ui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/zMoooooritz/go-let-loose/pkg/rconv2"
	"github.com/zMoooooritz/go-let-observer/pkg/util"
)

type LoginView struct {
	*BaseViewer

	activeField   int
	hostInput     string
	portInput     string
	passwordInput string
	errorMessage  string

	openMapView func(int, *rconv2.Rcon, string)
	recordPath  string
}

func NewLoginView(bv *BaseViewer, openMapView func(int, *rconv2.Rcon, string), recordPath string) *LoginView {
	lv := &LoginView{
		BaseViewer:  bv,
		openMapView: openMapView,
		recordPath:  recordPath,
	}
	lv.backgroundImage = util.LoadGreeterImage()
	return lv
}

func (lv *LoginView) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyControl) && ebiten.IsKeyPressed(ebiten.KeyC) || ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return ebiten.Termination
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyTab) {
		lv.activeField = (lv.activeField + 1) % 3
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		if lv.hostInput != "" && lv.portInput != "" && lv.passwordInput != "" {
			cfg := rconv2.ServerConfig{
				Host:     lv.hostInput,
				Port:     lv.portInput,
				Password: lv.passwordInput,
			}
			rcn, err := rconv2.NewRcon(cfg, RCON_WORKER_COUNT)
			if err != nil {
				lv.errorMessage = "Invalid credentials or connection error"
			} else {
				lv.openMapView(lv.dim.sizeX, rcn, lv.recordPath)
			}
		} else {
			lv.errorMessage = "All fields are required"
		}
	}

	if chars := ebiten.AppendInputChars(nil); len(chars) > 0 {
		for _, c := range chars {
			switch lv.activeField {
			case 0:
				lv.hostInput += string(c)
			case 1:
				lv.portInput += string(c)
			case 2:
				lv.passwordInput += string(c)
			}
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) {
		switch lv.activeField {
		case 0:
			if len(lv.hostInput) > 0 {
				lv.hostInput = lv.hostInput[:len(lv.hostInput)-1]
			}
		case 1:
			if len(lv.portInput) > 0 {
				lv.portInput = lv.portInput[:len(lv.portInput)-1]
			}
		case 2:
			if len(lv.passwordInput) > 0 {
				lv.passwordInput = lv.passwordInput[:len(lv.passwordInput)-1]
			}
		}
	}

	return nil
}

func (lv *LoginView) Draw(screen *ebiten.Image) {
	lv.DrawBackground(screen)

	util.DrawScaledRect(screen, 0, 0, 1000, 400, CLR_OVERLAY)

	util.DrawText(screen, "Login to HLL Observer", 20, 40, CLR_WHITE, util.Font.Title)

	util.DrawText(screen, "Host:", 50, 100, CLR_WHITE, util.Font.Normal)
	util.DrawText(screen, "Port:", 50, 160, CLR_WHITE, util.Font.Normal)
	util.DrawText(screen, "Password:", 50, 220, CLR_WHITE, util.Font.Normal)

	hostRectColor := CLR_WHITE
	portRectColor := CLR_WHITE
	passwordRectColor := CLR_WHITE

	switch lv.activeField {
	case 0:
		hostRectColor = CLR_SELECTED
	case 1:
		portRectColor = CLR_SELECTED
	case 2:
		passwordRectColor = CLR_SELECTED
	}

	util.DrawScaledRect(screen, 180, 80, 300, 30, hostRectColor)
	util.DrawScaledRect(screen, 180, 140, 300, 30, portRectColor)
	util.DrawScaledRect(screen, 180, 200, 300, 30, passwordRectColor)

	util.DrawText(screen, lv.hostInput, 185, 100, CLR_BLACK, util.Font.Normal)
	util.DrawText(screen, lv.portInput, 185, 160, CLR_BLACK, util.Font.Normal)
	util.DrawText(screen, lv.passwordInput, 185, 220, CLR_BLACK, util.Font.Normal)

	if lv.errorMessage != "" {
		util.DrawText(screen, lv.errorMessage, 50, 280, color.RGBA{255, 0, 0, 255}, util.Font.Normal)
	}

	util.DrawText(screen, "Press Enter to confirm, Tab to switch fields", 50, 340, color.Gray{Y: 200}, util.Font.Normal)
}
