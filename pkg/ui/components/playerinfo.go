package components

import (
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/zMoooooritz/go-let-loose/pkg/hll"
	"github.com/zMoooooritz/go-let-observer/pkg/ui/shared"
	"github.com/zMoooooritz/go-let-observer/pkg/util"
)

var playerInfoCache struct {
	image      *ebiten.Image
	lastPlayer hll.DetailedPlayerInfo
}

func DrawPlayerInfoOverlay(screen *ebiten.Image, player hll.DetailedPlayerInfo) {
	overlayWidth := 250
	overlayHeight := shared.ROOT_SCALING_SIZE
	overlayX := shared.ROOT_SCALING_SIZE - overlayWidth

	if playerInfoCache.image == nil || player != playerInfoCache.lastPlayer {
		playerInfoCache.image = util.NewScaledImage(overlayWidth, overlayHeight)
		playerInfoCache.lastPlayer = player

		util.DrawScaledRect(playerInfoCache.image, 0, 0, overlayWidth, overlayHeight, shared.CLR_OVERLAY)

		textX := 10
		textY := 30
		lineHeight := 20
		dividerHeight := 30

		util.DrawText(playerInfoCache.image, "Player Info", textX, textY, shared.CLR_WHITE, util.Font.Normal)
		textY += dividerHeight
		textX += 10
		util.DrawText(playerInfoCache.image, "Name: "+player.Name, textX, textY, shared.CLR_WHITE, util.Font.Small)
		textY += lineHeight
		util.DrawText(playerInfoCache.image, "Clantag: "+player.ClanTag, textX, textY, shared.CLR_WHITE, util.Font.Small)
		textY += lineHeight
		util.DrawText(playerInfoCache.image, "Level: "+strconv.Itoa(player.Level), textX, textY, shared.CLR_WHITE, util.Font.Small)
		textY += dividerHeight
		util.DrawText(playerInfoCache.image, "Team: "+string(player.Team), textX, textY, shared.CLR_WHITE, util.Font.Small)
		textY += lineHeight
		util.DrawText(playerInfoCache.image, "Unit: "+string(player.Unit.Name), textX, textY, shared.CLR_WHITE, util.Font.Small)
		textY += lineHeight
		util.DrawText(playerInfoCache.image, "Role: "+string(player.Role), textX, textY, shared.CLR_WHITE, util.Font.Small)
		textY += lineHeight
		util.DrawText(playerInfoCache.image, "Loadout: "+player.Loadout, textX, textY, shared.CLR_WHITE, util.Font.Small)
		textY += dividerHeight
		util.DrawText(playerInfoCache.image, "Kills: "+strconv.Itoa(player.Kills), textX, textY, shared.CLR_WHITE, util.Font.Small)
		textY += lineHeight
		util.DrawText(playerInfoCache.image, "Deaths: "+strconv.Itoa(player.Deaths), textX, textY, shared.CLR_WHITE, util.Font.Small)
		textY += lineHeight
		util.DrawText(playerInfoCache.image, "Score:", textX, textY, shared.CLR_WHITE, util.Font.Small)
		textY += lineHeight
		util.DrawText(playerInfoCache.image, "Combat : "+strconv.Itoa(player.Score.Combat), textX+10, textY, shared.CLR_WHITE, util.Font.Small)
		textY += lineHeight
		util.DrawText(playerInfoCache.image, "Offense: "+strconv.Itoa(player.Score.Offense), textX+10, textY, shared.CLR_WHITE, util.Font.Small)
		textY += lineHeight
		util.DrawText(playerInfoCache.image, "Defense: "+strconv.Itoa(player.Score.Defense), textX+10, textY, shared.CLR_WHITE, util.Font.Small)
		textY += lineHeight
		util.DrawText(playerInfoCache.image, "Support: "+strconv.Itoa(player.Score.Support), textX+10, textY, shared.CLR_WHITE, util.Font.Small)
	}

	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(util.ScaledDim(overlayX)), 0)
	screen.DrawImage(playerInfoCache.image, options)
}
