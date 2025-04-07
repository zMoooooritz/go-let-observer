package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/zMoooooritz/go-let-loose/pkg/rconv2"
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
	// Clear the screen
	if g.backgroundImage != nil {
		screenSize := screen.Bounds().Size()
		imageSize := g.backgroundImage.Bounds().Size()
		scale := float64(screenSize.X) / float64(imageSize.X)

		options := &ebiten.DrawImageOptions{}
		options.GeoM.Scale(scale, scale)
		screen.DrawImage(g.backgroundImage, options)

		vector.DrawFilledRect(screen, 0, 0, 1000, 400, color.RGBA{0, 0, 0, 200}, false)
	} else {
		screen.Fill(color.RGBA{31, 31, 31, 255})
	}

	// Draw title
	drawTextNoShift(screen, "Login to HLL Observer", 20, 40, color.White, g.fnt.huge)

	// Draw input labels
	drawTextNoShift(screen, "Host:", 50, 100, color.White, g.fnt.title)
	drawTextNoShift(screen, "Port:", 50, 160, color.White, g.fnt.title)
	drawTextNoShift(screen, "Password:", 50, 220, color.White, g.fnt.title)

	// Draw rectangles around input fields
	hostRectColor := WHITE
	portRectColor := WHITE
	passwordRectColor := WHITE

	// Highlight the active field
	switch g.loginView.activeField {
	case 0:
		hostRectColor = GREEN
	case 1:
		portRectColor = GREEN
	case 2:
		passwordRectColor = GREEN
	}

	// Draw rectangles
	vector.DrawFilledRect(screen, 180, 80, 300, 30, hostRectColor, false)      // Host field
	vector.DrawFilledRect(screen, 180, 140, 300, 30, portRectColor, false)     // Port field
	vector.DrawFilledRect(screen, 180, 200, 300, 30, passwordRectColor, false) // Password field

	// Draw input fields
	drawTextNoShift(screen, g.loginView.hostInput, 185, 100, color.Black, g.fnt.title)
	drawTextNoShift(screen, g.loginView.portInput, 185, 160, color.Black, g.fnt.title)
	drawTextNoShift(screen, g.loginView.passwordInput, 185, 220, color.Black, g.fnt.title)

	// Draw error message if any
	if g.loginView.errorMessage != "" {
		drawTextNoShift(screen, g.loginView.errorMessage, 50, 280, color.RGBA{255, 0, 0, 255}, g.fnt.title)
	}

	// Draw instructions
	drawTextNoShift(screen, "Press Enter to confirm, Tab to switch fields", 50, 340, color.Gray{Y: 200}, g.fnt.title)
}
