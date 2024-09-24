// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flat

import "strconv"

type GameEventOption byte

const (
	GameEventOptionDefault GameEventOption = 0
	GameEventOptionHaunted GameEventOption = 1
	GameEventOptionRugby   GameEventOption = 2
)

var EnumNamesGameEventOption = map[GameEventOption]string{
	GameEventOptionDefault: "Default",
	GameEventOptionHaunted: "Haunted",
	GameEventOptionRugby:   "Rugby",
}

var EnumValuesGameEventOption = map[string]GameEventOption{
	"Default": GameEventOptionDefault,
	"Haunted": GameEventOptionHaunted,
	"Rugby":   GameEventOptionRugby,
}

func (v GameEventOption) String() string {
	if s, ok := EnumNamesGameEventOption[v]; ok {
		return s
	}
	return "GameEventOption(" + strconv.FormatInt(int64(v), 10) + ")"
}