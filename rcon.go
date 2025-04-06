package main

import (
	"log"

	"github.com/zMoooooritz/go-let-loose/pkg/hll"
	"github.com/zMoooooritz/go-let-observer/assets"
)

func (g *Game) fetchRconData() {
	players, err := g.rcon.GetPlayersInfo()
	if err == nil {
		playerMap := map[string]hll.DetailedPlayerInfo{}
		for _, player := range players {
			playerMap[player.ID] = player
		}
		g.players = playerMap
	}
	currMap, err := g.rcon.GetCurrentMap()
	if err == nil {
		currMapName := assets.ToFileName(currMap.ID)
		g.currentMapOrientation = currMap.Orientation
		if currMapName != g.currentMapName {
			img, err := loadMapImage(currMapName)
			if err == nil {
				g.backgroundImage = img
			} else {
				log.Println("Error loading background image:", err)
			}
		}
	}

	sessionInfo, err := g.rcon.GetSessionInfo()
	if err == nil {
		g.serverName = sessionInfo.ServerName
		g.playerCurrCount = sessionInfo.PlayerCount
		g.playerMaxCount = sessionInfo.MaxPlayerCount
	}

}
