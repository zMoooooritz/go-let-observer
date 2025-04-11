package game

import (
	"sync"
	"time"

	"github.com/zMoooooritz/go-let-loose/pkg/hll"
)

type SpawnType string

const (
	SpawnTypeGarrison SpawnType = "garrison"
	SpawnTypeOutpost  SpawnType = "outpost"
	SpawnTypeNone     SpawnType = "unknown"
)

const (
	SPAWN_DISTANCE_DELTA          = 3000 // Distance in centimeters (30m)
	OUTPOST_DESTRUCTION_DISTANCE  = 1000 // Distance in centimeters (10m)
	GARRISON_DESTRUCTION_DISTANCE = 500  // Distance in centimeters (5m)
)

var spawnTTL = map[SpawnType]time.Duration{
	SpawnTypeGarrison: 5 * time.Minute,
	SpawnTypeOutpost:  3 * time.Minute,
}

type SpawnPoint struct {
	position   hll.Position
	team       hll.Team
	spawnType  SpawnType
	lastSeen   time.Time
	spawnCount int
	unit       string
	usedBy     map[string]bool
}

type SpawnTracker struct {
	spawns []SpawnPoint
	mu     sync.Mutex
}

func NewSpawnTracker() *SpawnTracker {
	return &SpawnTracker{
		spawns: []SpawnPoint{},
	}
}

func (st *SpawnTracker) TrackPlayerPosition(previousState, currentState hll.DetailedPlayerInfo) {
	st.mu.Lock()
	defer st.mu.Unlock()

	if hasJustSpawned(previousState, currentState) {
		st.handlePlayerSpawn(currentState)
	}
	st.destroyNearbySpawns(currentState)
}

func (st *SpawnTracker) ResetSpawns() {
	st.mu.Lock()
	defer st.mu.Unlock()

	st.spawns = []SpawnPoint{}
}

func (st *SpawnTracker) CleanExpiredSpawns() {
	st.mu.Lock()
	defer st.mu.Unlock()

	active := []SpawnPoint{}
	for _, spawn := range st.spawns {
		if ttl, ok := spawnTTL[spawn.spawnType]; ok {
			if time.Since(spawn.lastSeen) < ttl {
				active = append(active, spawn)
			}
		}
	}
	st.spawns = active
}

func (st *SpawnTracker) GetSpawns() []SpawnPoint {
	st.mu.Lock()
	defer st.mu.Unlock()
	return st.spawns
}

func hasJustSpawned(previous, current hll.DetailedPlayerInfo) bool {
	return previous.ID == current.ID && !previous.IsSpawned() && current.IsSpawned()
}

func (st *SpawnTracker) destroyNearbySpawns(player hll.DetailedPlayerInfo) {
	for i := len(st.spawns) - 1; i >= 0; i-- {
		spawn := &st.spawns[i]
		if spawn.team != player.Team {
			if spawn.spawnType == SpawnTypeOutpost && player.PlanarDistanceTo(spawn.position) <= OUTPOST_DESTRUCTION_DISTANCE {
				st.spawns = append(st.spawns[:i], st.spawns[i+1:]...)
			}
			if spawn.spawnType == SpawnTypeGarrison && player.PlanarDistanceTo(spawn.position) <= GARRISON_DESTRUCTION_DISTANCE {
				st.spawns = append(st.spawns[:i], st.spawns[i+1:]...)
			}
		}
	}

}

func (st *SpawnTracker) handlePlayerSpawn(player hll.DetailedPlayerInfo) {
	index, isNearExisting := st.isNearExistingSpawn(player)

	if isNearExisting {
		st.updateSpawnPoint(index, player)
	} else {
		st.addNewSpawnPoint(player)
	}

	st.analyzeSpawnTypes()
}

func (st *SpawnTracker) isNearExistingSpawn(player hll.DetailedPlayerInfo) (int, bool) {
	for i, spawn := range st.spawns {
		if spawn.team == player.Team {
			if player.PlanarDistanceTo(spawn.position) <= SPAWN_DISTANCE_DELTA {
				return i, true
			}
		}
	}
	return -1, false
}

func (st *SpawnTracker) updateSpawnPoint(index int, player hll.DetailedPlayerInfo) {
	spawn := &st.spawns[index]
	spawn.lastSeen = time.Now()
	spawn.spawnCount += 1

	spawn.position.X = (spawn.position.X*float64(spawn.spawnCount-1) + player.Position.X) / float64(spawn.spawnCount)
	spawn.position.Y = (spawn.position.Y*float64(spawn.spawnCount-1) + player.Position.Y) / float64(spawn.spawnCount)
	spawn.position.Z = (spawn.position.Z*float64(spawn.spawnCount-1) + player.Position.Z) / float64(spawn.spawnCount)

	if spawn.usedBy == nil {
		spawn.usedBy = make(map[string]bool)
	}
	spawn.usedBy[player.Unit.Name] = true
}

func (st *SpawnTracker) addNewSpawnPoint(player hll.DetailedPlayerInfo) {
	usedBy := make(map[string]bool)
	usedBy[player.Unit.Name] = true

	spawn := SpawnPoint{
		position:   player.Position,
		team:       player.Team,
		spawnType:  SpawnTypeNone,
		lastSeen:   time.Now(),
		spawnCount: 1,
		unit:       player.Unit.Name,
		usedBy:     usedBy,
	}

	st.spawns = append(st.spawns, spawn)
}

func (st *SpawnTracker) analyzeSpawnTypes() {
	for i := range st.spawns {
		spawn := &st.spawns[i]

		if len(spawn.usedBy) > 1 {
			spawn.spawnType = SpawnTypeGarrison
		} else {
			spawn.spawnType = SpawnTypeOutpost
		}
	}
}
