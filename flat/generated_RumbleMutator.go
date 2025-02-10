// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flat

import "strconv"

/// Rumble mutator options.
type RumbleMutator byte

const (
	RumbleMutatorNoRumble         RumbleMutator = 0
	RumbleMutatorDefaultRumble    RumbleMutator = 1
	RumbleMutatorSlow             RumbleMutator = 2
	RumbleMutatorCivilized        RumbleMutator = 3
	RumbleMutatorDestructionDerby RumbleMutator = 4
	RumbleMutatorSpringLoaded     RumbleMutator = 5
	RumbleMutatorSpikesOnly       RumbleMutator = 6
	RumbleMutatorSpikeRush        RumbleMutator = 7
	RumbleMutatorHauntedBallBeam  RumbleMutator = 8
	RumbleMutatorTactical         RumbleMutator = 9
	RumbleMutatorBatmanRumble     RumbleMutator = 10
)

var EnumNamesRumbleMutator = map[RumbleMutator]string{
	RumbleMutatorNoRumble:         "NoRumble",
	RumbleMutatorDefaultRumble:    "DefaultRumble",
	RumbleMutatorSlow:             "Slow",
	RumbleMutatorCivilized:        "Civilized",
	RumbleMutatorDestructionDerby: "DestructionDerby",
	RumbleMutatorSpringLoaded:     "SpringLoaded",
	RumbleMutatorSpikesOnly:       "SpikesOnly",
	RumbleMutatorSpikeRush:        "SpikeRush",
	RumbleMutatorHauntedBallBeam:  "HauntedBallBeam",
	RumbleMutatorTactical:         "Tactical",
	RumbleMutatorBatmanRumble:     "BatmanRumble",
}

var EnumValuesRumbleMutator = map[string]RumbleMutator{
	"NoRumble":         RumbleMutatorNoRumble,
	"DefaultRumble":    RumbleMutatorDefaultRumble,
	"Slow":             RumbleMutatorSlow,
	"Civilized":        RumbleMutatorCivilized,
	"DestructionDerby": RumbleMutatorDestructionDerby,
	"SpringLoaded":     RumbleMutatorSpringLoaded,
	"SpikesOnly":       RumbleMutatorSpikesOnly,
	"SpikeRush":        RumbleMutatorSpikeRush,
	"HauntedBallBeam":  RumbleMutatorHauntedBallBeam,
	"Tactical":         RumbleMutatorTactical,
	"BatmanRumble":     RumbleMutatorBatmanRumble,
}

func (v RumbleMutator) String() string {
	if s, ok := EnumNamesRumbleMutator[v]; ok {
		return s
	}
	return "RumbleMutator(" + strconv.FormatInt(int64(v), 10) + ")"
}
