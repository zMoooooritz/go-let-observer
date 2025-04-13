package rcndata

import (
	"time"

	"github.com/zMoooooritz/go-let-loose/pkg/hll"
	"github.com/zMoooooritz/go-let-loose/pkg/rconv2"
)

type RconDataSnapshot struct {
	Players     []hll.DetailedPlayerInfo
	CurrentMap  hll.GameMap
	SessionInfo hll.SessionInfo
	FetchTime   time.Time
}

func FetchRconDataSnapshot(rcon *rconv2.Rcon) (*RconDataSnapshot, error) {
	players, err := rcon.GetPlayersInfo()
	if err != nil {
		return nil, err
	}

	currMap, err := rcon.GetCurrentMap()
	if err != nil {
		return nil, err
	}

	sessionInfo, err := rcon.GetSessionInfo()
	if err != nil {
		return nil, err
	}

	return &RconDataSnapshot{
		Players:     players,
		CurrentMap:  currMap,
		SessionInfo: sessionInfo,
		FetchTime:   time.Now(),
	}, nil
}
