// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flat

import "strconv"

/// Game speed mutator options.
type GameSpeedMutator byte

const (
	GameSpeedMutatorDefault  GameSpeedMutator = 0
	GameSpeedMutatorSloMo    GameSpeedMutator = 1
	GameSpeedMutatorTimeWarp GameSpeedMutator = 2
)

var EnumNamesGameSpeedMutator = map[GameSpeedMutator]string{
	GameSpeedMutatorDefault:  "Default",
	GameSpeedMutatorSloMo:    "SloMo",
	GameSpeedMutatorTimeWarp: "TimeWarp",
}

var EnumValuesGameSpeedMutator = map[string]GameSpeedMutator{
	"Default":  GameSpeedMutatorDefault,
	"SloMo":    GameSpeedMutatorSloMo,
	"TimeWarp": GameSpeedMutatorTimeWarp,
}

func (v GameSpeedMutator) String() string {
	if s, ok := EnumNamesGameSpeedMutator[v]; ok {
		return s
	}
	return "GameSpeedMutator(" + strconv.FormatInt(int64(v), 10) + ")"
}
