// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flat

import "strconv"

/// Max score mutator options.
type MaxScoreMutator byte

const (
	MaxScoreMutatorDefault    MaxScoreMutator = 0
	MaxScoreMutatorOneGoal    MaxScoreMutator = 1
	MaxScoreMutatorThreeGoals MaxScoreMutator = 2
	MaxScoreMutatorFiveGoals  MaxScoreMutator = 3
	MaxScoreMutatorSevenGoals MaxScoreMutator = 4
	MaxScoreMutatorUnlimited  MaxScoreMutator = 5
)

var EnumNamesMaxScoreMutator = map[MaxScoreMutator]string{
	MaxScoreMutatorDefault:    "Default",
	MaxScoreMutatorOneGoal:    "OneGoal",
	MaxScoreMutatorThreeGoals: "ThreeGoals",
	MaxScoreMutatorFiveGoals:  "FiveGoals",
	MaxScoreMutatorSevenGoals: "SevenGoals",
	MaxScoreMutatorUnlimited:  "Unlimited",
}

var EnumValuesMaxScoreMutator = map[string]MaxScoreMutator{
	"Default":    MaxScoreMutatorDefault,
	"OneGoal":    MaxScoreMutatorOneGoal,
	"ThreeGoals": MaxScoreMutatorThreeGoals,
	"FiveGoals":  MaxScoreMutatorFiveGoals,
	"SevenGoals": MaxScoreMutatorSevenGoals,
	"Unlimited":  MaxScoreMutatorUnlimited,
}

func (v MaxScoreMutator) String() string {
	if s, ok := EnumNamesMaxScoreMutator[v]; ok {
		return s
	}
	return "MaxScoreMutator(" + strconv.FormatInt(int64(v), 10) + ")"
}
