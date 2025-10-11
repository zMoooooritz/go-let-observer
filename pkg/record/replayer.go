package record

import (
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/zMoooooritz/go-let-loose/pkg/hll"
	"github.com/zMoooooritz/go-let-observer/pkg/rcndata"
	"google.golang.org/protobuf/proto"
)

type MatchReplayer struct {
	Header           *MatchHeader
	currentTimeStamp time.Time
	isPaused         bool
	snapshots        []*Snapshot
}

func NewMatchReplayer(filePath string) (*MatchReplayer, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	dataBytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	matchData := &MatchData{}
	if err := proto.Unmarshal(dataBytes, matchData); err != nil {
		return nil, err
	}

	mr := &MatchReplayer{
		Header:    matchData.Header,
		isPaused:  true,
		snapshots: matchData.Snapshots,
	}

	if mr.Header.Version != VERSION {
		fmt.Println("Warning: Version mismatch. Expected:", VERSION, "but got:", mr.Header.Version, "this may cause issues.")
	}

	mr.currentTimeStamp = mr.firstTimestamp()
	return mr, nil
}

func (r *MatchReplayer) FetchRconDataSnapshot() (*rcndata.RconDataSnapshot, error) {
	state, snapshot, err := r.getStateAndSnapshotAt(r.currentTimeStamp)
	if err != nil {
		return nil, err
	}
	playerStateMap := make(map[int]*PlayerState)
	for _, player := range state.Players {
		playerStateMap[int(player.PlayerId)] = player
	}

	playerMap := make(map[string]hll.DetailedPlayerInfo)
	players := []hll.DetailedPlayerInfo{}
	for _, player := range r.Header.Players {
		playerState, ok := playerStateMap[int(player.RecordId)]
		if !ok {
			continue
		}

		unitID := int(playerState.Unit)
		unitName := hll.UnitIDToName(unitID)
		detailedPlayer := hll.DetailedPlayerInfo{
			PlayerInfo: hll.PlayerInfo{
				ID:   player.Id,
				Name: player.Name,
			},
			ClanTag: player.ClanTag,
			Level:   int(player.Level),
			Position: hll.Position{
				X: float64(playerState.X),
				Y: float64(playerState.Y),
				Z: 0.0,
			},
			Kills:  int(playerState.Kills),
			Deaths: int(playerState.Deaths),
			Unit:   hll.Unit{ID: unitID, Name: unitName},
			Team:   hll.TeamFromInt(int(playerState.Team)),
			Role:   hll.RoleFromInt(int(playerState.Role)),
		}
		playerMap[player.Id] = detailedPlayer
		players = append(players, detailedPlayer)
	}

	return &rcndata.RconDataSnapshot{
		Players:    players,
		PlayerMap:  playerMap,
		CurrentMap: hll.MapToGameMap(hll.Map(r.Header.MapId)),
		SessionInfo: hll.SessionInfo{
			AlliedScore:        int(snapshot.AlliedScore),
			AxisScore:          int(snapshot.AxisScore),
			RemainingMatchTime: snapshot.RemainingTime.AsDuration(),
		},
		FetchTime: time.Now(),
	}, nil
}

func (r *MatchReplayer) StartCurrentEndTime() (time.Time, time.Time, time.Time) {
	return r.Header.StartTime.AsTime(), r.currentTimeStamp, r.Header.EndTime.AsTime()
}

func (r *MatchReplayer) IsUserSeekable() bool { return true }

func (r *MatchReplayer) IsPaused() bool { return r.isPaused }

func (r *MatchReplayer) Pause() { r.isPaused = true }

func (r *MatchReplayer) Continue() { r.isPaused = false }

func (r *MatchReplayer) Seek(duration time.Duration) {
	r.currentTimeStamp = r.currentTimeStamp.Add(duration)
	if r.currentTimeStamp.Before(r.firstTimestamp()) {
		r.currentTimeStamp = r.firstTimestamp()
	}
	if r.currentTimeStamp.After(r.Header.EndTime.AsTime()) {
		r.currentTimeStamp = r.Header.EndTime.AsTime()
	}
}

func (r *MatchReplayer) firstTimestamp() time.Time {
	if len(r.snapshots) > 0 {
		return r.snapshots[0].Timestamp.AsTime()
	}
	return r.Header.StartTime.AsTime()
}

func (r *MatchReplayer) getStateAt(timestamp time.Time) (*FullSnapshot, error) {
	state, _, err := r.getStateAndSnapshotAt(timestamp)
	return state, err
}

func (r *MatchReplayer) getStateAndSnapshotAt(timestamp time.Time) (*FullSnapshot, *Snapshot, error) {
	var baseSnapshot *FullSnapshot
	var currentSnapshot *Snapshot
	state := &FullSnapshot{}

	for _, snapshot := range r.snapshots {
		if snapshot.Timestamp.AsTime().After(timestamp) {
			break
		}
		currentSnapshot = snapshot

		switch s := snapshot.Data.(type) {
		case *Snapshot_FullSnapshot:
			baseSnapshot = s.FullSnapshot
			state = proto.Clone(baseSnapshot).(*FullSnapshot)
		case *Snapshot_DeltaSnapshot:
			if baseSnapshot != nil {
				applyDelta(state, s.DeltaSnapshot)
			}
		}
	}

	if baseSnapshot == nil {
		return nil, nil, errors.New("no full snapshot found before the given timestamp")
	}

	return state, currentSnapshot, nil
}

func applyDelta(state *FullSnapshot, delta *DeltaSnapshot) {
	playerMap := make(map[int32]*PlayerState)
	for _, player := range state.Players {
		playerMap[player.PlayerId] = player
	}

	for _, deltaPlayer := range delta.Players {
		if player, exists := playerMap[deltaPlayer.PlayerId]; exists {
			if deltaPlayer.X != 0 || deltaPlayer.Y != 0 {
				player.X += deltaPlayer.X
				player.Y += deltaPlayer.Y
			}
			if deltaPlayer.Kills != nil {
				player.Kills += *deltaPlayer.Kills
			}
			if deltaPlayer.Deaths != nil {
				player.Deaths += *deltaPlayer.Deaths
			}
			if deltaPlayer.Team != nil {
				player.Team = *deltaPlayer.Team
			}
			if deltaPlayer.Unit != nil {
				player.Unit = *deltaPlayer.Unit
			}
			if deltaPlayer.Role != nil {
				player.Role = *deltaPlayer.Role
			}
		}
	}
}
