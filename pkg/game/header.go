package ui

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/zMoooooritz/go-let-observer/pkg/util"
)

func (mv *MapView) drawHeader(screen *ebiten.Image) {
	mv.drawServerName(screen)
	mv.drawPlayerCount(screen)
}

func (mv *MapView) drawServerName(screen *ebiten.Image) {
	overlayWidth := 750
	overlayHeight := 50
	overlayX := 0
	overlayY := 0

	util.DrawScaledRect(screen, overlayX, overlayY, overlayWidth, overlayHeight, CLR_OVERLAY)

	textX := overlayX + 10
	textY := overlayY + 30
	util.DrawText(screen, mv.serverName, textX, textY, CLR_WHITE, util.Font.Normal)
}

func (mv *MapView) drawPlayerCount(screen *ebiten.Image) {
	overlayWidth := 200
	overlayHeight := 50
	overlayX := 0
	overlayY := 50

	util.DrawScaledRect(screen, overlayX, overlayY, overlayWidth, overlayHeight, CLR_OVERLAY)

	textX := overlayX + 10
	textY := overlayY + 30
	info := fmt.Sprintf("Players: %d/%d", mv.playerCurrCount, mv.playerMaxCount)
	util.DrawText(screen, info, textX, textY, CLR_WHITE, util.Font.Normal)
}
