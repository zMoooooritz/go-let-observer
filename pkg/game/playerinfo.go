package game

import (
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/zMoooooritz/go-let-loose/pkg/hll"
	"github.com/zMoooooritz/go-let-observer/pkg/util"
)

func (g *Game) drawPlayerOverlay(screen *ebiten.Image, player hll.DetailedPlayerInfo) {
	overlayWidth := 250
	overlayHeight := ROOT_SCALING_SIZE
	overlayX := ROOT_SCALING_SIZE - overlayWidth

	util.DrawScaledRect(screen, overlayX, 0, overlayWidth, overlayHeight, CLR_OVERLAY)

	// Display player information
	textX := overlayX + 10
	textY := 80
	lineHeight := 20
	dividerHeight := 30

	util.DrawText(screen, "Player Info", textX, textY, CLR_WHITE, g.fnt.Normal)
	textY += dividerHeight
	textX += 10
	util.DrawText(screen, "Name: "+player.Name, textX, textY, CLR_WHITE, g.fnt.Small)
	textY += lineHeight
	util.DrawText(screen, "Clantag: "+player.ClanTag, textX, textY, CLR_WHITE, g.fnt.Small)
	textY += lineHeight
	util.DrawText(screen, "Level: "+strconv.Itoa(player.Level), textX, textY, CLR_WHITE, g.fnt.Small)
	textY += dividerHeight
	util.DrawText(screen, "Team: "+string(player.Team), textX, textY, CLR_WHITE, g.fnt.Small)
	textY += lineHeight
	util.DrawText(screen, "Unit: "+string(player.Unit.Name), textX, textY, CLR_WHITE, g.fnt.Small)
	textY += lineHeight
	util.DrawText(screen, "Role: "+string(player.Role), textX, textY, CLR_WHITE, g.fnt.Small)
	textY += lineHeight
	util.DrawText(screen, "Loadout: "+player.Loadout, textX, textY, CLR_WHITE, g.fnt.Small)
	textY += dividerHeight
	util.DrawText(screen, "Kills: "+strconv.Itoa(player.Kills), textX, textY, CLR_WHITE, g.fnt.Small)
	textY += lineHeight
	util.DrawText(screen, "Deaths: "+strconv.Itoa(player.Deaths), textX, textY, CLR_WHITE, g.fnt.Small)
	textY += lineHeight
	util.DrawText(screen, "Score:", textX, textY, CLR_WHITE, g.fnt.Small)
	textY += lineHeight
	util.DrawText(screen, "Combat : "+strconv.Itoa(player.Score.Combat), textX+10, textY, CLR_WHITE, g.fnt.Small)
	textY += lineHeight
	util.DrawText(screen, "Offense: "+strconv.Itoa(player.Score.Offense), textX+10, textY, CLR_WHITE, g.fnt.Small)
	textY += lineHeight
	util.DrawText(screen, "Defense: "+strconv.Itoa(player.Score.Defense), textX+10, textY, CLR_WHITE, g.fnt.Small)
	textY += lineHeight
	util.DrawText(screen, "Support: "+strconv.Itoa(player.Score.Support), textX+10, textY, CLR_WHITE, g.fnt.Small)
	textY += lineHeight
}
