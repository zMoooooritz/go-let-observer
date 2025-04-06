package main

import (
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/zMoooooritz/go-let-loose/pkg/hll"
)

func (g *Game) drawPlayerOverlay(screen *ebiten.Image, player hll.DetailedPlayerInfo) {
	overlayWidth := 250
	overlayHeight := screen.Bounds().Dy()
	overlayX := screen.Bounds().Dx() - overlayWidth

	// Draw the overlay background
	overlayColor := color.RGBA{0, 0, 0, 200}
	vector.DrawFilledRect(screen, float32(overlayX), 0, float32(overlayWidth), float32(overlayHeight), overlayColor, false)

	// Display player information
	textX := overlayX + 10
	textY := 60
	lineHeight := 20
	dividerHeight := 30

	drawText(screen, "Player Info", textX, textY, color.White, g.fnt.title)
	textY += dividerHeight
	textX += 10
	drawText(screen, "Name: "+player.Name, textX, textY, color.White, g.fnt.normal)
	textY += lineHeight
	drawText(screen, "ClanTag: "+player.ClanTag, textX, textY, color.White, g.fnt.normal)
	textY += lineHeight
	drawText(screen, "Level: "+strconv.Itoa(player.Level), textX, textY, color.White, g.fnt.normal)
	textY += dividerHeight
	drawText(screen, "Team: "+string(player.Team), textX, textY, color.White, g.fnt.normal)
	textY += lineHeight
	drawText(screen, "Unit: "+string(player.Unit.Name), textX, textY, color.White, g.fnt.normal)
	textY += lineHeight
	drawText(screen, "Role: "+string(player.Role), textX, textY, color.White, g.fnt.normal)
	textY += lineHeight
	drawText(screen, "Loadout: "+player.Loadout, textX, textY, color.White, g.fnt.normal)
	textY += dividerHeight
	drawText(screen, "Kills: "+strconv.Itoa(player.Kills), textX, textY, color.White, g.fnt.normal)
	textY += lineHeight
	drawText(screen, "Deaths: "+strconv.Itoa(player.Deaths), textX, textY, color.White, g.fnt.normal)
	textY += lineHeight
	drawText(screen, "Score:", textX, textY, color.White, g.fnt.normal)
	textY += lineHeight
	drawText(screen, "Combat : "+strconv.Itoa(player.Score.Combat), textX+10, textY, color.White, g.fnt.normal)
	textY += lineHeight
	drawText(screen, "Offense: "+strconv.Itoa(player.Score.Offense), textX+10, textY, color.White, g.fnt.normal)
	textY += lineHeight
	drawText(screen, "Defense: "+strconv.Itoa(player.Score.Defense), textX+10, textY, color.White, g.fnt.normal)
	textY += lineHeight
	drawText(screen, "Support: "+strconv.Itoa(player.Score.Support), textX+10, textY, color.White, g.fnt.normal)
	textY += lineHeight
}
