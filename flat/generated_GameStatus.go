// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flat

import "strconv"

type GameStatus byte

const (
	/// Game has not been created yet
	GameStatusInactive   GameStatus = 0
	/// 3-2-1 countdown
	GameStatusCountdown  GameStatus = 1
	/// After countdown, but before ball has been hit
	GameStatusKickoff    GameStatus = 2
	/// Ball has been hit
	GameStatusActive     GameStatus = 3
	/// A goal was scored. Waiting for replay.
	GameStatusGoalScored GameStatus = 4
	/// Watching replay
	GameStatusReplay     GameStatus = 5
	/// Game paused
	GameStatusPaused     GameStatus = 6
	/// Match has ended
	GameStatusEnded      GameStatus = 7
)

var EnumNamesGameStatus = map[GameStatus]string{
	GameStatusInactive:   "Inactive",
	GameStatusCountdown:  "Countdown",
	GameStatusKickoff:    "Kickoff",
	GameStatusActive:     "Active",
	GameStatusGoalScored: "GoalScored",
	GameStatusReplay:     "Replay",
	GameStatusPaused:     "Paused",
	GameStatusEnded:      "Ended",
}

var EnumValuesGameStatus = map[string]GameStatus{
	"Inactive":   GameStatusInactive,
	"Countdown":  GameStatusCountdown,
	"Kickoff":    GameStatusKickoff,
	"Active":     GameStatusActive,
	"GoalScored": GameStatusGoalScored,
	"Replay":     GameStatusReplay,
	"Paused":     GameStatusPaused,
	"Ended":      GameStatusEnded,
}

func (v GameStatus) String() string {
	if s, ok := EnumNamesGameStatus[v]; ok {
		return s
	}
	return "GameStatus(" + strconv.FormatInt(int64(v), 10) + ")"
}