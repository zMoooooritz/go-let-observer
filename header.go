package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
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

	// Draw the overlay background
	overlayColor := color.RGBA{0, 0, 0, 200}
	vector.DrawFilledRect(screen, float32(overlayX), float32(overlayY), float32(overlayWidth), float32(overlayHeight), overlayColor, false)

	textX := overlayX + 10
	textY := overlayY + 20
	drawText(screen, g.serverName, textX, textY, color.White, g.fnt.title)
}

func (g *Game) drawPlayerCount(screen *ebiten.Image) {
	overlayWidth := 200
	overlayHeight := 50
	overlayX := 0
	overlayY := 50

	// Draw the overlay background
	overlayColor := color.RGBA{0, 0, 0, 200}
	vector.DrawFilledRect(screen, float32(overlayX), float32(overlayY), float32(overlayWidth), float32(overlayHeight), overlayColor, false)

	textX := overlayX + 10
	textY := overlayY + 20
	info := fmt.Sprintf("Players: %d/%d", g.playerCurrCount, g.playerMaxCount)
	drawText(screen, info, textX, textY, color.White, g.fnt.title)
}
