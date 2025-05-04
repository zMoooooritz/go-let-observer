package components

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/zMoooooritz/go-let-observer/pkg/ui/shared"
	"github.com/zMoooooritz/go-let-observer/pkg/util"
)

func DrawProgressBar(screen *ebiten.Image, progress float64) {
	progressBarWidth := 600
	progressBarHeight := 30
	screenWidth := shared.ROOT_SCALING_SIZE
	screenHeight := shared.ROOT_SCALING_SIZE
	progressBarX := (screenWidth - progressBarWidth) / 2
	progressBarY := screenHeight - progressBarHeight - 20

	util.DrawScaledRect(screen, progressBarX, progressBarY, progressBarWidth, progressBarHeight, shared.CLR_OVERLAY)

	filledWidth := int(progress * float64(progressBarWidth))

	util.DrawScaledRect(screen, progressBarX, progressBarY, filledWidth, progressBarHeight, shared.CLR_WHITE)
}
