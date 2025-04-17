package ui

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/zMoooooritz/go-let-loose/pkg/hll"
	"github.com/zMoooooritz/go-let-observer/pkg/util"
)

func drawScoreboard(screen *ebiten.Image, playerList []hll.DetailedPlayerInfo) {
	scoreboardWidth := 800
	scoreboardHeight := 500
	screenWidth := ROOT_SCALING_SIZE
	screenHeight := ROOT_SCALING_SIZE
	scoreboardX := (screenWidth - scoreboardWidth) / 2
	scoreboardY := (screenHeight - scoreboardHeight) / 2

	util.DrawScaledRect(screen, scoreboardX, scoreboardY, scoreboardWidth, scoreboardHeight, CLR_OVERLAY)

	textX := scoreboardX + 20
	textY := scoreboardY + 40
	lineHeight := 30
	util.DrawText(screen, "Scoreboard (Top 25 Players)", textX, textY, CLR_WHITE, util.Font.Normal)
	textY += lineHeight

	sortedPlayers := playerList
	sort.Slice(sortedPlayers, func(i, j int) bool {
		return sortedPlayers[i].Score.Combat > sortedPlayers[j].Score.Combat // TODO: sort by kills when data is present in data recv from server
	})

	util.DrawText(screen, formatScoreboardLine("Rank", "Name", "Kills", "Deaths", "K/D", "Lvl", "Cbt", "Off", "Def", "Sup"), textX, textY, CLR_WHITE, util.Font.Small)
	textY += 25
	for i, player := range sortedPlayers {
		if i >= 25 {
			break
		}
		kdStr := "0.0"
		if player.Deaths > 0 {
			kdStr = fmt.Sprintf("%.2f", float32(player.Kills)/float32(player.Deaths))
		} else {
			kdStr = fmt.Sprintf("%.2f", float32(player.Kills))
		}
		playerInfo := formatScoreboardLine(strconv.Itoa(i+1), player.Name, strconv.Itoa(player.Kills), strconv.Itoa(player.Deaths), kdStr, strconv.Itoa(player.Level),
			strconv.Itoa(player.Score.Combat), strconv.Itoa(player.Score.Offense), strconv.Itoa(player.Score.Defense), strconv.Itoa(player.Score.Support))
		util.DrawText(screen, playerInfo, textX, textY, CLR_WHITE, util.Font.Small)
		textY += 15
	}
}

func formatScoreboardLine(rank, name, kills, deaths, kd, level, combat, offense, defense, support string) string {
	return fmt.Sprintf("%4s %5s  %-30s %7s %7s %5s %5s %5s %5s %5s", rank, level, name, kills, deaths, kd, combat, offense, defense, support)
}
