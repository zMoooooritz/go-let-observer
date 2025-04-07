package main

import (
	"fmt"
	"image/color"
	"sort"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/zMoooooritz/go-let-loose/pkg/hll"
)

func (g *Game) drawScoreboard(screen *ebiten.Image) {
	// Define scoreboard dimensions
	scoreboardWidth := 800
	scoreboardHeight := 500
	screenWidth := screen.Bounds().Dx()
	screenHeight := screen.Bounds().Dy()
	scoreboardX := (screenWidth - scoreboardWidth) / 2
	scoreboardY := (screenHeight - scoreboardHeight) / 2

	// Draw the scoreboard background
	backgroundColor := color.RGBA{0, 0, 0, 200}
	vector.DrawFilledRect(screen, float32(scoreboardX), float32(scoreboardY), float32(scoreboardWidth), float32(scoreboardHeight), backgroundColor, false)

	// Display scoreboard title
	textX := scoreboardX + 20
	textY := scoreboardY + 20
	lineHeight := 30
	drawText(screen, "Scoreboard (Top 25 Players)", textX, textY, color.White, g.fnt.title)
	textY += lineHeight

	sortedPlayers := []hll.DetailedPlayerInfo{}
	for _, player := range g.mapView.players {
		sortedPlayers = append(sortedPlayers, player)
	}
	sort.Slice(sortedPlayers, func(i, j int) bool {
		return sortedPlayers[i].Score.Combat > sortedPlayers[j].Score.Combat ||
			sortedPlayers[i].Score.Combat == sortedPlayers[j].Score.Combat &&
				sortedPlayers[i].Name > sortedPlayers[j].Name
	})

	drawText(screen, formatScoreboardLine("Name", "Kills", "Deaths", "K/D", "Team", "Cbt", "Off", "Def", "Sup"), textX, textY, color.White, g.fnt.normal)
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
		playerInfo := formatScoreboardLine(player.Name, strconv.Itoa(player.Kills), strconv.Itoa(player.Deaths), kdStr, string(player.Team),
			strconv.Itoa(player.Score.Combat), strconv.Itoa(player.Score.Offense), strconv.Itoa(player.Score.Defense), strconv.Itoa(player.Score.Support))
		drawText(screen, playerInfo, textX, textY, color.White, g.fnt.normal)
		textY += 15
	}
}

func formatScoreboardLine(name, kills, deaths, kd, team, combat, offense, defense, support string) string {
	return fmt.Sprintf("%10s %30s %7s %7s %5s %5s %5s %5s %5s", team, name, kills, deaths, kd, combat, offense, defense, support)
}
