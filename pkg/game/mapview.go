package game

import (
	"image/color"
	"strings"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/zMoooooritz/go-let-loose/pkg/hll"
	"github.com/zMoooooritz/go-let-observer/pkg/util"
)

func (g *Game) updateMap() error {
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		return ebiten.Termination
	}

	if time.Since(g.mapView.lastUpdateTime) >= fetchIntervalSteps[g.mapView.intervalIndex] {
		g.mapView.fetchMutex.Lock()
		if !g.mapView.isFetching {
			g.mapView.isFetching = true
			go func() {
				g.fetchRconData()
				g.mapView.fetchMutex.Lock()
				g.mapView.isFetching = false
				g.mapView.fetchMutex.Unlock()
			}()
		}
		g.mapView.fetchMutex.Unlock()
	}

	g.handleKeyboardInput()
	g.handleMouseInput()

	return nil
}

func (g *Game) drawMap(screen *ebiten.Image) {
	g.drawBackground(screen)

	if !g.mapView.initialDataLoaded {
		return
	}

	g.drawMapView(screen)

	if g.mapView.showHeader {
		g.drawHeader(screen)
	}

	if g.mapView.showScoreboard {
		g.drawScoreboard(screen)
	} else if g.mapView.selectedPlayerID != "" {
		if player, ok := g.mapView.playerMap[g.mapView.selectedPlayerID]; ok {
			g.drawPlayerOverlay(screen, player)
		}
	}
}

func (g *Game) drawBackground(screen *ebiten.Image) {
	if g.backgroundImage != nil {
		screenSize := screen.Bounds().Size()
		imageSize := g.backgroundImage.Bounds().Size()
		scale := (float64(screenSize.X) / float64(imageSize.X)) * g.mapView.zoomLevel

		options := &ebiten.DrawImageOptions{}
		options.GeoM.Scale(scale, scale)
		options.GeoM.Translate(g.mapView.panX, g.mapView.panY)
		screen.DrawImage(g.backgroundImage, options)
	}
}

func (g *Game) drawMapView(screen *ebiten.Image) {
	screenSize := screen.Bounds().Size()

	if g.mapView.showGrid {
		g.drawGrid(screen, screenSize.X, screenSize.Y)
	}

	g.drawPlayers(screen)
}

func (g *Game) drawPlayers(screen *ebiten.Image) {
	for _, player := range g.mapView.playerList {
		if !player.IsSpawned() {
			continue
		}

		x, y := util.TranslateCoords(g.dim.sizeX, g.dim.sizeY, player.Position)
		x = x*g.mapView.zoomLevel + g.mapView.panX
		y = y*g.mapView.zoomLevel + g.mapView.panY

		clr := CLR_ALLIES
		if player.Team == hll.TmAxis {
			clr = CLR_AXIS
		}
		if player.ID == g.mapView.selectedPlayerID {
			clr = CLR_SELECTED
		}

		vector.DrawFilledCircle(screen, float32(x), float32(y), float32(util.PlayerCircleRadius(g.mapView.zoomLevel)), clr, false)

		roleImage, ok := g.mapView.roleImages[strings.ToLower(string(player.Role))]
		if ok {
			targetSize := util.PlayerIconSize(g.mapView.zoomLevel)
			iconScale := targetSize / float64(roleImage.Bounds().Dx())

			options := &ebiten.DrawImageOptions{}
			options.GeoM.Scale(iconScale, iconScale)
			options.GeoM.Translate(x-targetSize/2, y-targetSize/2)
			screen.DrawImage(roleImage, options)
		}

	}
}

func (g *Game) drawGrid(screen *ebiten.Image, width, height int) {
	bgWidth := float64(g.dim.sizeX) * g.mapView.zoomLevel
	bgHeight := float64(g.dim.sizeY) * g.mapView.zoomLevel

	cellWidth := bgWidth / 5
	cellHeight := bgHeight / 5

	active := []int{1, 2, 1, 0, 0}

	gridColor := color.RGBA{100, 100, 100, 255}
	// fillColor := color.RGBA{50, 50, 50, 100}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			x := float64(i)*cellWidth + g.mapView.panX
			y := float64(j)*cellHeight + g.mapView.panY

			if g.mapView.currentMapOrientation == hll.OriHorizontal {
				if j == 0 || j == 4 {
					continue
				}

				if active[i]+1 == j {
					// vector.DrawFilledRect(screen, float32(x), float32(y), float32(cellWidth), float32(cellHeight), fillColor, false)
				}
			}

			if g.mapView.currentMapOrientation == hll.OriVertical {
				if i == 0 || i == 4 {
					continue
				}

				if active[j]+1 == i {
					// vector.DrawFilledRect(screen, float32(x), float32(y), float32(cellWidth), float32(cellHeight), fillColor, false)
				}
			}

			vector.StrokeLine(screen, float32(x), float32(y), float32(x+cellWidth), float32(y), 3, gridColor, false)
			vector.StrokeLine(screen, float32(x), float32(y+cellHeight), float32(x+cellWidth), float32(y+cellHeight), 3, gridColor, false)
			vector.StrokeLine(screen, float32(x), float32(y), float32(x), float32(y+cellHeight), 3, gridColor, false)
			vector.StrokeLine(screen, float32(x+cellWidth), float32(y), float32(x+cellWidth), float32(y+cellHeight), 3, gridColor, false)
		}
	}
}

func (g *Game) handleMouseInput() {
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

		mouseWorldX := (float64(mouseX) - g.mapView.panX) / oldZoom
		mouseWorldY := (float64(mouseY) - g.mapView.panY) / oldZoom
		g.mapView.panX -= mouseWorldX * (g.mapView.zoomLevel - oldZoom)
		g.mapView.panY -= mouseWorldY * (g.mapView.zoomLevel - oldZoom)
	}

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

		g.mapView.panX = util.Clamp(g.mapView.panX, float64(g.dim.sizeX)*(MIN_ZOOM_LEVEL-g.mapView.zoomLevel), 0)
		g.mapView.panY = util.Clamp(g.mapView.panY, float64(g.dim.sizeY)*(MIN_ZOOM_LEVEL-g.mapView.zoomLevel), 0)
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		foundPlayer := false
		mouseX, mouseY := ebiten.CursorPosition()
		for _, player := range g.mapView.playerMap {
			if !player.IsSpawned() {
				continue
			}

			x, y := util.TranslateCoords(g.dim.sizeX, g.dim.sizeY, player.Position)
			x = x*g.mapView.zoomLevel + g.mapView.panX
			y = y*g.mapView.zoomLevel + g.mapView.panY

			radius := util.PlayerCircleRadius(g.mapView.zoomLevel)
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
	if ebiten.IsKeyPressed(ebiten.KeyTab) {
		g.mapView.showScoreboard = true
		g.mapView.selectedPlayerID = ""
	} else {
		g.mapView.showScoreboard = false
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyG) {
		g.mapView.showGrid = !g.mapView.showGrid
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyH) {
		g.mapView.showHeader = !g.mapView.showHeader
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyNumpadAdd) || inpututil.IsKeyJustPressed(ebiten.KeyRightBracket) {
		if g.mapView.intervalIndex < len(fetchIntervalSteps)-1 {
			g.mapView.intervalIndex += 1
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyNumpadSubtract) || inpututil.IsKeyJustPressed(ebiten.KeySlash) {
		if g.mapView.intervalIndex > 0 {
			g.mapView.intervalIndex -= 1
		}
	}
}
