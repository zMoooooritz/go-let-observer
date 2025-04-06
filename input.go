package main

import "github.com/hajimehoshi/ebiten/v2"

func (g *Game) handleMouseInput() {
	// Handle zooming with mouse wheel
	mouseX, mouseY := ebiten.CursorPosition()
	_, wheelY := ebiten.Wheel()
	if wheelY != 0 {
		oldZoom := g.zoomLevel
		g.zoomLevel += float64(wheelY * ZOOM_STEP_MULTIPLIER)
		if g.zoomLevel < MIN_ZOOM_LEVEL {
			g.zoomLevel = MIN_ZOOM_LEVEL
		} else if g.zoomLevel > MAX_ZOOM_LEVEL {
			g.zoomLevel = MAX_ZOOM_LEVEL
		}

		// Adjust pan to zoom into the mouse pointer location
		mouseWorldX := (float64(mouseX) - g.panX) / oldZoom
		mouseWorldY := (float64(mouseY) - g.panY) / oldZoom
		g.panX -= mouseWorldX * (g.zoomLevel - oldZoom)
		g.panY -= mouseWorldY * (g.zoomLevel - oldZoom)
	}

	// Handle panning with mouse drag
	if g.zoomLevel == MIN_ZOOM_LEVEL {
		g.panX = 0
		g.panY = 0
	} else {
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
			x, y := ebiten.CursorPosition()
			if g.isDragging {
				g.panX += float64(x - g.lastMouseX)
				g.panY += float64(y - g.lastMouseY)
			}
			g.lastMouseX = x
			g.lastMouseY = y
			g.isDragging = true
		} else {
			g.isDragging = false
		}

		g.panX = clamp(g.panX, float64(g.sizeX)*(MIN_ZOOM_LEVEL-g.zoomLevel), 0)
		g.panY = clamp(g.panY, float64(g.sizeY)*(MIN_ZOOM_LEVEL-g.zoomLevel), 0)
	}

	// Detect clicks on player circles
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		foundPlayer := false
		mouseX, mouseY := ebiten.CursorPosition()
		for _, player := range g.players {
			if !player.IsSpawned() {
				continue
			}

			x, y := translateCoords(g.sizeX, g.sizeY, player.Position)
			x = x*g.zoomLevel + g.panX
			y = y*g.zoomLevel + g.panY

			// Check if the mouse click is within the circle
			radius := playerCircleRadius(g.zoomLevel)
			if (float64(mouseX)-x)*(float64(mouseX)-x)+(float64(mouseY)-y)*(float64(mouseY)-y) <= radius*radius {
				g.selectedPlayerID = player.ID
				foundPlayer = true
				break
			}
		}
		if !foundPlayer {
			g.selectedPlayerID = ""
		}
	}
}

func (g *Game) handleKeyboardInput() {
	// Toggle scoreboard when holding Tab
	if ebiten.IsKeyPressed(ebiten.KeyTab) {
		g.showScoreboard = true
		g.selectedPlayerID = ""
	} else {
		g.showScoreboard = false
	}

	// Toggle grid visibility when 'G' is pressed
	if ebiten.IsKeyPressed(ebiten.KeyG) {
		if !g.gridKeyPressed {
			g.showGrid = !g.showGrid
			g.gridKeyPressed = true
		}
	} else {
		g.gridKeyPressed = false
	}
}
