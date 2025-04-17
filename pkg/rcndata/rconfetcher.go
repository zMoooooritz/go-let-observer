package rcndata

import (
	"time"

	"github.com/zMoooooritz/go-let-loose/pkg/hll"
	"github.com/zMoooooritz/go-let-loose/pkg/rconv2"
)

type RconDataSnapshot struct {
	Players     []hll.DetailedPlayerInfo
	PlayerMap   map[string]hll.DetailedPlayerInfo
	CurrentMap  hll.GameMap
	SessionInfo hll.SessionInfo
	FetchTime   time.Time
}

type DataFetcher interface {
	FetchRconDataSnapshot() (*RconDataSnapshot, error)
	StartCurrentEndTime() (time.Time, time.Time, time.Time)
	IsUserSeekable() bool
	IsPaused() bool
	Pause()
	Continue()
	Seek(time.Duration)
}

type RconDataFetcher struct {
	rcon *rconv2.Rcon
}

func NewRconDataFetcher(rcon *rconv2.Rcon) *RconDataFetcher {
	return &RconDataFetcher{
		rcon: rcon,
	}
}

func (f *RconDataFetcher) FetchRconDataSnapshot() (*RconDataSnapshot, error) {
	players, err := f.rcon.GetPlayersInfo()
	if err != nil {
		return nil, err
	}
	playerMap := make(map[string]hll.DetailedPlayerInfo)
	for _, player := range players {
		playerMap[player.ID] = player
	}

	currMap, err := f.rcon.GetCurrentMap()
	if err != nil {
		return nil, err
	}

	sessionInfo, err := f.rcon.GetSessionInfo()
	if err != nil {
		return nil, err
	}

	return &RconDataSnapshot{
		Players:     players,
		PlayerMap:   playerMap,
		CurrentMap:  currMap,
		SessionInfo: sessionInfo,
		FetchTime:   time.Now(),
	}, nil
}

func (f *RconDataFetcher) StartCurrentEndTime() (time.Time, time.Time, time.Time) {
	return time.Time{}, time.Time{}, time.Time{}
}
func (f *RconDataFetcher) IsUserSeekable() bool { return false }
func (f *RconDataFetcher) IsPaused() bool       { return false }
func (f *RconDataFetcher) Pause()               {}
func (f *RconDataFetcher) Continue()            {}
func (f *RconDataFetcher) Seek(time.Duration)   {}
