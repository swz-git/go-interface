// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flat

import "strconv"

/// Game event mutator options.
type GameEventMutator byte

const (
	GameEventMutatorDefault GameEventMutator = 0
	GameEventMutatorHaunted GameEventMutator = 1
	GameEventMutatorRugby   GameEventMutator = 2
)

var EnumNamesGameEventMutator = map[GameEventMutator]string{
	GameEventMutatorDefault: "Default",
	GameEventMutatorHaunted: "Haunted",
	GameEventMutatorRugby:   "Rugby",
}

var EnumValuesGameEventMutator = map[string]GameEventMutator{
	"Default": GameEventMutatorDefault,
	"Haunted": GameEventMutatorHaunted,
	"Rugby":   GameEventMutatorRugby,
}

func (v GameEventMutator) String() string {
	if s, ok := EnumNamesGameEventMutator[v]; ok {
		return s
	}
	return "GameEventMutator(" + strconv.FormatInt(int64(v), 10) + ")"
}
