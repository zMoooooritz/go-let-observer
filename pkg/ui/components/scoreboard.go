package components

import (
	"fmt"
	"sort"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/zMoooooritz/go-let-loose/pkg/hll"
	"github.com/zMoooooritz/go-let-observer/pkg/ui/shared"
	"github.com/zMoooooritz/go-let-observer/pkg/util"
)

var (
	cachedScoreboard     *ebiten.Image
	lastScoreboardUpdate time.Time
)

const (
	SCOREBOARD_WIDTH  = 800
	SCOREBOARD_HEIGHT = 400
)

func DrawScoreboard(screen *ebiten.Image, playerList []hll.DetailedPlayerInfo) {
	currentTime := time.Now()
	if cachedScoreboard == nil || currentTime.Sub(lastScoreboardUpdate) >= time.Second {
		cachedScoreboard = util.NewScaledImage(SCOREBOARD_WIDTH, SCOREBOARD_HEIGHT)

		util.DrawScaledRect(cachedScoreboard, 0, 0, SCOREBOARD_WIDTH, SCOREBOARD_HEIGHT, shared.CLR_OVERLAY)

		textX := 20
		textY := 30
		lineHeight := 30
		util.DrawText(cachedScoreboard, "Scoreboard (Top 20 Players)", textX, textY, shared.CLR_WHITE, util.Font.Normal)
		textY += lineHeight

		sortedPlayers := playerList
		sort.Slice(sortedPlayers, func(i, j int) bool {
			return sortedPlayers[i].Score.Combat > sortedPlayers[j].Score.Combat
		})

		util.DrawText(cachedScoreboard, formatScoreboardLine("Rank", "Name", "Kills", "Deaths", "K/D", "Lvl", "Cbt", "Off", "Def", "Sup"), textX, textY, shared.CLR_WHITE, util.Font.Small)
		textY += 25
		for i, player := range sortedPlayers {
			if i >= 20 {
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
			util.DrawText(cachedScoreboard, playerInfo, textX, textY, shared.CLR_WHITE, util.Font.Small)
			textY += 15
		}
		lastScoreboardUpdate = currentTime
	}

	screenWidth := shared.ROOT_SCALING_SIZE
	screenHeight := shared.ROOT_SCALING_SIZE
	scoreboardX := (screenWidth - SCOREBOARD_WIDTH) / 2
	scoreboardY := (screenHeight - SCOREBOARD_HEIGHT) / 2

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(util.ScaledDim(scoreboardX)), float64(util.ScaledDim(scoreboardY)))
	screen.DrawImage(cachedScoreboard, op)
}

func formatScoreboardLine(rank, name, kills, deaths, kd, level, combat, offense, defense, support string) string {
	return fmt.Sprintf("%4s %5s  %-30s %7s %7s %5s %5s %5s %5s %5s", rank, level, name, kills, deaths, kd, combat, offense, defense, support)
}
