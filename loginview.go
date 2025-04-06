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
		g.activeField = (g.activeField + 1) % 3
	}

	// Handle Enter key to attempt login
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		if g.hostInput != "" && g.portInput != "" && g.passwordInput != "" {
			cfg := rconv2.ServerConfig{
				Host:     g.hostInput,
				Port:     g.portInput,
				Password: g.passwordInput,
			}
			rcn, err := rconv2.NewRcon(cfg, 3)
			if err != nil {
				g.errorMessage = "Invalid credentials or connection error"
			} else {
				g.rcon = rcn
				g.stage = StageMap
			}
		} else {
			g.errorMessage = "All fields are required"
		}
	}

	// Handle text input
	if chars := ebiten.AppendInputChars(nil); len(chars) > 0 {
		for _, c := range chars {
			switch g.activeField {
			case 0:
				g.hostInput += string(c)
			case 1:
				g.portInput += string(c)
			case 2:
				g.passwordInput += string(c)
			}
		}
	}

	// Handle backspace
	if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) {
		switch g.activeField {
		case 0:
			if len(g.hostInput) > 0 {
				g.hostInput = g.hostInput[:len(g.hostInput)-1]
			}
		case 1:
			if len(g.portInput) > 0 {
				g.portInput = g.portInput[:len(g.portInput)-1]
			}
		case 2:
			if len(g.passwordInput) > 0 {
				g.passwordInput = g.passwordInput[:len(g.passwordInput)-1]
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
		scale := (float64(screenSize.X) / float64(imageSize.X)) * g.zoomLevel

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
	switch g.activeField {
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
	drawTextNoShift(screen, g.hostInput, 185, 100, color.Black, g.fnt.title)
	drawTextNoShift(screen, g.portInput, 185, 160, color.Black, g.fnt.title)
	drawTextNoShift(screen, g.passwordInput, 185, 220, color.Black, g.fnt.title)

	// Draw error message if any
	if g.errorMessage != "" {
		drawTextNoShift(screen, g.errorMessage, 50, 280, color.RGBA{255, 0, 0, 255}, g.fnt.title)
	}

	// Draw instructions
	drawTextNoShift(screen, "Press Enter to confirm, Tab to switch fields", 50, 340, color.Gray{Y: 200}, g.fnt.title)
}
