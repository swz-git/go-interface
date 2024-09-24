// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flat

import "strconv"

type MaxTimeOption byte

const (
	MaxTimeOptionDefault        MaxTimeOption = 0
	MaxTimeOptionEleven_Minutes MaxTimeOption = 1
)

var EnumNamesMaxTimeOption = map[MaxTimeOption]string{
	MaxTimeOptionDefault:        "Default",
	MaxTimeOptionEleven_Minutes: "Eleven_Minutes",
}

var EnumValuesMaxTimeOption = map[string]MaxTimeOption{
	"Default":        MaxTimeOptionDefault,
	"Eleven_Minutes": MaxTimeOptionEleven_Minutes,
}

func (v MaxTimeOption) String() string {
	if s, ok := EnumNamesMaxTimeOption[v]; ok {
		return s
	}
	return "MaxTimeOption(" + strconv.FormatInt(int64(v), 10) + ")"
}
