// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flat

import "strconv"

/// Possible phases of the match.
type MatchPhase byte

const (
	/// Match has not been created yet.
	MatchPhaseInactive   MatchPhase = 0
	/// 3-2-1 countdown of a kickoff.
	MatchPhaseCountdown  MatchPhase = 1
	/// After kickoff countdown, but before ball has been hit.
	/// The match automatically proceeds to Active after 2 seconds.
	MatchPhaseKickoff    MatchPhase = 2
	/// The ball is in play and time is ticking.
	MatchPhaseActive     MatchPhase = 3
	/// A goal was just scored. Waiting for replay to start.
	MatchPhaseGoalScored MatchPhase = 4
	/// Goal replay is being shown.
	MatchPhaseReplay     MatchPhase = 5
	/// The match is paused.
	MatchPhasePaused     MatchPhase = 6
	/// The match has ended.
	MatchPhaseEnded      MatchPhase = 7
)

var EnumNamesMatchPhase = map[MatchPhase]string{
	MatchPhaseInactive:   "Inactive",
	MatchPhaseCountdown:  "Countdown",
	MatchPhaseKickoff:    "Kickoff",
	MatchPhaseActive:     "Active",
	MatchPhaseGoalScored: "GoalScored",
	MatchPhaseReplay:     "Replay",
	MatchPhasePaused:     "Paused",
	MatchPhaseEnded:      "Ended",
}

var EnumValuesMatchPhase = map[string]MatchPhase{
	"Inactive":   MatchPhaseInactive,
	"Countdown":  MatchPhaseCountdown,
	"Kickoff":    MatchPhaseKickoff,
	"Active":     MatchPhaseActive,
	"GoalScored": MatchPhaseGoalScored,
	"Replay":     MatchPhaseReplay,
	"Paused":     MatchPhasePaused,
	"Ended":      MatchPhaseEnded,
}

func (v MatchPhase) String() string {
	if s, ok := EnumNamesMatchPhase[v]; ok {
		return s
	}
	return "MatchPhase(" + strconv.FormatInt(int64(v), 10) + ")"
}
