package ui

import (
	"time"
	"unicode"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/zMoooooritz/go-let-observer/pkg/util"
)

func (mv *MapView) handleMouseInput() {
	mouseX, mouseY := ebiten.CursorPosition()
	_, wheelY := ebiten.Wheel()
	if wheelY != 0 {
		oldZoom := mv.dim.zoomLevel
		mv.dim.zoomLevel += float64(wheelY * ZOOM_STEP_MULTIPLIER)
		if mv.dim.zoomLevel < MIN_ZOOM_LEVEL {
			mv.dim.zoomLevel = MIN_ZOOM_LEVEL
		} else if mv.dim.zoomLevel > MAX_ZOOM_LEVEL {
			mv.dim.zoomLevel = MAX_ZOOM_LEVEL
		}

		mouseWorldX := (float64(mouseX) - mv.dim.panX) / oldZoom
		mouseWorldY := (float64(mouseY) - mv.dim.panY) / oldZoom
		mv.dim.panX -= mouseWorldX * (mv.dim.zoomLevel - oldZoom)
		mv.dim.panY -= mouseWorldY * (mv.dim.zoomLevel - oldZoom)
	}

	if mv.dim.zoomLevel == MIN_ZOOM_LEVEL {
		mv.dim.panX = 0
		mv.dim.panY = 0
	} else {
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
			x, y := ebiten.CursorPosition()
			if mv.isDragging {
				mv.dim.panX += float64(x - mv.lastMouseX)
				mv.dim.panY += float64(y - mv.lastMouseY)
			}
			mv.lastMouseX = x
			mv.lastMouseY = y
			mv.isDragging = true
		} else {
			mv.isDragging = false
		}

		mv.dim.panX = util.Clamp(mv.dim.panX, float64(mv.dim.sizeX)*(MIN_ZOOM_LEVEL-mv.dim.zoomLevel), 0)
		mv.dim.panY = util.Clamp(mv.dim.panY, float64(mv.dim.sizeY)*(MIN_ZOOM_LEVEL-mv.dim.zoomLevel), 0)
	}

	if mv.showPlayers && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		foundPlayer := false
		mouseX, mouseY := ebiten.CursorPosition()
		for _, player := range mv.playerMap {
			if !player.IsSpawned() {
				continue
			}

			x, y := util.TranslateCoords(mv.dim.sizeX, mv.dim.sizeY, player.Position)
			x = x*mv.dim.zoomLevel + mv.dim.panX
			y = y*mv.dim.zoomLevel + mv.dim.panY

			radius := util.IconCircleRadius(mv.dim.zoomLevel, PLAYER_SIZE_MODIFIER)
			if (float64(mouseX)-x)*(float64(mouseX)-x)+(float64(mouseY)-y)*(float64(mouseY)-y) <= radius*radius {
				mv.selectedPlayerID = player.ID
				foundPlayer = true
				break
			}
		}
		if !foundPlayer {
			mv.selectedPlayerID = ""
		}
	}
}

func (mv *MapView) handleKeyboardInput() {
	typed := ebiten.AppendInputChars(nil)
	for _, r := range typed {
		typedKey := string(unicode.ToLower(r))

		if typedKey == "g" {
			mv.showGrid = !mv.showGrid
		}

		if typedKey == "p" {
			mv.showPlayers = !mv.showPlayers
		}

		if typedKey == "i" {
			mv.showPlayerInfo = !mv.showPlayerInfo
		}

		if typedKey == "s" {
			mv.showSpawns = !mv.showSpawns
		}

		if typedKey == "t" {
			mv.showTanks = !mv.showTanks
		}

		if typedKey == "h" {
			mv.showServerInfo = !mv.showServerInfo
		}

		if typedKey == "+" {
			if mv.intervalIndex < len(fetchIntervalSteps)-1 {
				mv.intervalIndex++
			}
		}

		if typedKey == "-" {
			if mv.intervalIndex > 0 {
				mv.intervalIndex--
			}
		}

		if typedKey == "?" {
			mv.showHelp = !mv.showHelp
		}

	}

	mv.showScoreboard = false
	if ebiten.IsKeyPressed(ebiten.KeyTab) {
		mv.showScoreboard = true
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		if mv.dataFetcher.IsPaused() {
			mv.dataFetcher.Continue()
		} else {
			mv.dataFetcher.Pause()
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		duration := -time.Second
		if ebiten.IsKeyPressed(ebiten.KeyShift) {
			duration = -time.Minute
		}
		mv.dataFetcher.Seek(duration)
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		mv.dataFetcher.Seek(-2 * time.Hour)
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		duration := time.Second
		if ebiten.IsKeyPressed(ebiten.KeyShift) {
			duration = time.Minute
		}
		mv.dataFetcher.Seek(duration)
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		mv.dataFetcher.Seek(2 * time.Hour)
	}
}
