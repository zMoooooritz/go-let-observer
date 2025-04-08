package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/zMoooooritz/go-let-loose/pkg/rconv2"
	"github.com/zMoooooritz/go-let-observer/pkg/util"
)

func (g *Game) updateLogin() error {
	// Switch active input field with Tab
	if inpututil.IsKeyJustPressed(ebiten.KeyTab) {
		g.loginView.activeField = (g.loginView.activeField + 1) % 3
	}

	// Handle Enter key to attempt login
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		if g.loginView.hostInput != "" && g.loginView.portInput != "" && g.loginView.passwordInput != "" {
			cfg := rconv2.ServerConfig{
				Host:     g.loginView.hostInput,
				Port:     g.loginView.portInput,
				Password: g.loginView.passwordInput,
			}
			rcn, err := rconv2.NewRcon(cfg, 3)
			if err != nil {
				g.loginView.errorMessage = "Invalid credentials or connection error"
			} else {
				g.rcon = rcn
				g.uiState = StateMap
			}
		} else {
			g.loginView.errorMessage = "All fields are required"
		}
	}

	// Handle text input
	if chars := ebiten.AppendInputChars(nil); len(chars) > 0 {
		for _, c := range chars {
			switch g.loginView.activeField {
			case 0:
				g.loginView.hostInput += string(c)
			case 1:
				g.loginView.portInput += string(c)
			case 2:
				g.loginView.passwordInput += string(c)
			}
		}
	}

	// Handle backspace
	if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) {
		switch g.loginView.activeField {
		case 0:
			if len(g.loginView.hostInput) > 0 {
				g.loginView.hostInput = g.loginView.hostInput[:len(g.loginView.hostInput)-1]
			}
		case 1:
			if len(g.loginView.portInput) > 0 {
				g.loginView.portInput = g.loginView.portInput[:len(g.loginView.portInput)-1]
			}
		case 2:
			if len(g.loginView.passwordInput) > 0 {
				g.loginView.passwordInput = g.loginView.passwordInput[:len(g.loginView.passwordInput)-1]
			}
		}
	}

	return nil
}

func (g *Game) drawLogin(screen *ebiten.Image) {
	// Calculate scaling factor
	screenSize := screen.Bounds().Size()

	// Clear the screen
	if g.backgroundImage != nil {
		imageSize := g.backgroundImage.Bounds().Size()
		imageScale := float64(screenSize.X) / float64(imageSize.X)

		options := &ebiten.DrawImageOptions{}
		options.GeoM.Scale(imageScale, imageScale)
		screen.DrawImage(g.backgroundImage, options)
	} else {
		screen.Fill(color.RGBA{31, 31, 31, 255})
	}

	util.DrawScaledRect(screen, 0, 0, 1000, 400, CLR_OVERLAY, g.dim.scaleFactor)

	// Draw.Title
	util.DrawTextNoShift(screen, "Login to HLL Observer", 20, 40, CLR_WHITE, g.fnt.Huge, g.dim.scaleFactor)

	// Draw input labels
	util.DrawTextNoShift(screen, "Host:", 50, 100, CLR_WHITE, g.fnt.Title, g.dim.scaleFactor)
	util.DrawTextNoShift(screen, "Port:", 50, 160, CLR_WHITE, g.fnt.Title, g.dim.scaleFactor)
	util.DrawTextNoShift(screen, "Password:", 50, 220, CLR_WHITE, g.fnt.Title, g.dim.scaleFactor)

	// Draw rectangles around input fields
	hostRectColor := CLR_WHITE
	portRectColor := CLR_WHITE
	passwordRectColor := CLR_WHITE

	// Highlight the active field
	switch g.loginView.activeField {
	case 0:
		hostRectColor = CLR_SELECTED
	case 1:
		portRectColor = CLR_SELECTED
	case 2:
		passwordRectColor = CLR_SELECTED
	}

	// Draw rectangles
	util.DrawScaledRect(screen, 180, 80, 300, 30, hostRectColor, g.dim.scaleFactor)
	util.DrawScaledRect(screen, 180, 140, 300, 30, portRectColor, g.dim.scaleFactor)
	util.DrawScaledRect(screen, 180, 200, 300, 30, passwordRectColor, g.dim.scaleFactor)

	// Draw input fields
	util.DrawTextNoShift(screen, g.loginView.hostInput, 185, 100, CLR_BLACK, g.fnt.Title, g.dim.scaleFactor)
	util.DrawTextNoShift(screen, g.loginView.portInput, 185, 160, CLR_BLACK, g.fnt.Title, g.dim.scaleFactor)
	util.DrawTextNoShift(screen, g.loginView.passwordInput, 185, 220, CLR_BLACK, g.fnt.Title, g.dim.scaleFactor)

	// Draw error message if any
	if g.loginView.errorMessage != "" {
		util.DrawTextNoShift(screen, g.loginView.errorMessage, 50, 280, color.RGBA{255, 0, 0, 255}, g.fnt.Title, g.dim.scaleFactor)
	}

	// Draw instructions
	util.DrawTextNoShift(screen, "Press Enter to confirm, Tab to switch fields", 50, 340, color.Gray{Y: 200}, g.fnt.Title, g.dim.scaleFactor)
}
