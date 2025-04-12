package util

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/zMoooooritz/go-let-loose/pkg/hll"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

var (
	ScaleFactor float32
)

func Clamp[T int | float64](value, min, max T) T {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

func DrawText(screen *ebiten.Image, txt string, x, y int, clr color.Color, face font.Face) {
	x = int(float32(x) * ScaleFactor)
	y = int(float32(y) * ScaleFactor)

	d := &font.Drawer{
		Dst:  screen,
		Src:  image.NewUniform(clr),
		Face: face,
		Dot: fixed.Point26_6{
			X: fixed.I(x),
			Y: fixed.I(y),
		},
	}
	d.DrawString(txt)
}

func DrawScaledRect(screen *ebiten.Image, x, y, width, height int, color color.Color) {
	scaledX := float32(x) * ScaleFactor
	scaledY := float32(y) * ScaleFactor
	scaledWidth := float32(width) * ScaleFactor
	scaledHeight := float32(height) * ScaleFactor
	vector.DrawFilledRect(screen, scaledX, scaledY, scaledWidth, scaledHeight, color, false)
}

func ScaledDim(val int) int {
	return int(float32(val) * ScaleFactor)
}

func IconCircleRadius(zoomLevel float64, sizeModifier float64) float64 {
	return (8 + 0.5*zoomLevel) * float64(ScaleFactor) * sizeModifier
}

func IconSize(zoomLevel float64, sizeModifier float64) float64 {
	return (10 + 0.5*zoomLevel) * float64(ScaleFactor) * sizeModifier
}

func TranslateCoords(sizeX, sizeY int, coords hll.Position) (float64, float64) {
	// Map X from [-100000, 100000] to [0, 1000]
	screenX := (float64(coords.X) + 100000) * float64(sizeX) / 200000
	// Map Y from [-100000, 100000] to [0, 1000]
	screenY := (float64(coords.Y) + 100000) * float64(sizeY) / 200000

	return screenX, screenY
}
