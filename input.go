package main

import "github.com/hajimehoshi/ebiten/v2"

func (g *Game) handleMouseInput() {
	// Handle zooming with mouse wheel
	mouseX, mouseY := ebiten.CursorPosition()
	_, wheelY := ebiten.Wheel()
	if wheelY != 0 {
		oldZoom := g.mapView.zoomLevel
		g.mapView.zoomLevel += float64(wheelY * ZOOM_STEP_MULTIPLIER)
		if g.mapView.zoomLevel < MIN_ZOOM_LEVEL {
			g.mapView.zoomLevel = MIN_ZOOM_LEVEL
		} else if g.mapView.zoomLevel > MAX_ZOOM_LEVEL {
			g.mapView.zoomLevel = MAX_ZOOM_LEVEL
		}

		// Adjust pan to zoom into the mouse pointer location
		mouseWorldX := (float64(mouseX) - g.mapView.panX) / oldZoom
		mouseWorldY := (float64(mouseY) - g.mapView.panY) / oldZoom
		g.mapView.panX -= mouseWorldX * (g.mapView.zoomLevel - oldZoom)
		g.mapView.panY -= mouseWorldY * (g.mapView.zoomLevel - oldZoom)
	}

	// Handle panning with mouse drag
	if g.mapView.zoomLevel == MIN_ZOOM_LEVEL {
		g.mapView.panX = 0
		g.mapView.panY = 0
	} else {
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
			x, y := ebiten.CursorPosition()
			if g.mapView.isDragging {
				g.mapView.panX += float64(x - g.mapView.lastMouseX)
				g.mapView.panY += float64(y - g.mapView.lastMouseY)
			}
			g.mapView.lastMouseX = x
			g.mapView.lastMouseY = y
			g.mapView.isDragging = true
		} else {
			g.mapView.isDragging = false
		}

		g.mapView.panX = clamp(g.mapView.panX, float64(g.dim.sizeX)*(MIN_ZOOM_LEVEL-g.mapView.zoomLevel), 0)
		g.mapView.panY = clamp(g.mapView.panY, float64(g.dim.sizeY)*(MIN_ZOOM_LEVEL-g.mapView.zoomLevel), 0)
	}

	// Detect clicks on player circles
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		foundPlayer := false
		mouseX, mouseY := ebiten.CursorPosition()
		for _, player := range g.mapView.players {
			if !player.IsSpawned() {
				continue
			}

			x, y := translateCoords(g.dim.sizeX, g.dim.sizeY, player.Position)
			x = x*g.mapView.zoomLevel + g.mapView.panX
			y = y*g.mapView.zoomLevel + g.mapView.panY

			// Check if the mouse click is within the circle
			radius := playerCircleRadius(g.mapView.zoomLevel)
			if (float64(mouseX)-x)*(float64(mouseX)-x)+(float64(mouseY)-y)*(float64(mouseY)-y) <= radius*radius {
				g.mapView.selectedPlayerID = player.ID
				foundPlayer = true
				break
			}
		}
		if !foundPlayer {
			g.mapView.selectedPlayerID = ""
		}
	}
}

func (g *Game) handleKeyboardInput() {
	// Toggle scoreboard when holding Tab
	if ebiten.IsKeyPressed(ebiten.KeyTab) {
		g.mapView.showScoreboard = true
		g.mapView.selectedPlayerID = ""
	} else {
		g.mapView.showScoreboard = false
	}

	// Toggle grid visibility when 'G' is pressed
	if ebiten.IsKeyPressed(ebiten.KeyG) {
		if !g.mapView.gridKeyPressed {
			g.mapView.showGrid = !g.mapView.showGrid
			g.mapView.gridKeyPressed = true
		}
	} else {
		g.mapView.gridKeyPressed = false
	}
}
