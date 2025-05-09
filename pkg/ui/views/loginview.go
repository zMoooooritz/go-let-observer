package views

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/zMoooooritz/go-let-loose/pkg/rconv2"
	"github.com/zMoooooritz/go-let-observer/pkg/ui/shared"
	"github.com/zMoooooritz/go-let-observer/pkg/util"
)

type LoginView struct {
	*BaseViewer

	targetMode shared.PresentationMode

	activeField   int
	hostInput     string
	portInput     string
	passwordInput string
	errorMessage  string
}

func NewLoginView(bv *BaseViewer, targetMode shared.PresentationMode) *LoginView {
	lv := &LoginView{
		BaseViewer: bv,
		targetMode: targetMode,
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

			bv := NewBaseViewer(lv.ctx)
			state, err := CreateState(bv, lv.targetMode, &cfg)
			if err != nil {
				lv.errorMessage = "Invalid credentials or connection error"
			} else {
				lv.ctx.TransitionTo(state)
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

	util.DrawScaledRect(screen, 0, 0, 1000, 400, shared.CLR_OVERLAY)

	util.DrawText(screen, "Login to HLL Observer", 20, 40, shared.CLR_WHITE, util.Font.Title)

	util.DrawText(screen, "Host:", 50, 100, shared.CLR_WHITE, util.Font.Normal)
	util.DrawText(screen, "Port:", 50, 160, shared.CLR_WHITE, util.Font.Normal)
	util.DrawText(screen, "Password:", 50, 220, shared.CLR_WHITE, util.Font.Normal)

	hostRectColor := shared.CLR_WHITE
	portRectColor := shared.CLR_WHITE
	passwordRectColor := shared.CLR_WHITE

	switch lv.activeField {
	case 0:
		hostRectColor = shared.CLR_SELECTED
	case 1:
		portRectColor = shared.CLR_SELECTED
	case 2:
		passwordRectColor = shared.CLR_SELECTED
	}

	util.DrawScaledRect(screen, 180, 80, 300, 30, hostRectColor)
	util.DrawScaledRect(screen, 180, 140, 300, 30, portRectColor)
	util.DrawScaledRect(screen, 180, 200, 300, 30, passwordRectColor)

	util.DrawText(screen, lv.hostInput, 185, 100, shared.CLR_BLACK, util.Font.Normal)
	util.DrawText(screen, lv.portInput, 185, 160, shared.CLR_BLACK, util.Font.Normal)
	util.DrawText(screen, lv.passwordInput, 185, 220, shared.CLR_BLACK, util.Font.Normal)

	if lv.errorMessage != "" {
		util.DrawText(screen, lv.errorMessage, 50, 280, color.RGBA{255, 0, 0, 255}, util.Font.Normal)
	}

	util.DrawText(screen, "Press Enter to confirm, Tab to switch fields", 50, 340, color.Gray{Y: 200}, util.Font.Normal)
}
