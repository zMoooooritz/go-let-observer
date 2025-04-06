package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/zMoooooritz/go-let-loose/pkg/hll"
	"golang.org/x/image/font"
)

func clamp(value, min, max float64) float64 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

func drawText(screen *ebiten.Image, txt string, x, y int, clr color.Color, face font.Face) {
	bounds, _ := font.BoundString(face, txt)
	height := (bounds.Max.Y - bounds.Min.Y).Ceil()
	text.Draw(screen, txt, face, x, y+height/2, clr)
}

func drawTextNoShift(screen *ebiten.Image, txt string, x, y int, clr color.Color, face font.Face) {
	text.Draw(screen, txt, face, x, y, clr)
}

func translateCoords(sizeX, sizeY int, coords hll.Position) (float64, float64) {
	// Map X from [-100000, 100000] to [0, 1000]
	screenX := (float64(coords.X) + 100000) * float64(sizeX) / 200000
	// Map Y from [-100000, 100000] to [0, 1000]
	screenY := (float64(coords.Y) + 100000) * float64(sizeY) / 200000

	return screenX, screenY
}

func playerCircleRadius(zoomLevel float64) float64 {
	return 8 + 0.5*zoomLevel
}

func playerIconSize(zoomLevel float64) float64 {
	return 10 + 0.5*zoomLevel
}
