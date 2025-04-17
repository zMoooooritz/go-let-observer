package ui

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/zMoooooritz/go-let-observer/pkg/util"
)

func drawServerName(screen *ebiten.Image, serverName string) {
	overlayWidth := 750
	overlayHeight := 50
	overlayX := 0
	overlayY := 0

	util.DrawScaledRect(screen, overlayX, overlayY, overlayWidth, overlayHeight, CLR_OVERLAY)

	textX := overlayX + 10
	textY := overlayY + 30
	util.DrawText(screen, serverName, textX, textY, CLR_WHITE, util.Font.Normal)
}

func drawPlayerCount(screen *ebiten.Image, playerCurrCount, playerMaxCount int) {
	overlayWidth := 200
	overlayHeight := 50
	overlayX := 0
	overlayY := 50

	util.DrawScaledRect(screen, overlayX, overlayY, overlayWidth, overlayHeight, CLR_OVERLAY)

	textX := overlayX + 10
	textY := overlayY + 30
	info := fmt.Sprintf("Players: %d/%d", playerCurrCount, playerMaxCount)
	util.DrawText(screen, info, textX, textY, CLR_WHITE, util.Font.Normal)
}
