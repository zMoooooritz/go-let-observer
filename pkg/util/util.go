package util

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/zMoooooritz/go-let-loose/pkg/hll"
	"golang.org/x/image/font"
)

// Generic Clamp function for int and float
func Clamp[T int | float64](value, min, max T) T {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

func DrawText(screen *ebiten.Image, txt string, x, y int, clr color.Color, face font.Face, scaleFactor float32) {
	bounds, _ := font.BoundString(face, txt)
	height := (bounds.Max.Y - bounds.Min.Y).Ceil()
	x = int(float32(x) * scaleFactor)
	y = int(float32(y) * scaleFactor)
	y += height / 2
	text.Draw(screen, txt, face, x, y, clr)
}

func DrawTextNoShift(screen *ebiten.Image, txt string, x, y int, clr color.Color, face font.Face, scaleFactor float32) {
	x = int(float32(x) * scaleFactor)
	y = int(float32(y) * scaleFactor)
	text.Draw(screen, txt, face, x, y, clr)
}

func DrawScaledRect(screen *ebiten.Image, x, y, width, height int, color color.Color, scaleFactor float32) {
	scaledX := float32(x) * scaleFactor
	scaledY := float32(y) * scaleFactor
	scaledWidth := float32(width) * scaleFactor
	scaledHeight := float32(height) * scaleFactor
	vector.DrawFilledRect(screen, scaledX, scaledY, scaledWidth, scaledHeight, color, false)
}

func ScaledDim(val int, scaleFactor float32) int {
	return int(float32(val) * scaleFactor)
}

func TranslateCoords(sizeX, sizeY int, coords hll.Position) (float64, float64) {
	// Map X from [-100000, 100000] to [0, 1000]
	screenX := (float64(coords.X) + 100000) * float64(sizeX) / 200000
	// Map Y from [-100000, 100000] to [0, 1000]
	screenY := (float64(coords.Y) + 100000) * float64(sizeY) / 200000

	return screenX, screenY
}

func PlayerCircleRadius(zoomLevel float64) float64 {
	return float64(8 + 0.5*zoomLevel)
}

func PlayerIconSize(zoomLevel float64) float64 {
	return 10 + 0.5*zoomLevel
}
