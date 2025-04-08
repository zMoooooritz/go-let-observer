package game

import (
	"log"
	"sort"
	"time"

	"github.com/zMoooooritz/go-let-loose/pkg/hll"
	"github.com/zMoooooritz/go-let-observer/assets"
	"github.com/zMoooooritz/go-let-observer/pkg/util"
)

func (g *Game) fetchRconData() {
	players, err := g.rcon.GetPlayersInfo()
	if err != nil {
		return
	}

	playerMap := map[string]hll.DetailedPlayerInfo{}
	for _, player := range players {
		playerMap[player.ID] = player
	}
	playerList := players
	sort.Slice(playerList, func(i, j int) bool {
		return playerList[i].ID > playerList[j].ID
	})

	g.mapView.playerMap = playerMap
	g.mapView.playerList = playerList

	currMap, err := g.rcon.GetCurrentMap()
	if err != nil {
		return
	}

	currMapName := assets.ToFileName(currMap.ID)
	g.mapView.currentMapOrientation = currMap.Orientation
	if currMapName != g.mapView.currentMapName {
		img, err := util.LoadMapImage(currMapName)
		if err == nil {
			g.backgroundImage = img
		} else {
			log.Println("Error loading background image:", err)
		}
	}

	sessionInfo, err := g.rcon.GetSessionInfo()
	if err != nil {
		return
	}
	g.mapView.serverName = sessionInfo.ServerName
	g.mapView.playerCurrCount = sessionInfo.PlayerCount
	g.mapView.playerMaxCount = sessionInfo.MaxPlayerCount

	g.mapView.lastUpdateTime = time.Now()
	g.mapView.initialDataLoaded = true
}
