package game

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/zMoooooritz/go-let-observer/pkg/util"
)

func (g *Game) drawHeader(screen *ebiten.Image) {
	g.drawServerName(screen)
	g.drawPlayerCount(screen)
}

func (g *Game) drawServerName(screen *ebiten.Image) {
	overlayWidth := 750
	overlayHeight := 50
	overlayX := 0
	overlayY := 0

	util.DrawScaledRect(screen, overlayX, overlayY, overlayWidth, overlayHeight, CLR_OVERLAY)

	textX := overlayX + 10
	textY := overlayY + 30
	util.DrawText(screen, g.mapView.serverName, textX, textY, CLR_WHITE, g.fnt.Title)
}

func (g *Game) drawPlayerCount(screen *ebiten.Image) {
	overlayWidth := 200
	overlayHeight := 50
	overlayX := 0
	overlayY := 50

	util.DrawScaledRect(screen, overlayX, overlayY, overlayWidth, overlayHeight, CLR_OVERLAY)

	textX := overlayX + 10
	textY := overlayY + 30
	info := fmt.Sprintf("Players: %d/%d", g.mapView.playerCurrCount, g.mapView.playerMaxCount)
	util.DrawText(screen, info, textX, textY, CLR_WHITE, g.fnt.Title)
}
