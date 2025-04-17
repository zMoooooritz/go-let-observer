package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/zMoooooritz/go-let-observer/pkg/util"
)

func drawProgressBar(screen *ebiten.Image, progress float64) {
	progressBarWidth := 600
	progressBarHeight := 30
	screenWidth := ROOT_SCALING_SIZE
	screenHeight := ROOT_SCALING_SIZE
	progressBarX := (screenWidth - progressBarWidth) / 2
	progressBarY := screenHeight - progressBarHeight - 20

	util.DrawScaledRect(screen, progressBarX, progressBarY, progressBarWidth, progressBarHeight, CLR_OVERLAY)

	filledWidth := int(progress * float64(progressBarWidth))

	util.DrawScaledRect(screen, progressBarX, progressBarY, filledWidth, progressBarHeight, CLR_WHITE)
}
