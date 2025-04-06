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

	if time.Since(g.lastUpdateTime) >= RCON_FETCH_INTERVAL {
		g.fetchMutex.Lock()
		if !g.isFetching {
			g.isFetching = true
			go func() {
				g.fetchRconData()
				g.fetchMutex.Lock()
				g.isFetching = false
				g.fetchMutex.Unlock()
			}()
		}
		g.fetchMutex.Unlock()

		g.lastUpdateTime = time.Now()
		g.initialDataLoaded = true
	}

	g.handleKeyboardInput()
	g.handleMouseInput()

	return nil
}

func (g *Game) drawMap(screen *ebiten.Image) {
	g.drawBackground(screen)

	if !g.initialDataLoaded {
		return
	}

	g.drawMapView(screen)

	g.drawHeader(screen)

	// Draw the scoreboard if active
	if g.showScoreboard {
		g.drawScoreboard(screen)
	} else if g.selectedPlayerID != "" {
		if player, ok := g.players[g.selectedPlayerID]; ok {
			g.drawPlayerOverlay(screen, player)
		}
	}
}

func (g *Game) drawBackground(screen *ebiten.Image) {
	if g.backgroundImage != nil {
		screenSize := screen.Bounds().Size()
		imageSize := g.backgroundImage.Bounds().Size()
		scale := (float64(screenSize.X) / float64(imageSize.X)) * g.zoomLevel

		options := &ebiten.DrawImageOptions{}
		options.GeoM.Scale(scale, scale)
		options.GeoM.Translate(g.panX, g.panY)
		screen.DrawImage(g.backgroundImage, options)
	}
}

func (g *Game) drawMapView(screen *ebiten.Image) {
	screenSize := screen.Bounds().Size()

	if g.showGrid {
		g.drawGrid(screen, screenSize.X, screenSize.Y)
	}

	g.drawPlayers(screen)
}

func (g *Game) drawPlayers(screen *ebiten.Image) {
	sortedPlayers := []hll.DetailedPlayerInfo{}
	for _, player := range g.players {
		sortedPlayers = append(sortedPlayers, player)
	}
	sort.Slice(sortedPlayers, func(i, j int) bool {
		return sortedPlayers[i].Name > sortedPlayers[j].Name
	})
	for _, player := range sortedPlayers {
		if !player.IsSpawned() {
			continue
		}

		x, y := translateCoords(g.sizeX, g.sizeY, player.Position)
		x = x*g.zoomLevel + g.panX
		y = y*g.zoomLevel + g.panY

		clr := BLUE
		if player.Team == hll.TmAxis {
			clr = RED
		}
		if player.ID == g.selectedPlayerID {
			clr = GREEN
		}

		// Draw the base circle
		vector.DrawFilledCircle(screen, float32(x), float32(y), float32(playerCircleRadius(g.zoomLevel)), clr, false)

		// Overlay the role icon
		roleImage, ok := g.roleImages[strings.ToLower(string(player.Role))]
		if ok {
			targetSize := playerIconSize(g.zoomLevel)
			iconScale := targetSize / float64(roleImage.Bounds().Dx())

			options := &ebiten.DrawImageOptions{}
			options.GeoM.Scale(iconScale, iconScale)
			options.GeoM.Translate(x-targetSize/2, y-targetSize/2)
			screen.DrawImage(roleImage, options)
		}

	}
}

func (g *Game) drawGrid(screen *ebiten.Image, width, height int) {
	bgWidth := float64(g.sizeX) * g.zoomLevel
	bgHeight := float64(g.sizeY) * g.zoomLevel

	cellWidth := bgWidth / 5
	cellHeight := bgHeight / 5

	active := []int{1, 2, 1, 0, 0}

	gridColor := color.RGBA{100, 100, 100, 255}
	// fillColor := color.RGBA{50, 50, 50, 100}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			x := float64(i)*cellWidth + g.panX
			y := float64(j)*cellHeight + g.panY

			if g.currentMapOrientation == hll.OriHorizontal {
				if j == 0 || j == 4 {
					continue
				}

				if active[i]+1 == j {
					// vector.DrawFilledRect(screen, float32(x), float32(y), float32(cellWidth), float32(cellHeight), fillColor, false)
				}
			}

			if g.currentMapOrientation == hll.OriVertical {
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
