package main

import (
	"log"
	"time"

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
		g.mapView.players = playerMap
	}
	currMap, err := g.rcon.GetCurrentMap()
	if err == nil {
		currMapName := assets.ToFileName(currMap.ID)
		g.mapView.currentMapOrientation = currMap.Orientation
		if currMapName != g.mapView.currentMapName {
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
		g.mapView.serverName = sessionInfo.ServerName
		g.mapView.playerCurrCount = sessionInfo.PlayerCount
		g.mapView.playerMaxCount = sessionInfo.MaxPlayerCount
	}

	g.mapView.lastUpdateTime = time.Now()
	g.mapView.initialDataLoaded = true
}
