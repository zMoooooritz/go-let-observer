package main

import (
	"image/color"
	"sort"
	"strings"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/zMoooooritz/go-let-loose/pkg/hll"
)

func (g *Game) updateMap() error {
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		return ebiten.Termination
	}

	if time.Since(g.mapView.lastUpdateTime) >= RCON_FETCH_INTERVAL {
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

	g.drawHeader(screen)

	// Draw the scoreboard if active
	if g.mapView.showScoreboard {
		g.drawScoreboard(screen)
	} else if g.mapView.selectedPlayerID != "" {
		if player, ok := g.mapView.players[g.mapView.selectedPlayerID]; ok {
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
	sortedPlayers := []hll.DetailedPlayerInfo{}
	for _, player := range g.mapView.players {
		sortedPlayers = append(sortedPlayers, player)
	}
	sort.Slice(sortedPlayers, func(i, j int) bool {
		return sortedPlayers[i].Name > sortedPlayers[j].Name
	})
	for _, player := range sortedPlayers {
		if !player.IsSpawned() {
			continue
		}

		x, y := translateCoords(g.dim.sizeX, g.dim.sizeY, player.Position)
		x = x*g.mapView.zoomLevel + g.mapView.panX
		y = y*g.mapView.zoomLevel + g.mapView.panY

		clr := BLUE
		if player.Team == hll.TmAxis {
			clr = RED
		}
		if player.ID == g.mapView.selectedPlayerID {
			clr = GREEN
		}

		// Draw the base circle
		vector.DrawFilledCircle(screen, float32(x), float32(y), float32(playerCircleRadius(g.mapView.zoomLevel)), clr, false)

		// Overlay the role icon
		roleImage, ok := g.mapView.roleImages[strings.ToLower(string(player.Role))]
		if ok {
			targetSize := playerIconSize(g.mapView.zoomLevel)
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
