// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flat

import "strconv"

/// Max time mutator options.
type MaxTimeMutator byte

const (
	MaxTimeMutatorDefault       MaxTimeMutator = 0
	MaxTimeMutatorElevenMinutes MaxTimeMutator = 1
)

var EnumNamesMaxTimeMutator = map[MaxTimeMutator]string{
	MaxTimeMutatorDefault:       "Default",
	MaxTimeMutatorElevenMinutes: "ElevenMinutes",
}

var EnumValuesMaxTimeMutator = map[string]MaxTimeMutator{
	"Default":       MaxTimeMutatorDefault,
	"ElevenMinutes": MaxTimeMutatorElevenMinutes,
}

func (v MaxTimeMutator) String() string {
	if s, ok := EnumNamesMaxTimeMutator[v]; ok {
		return s
	}
	return "MaxTimeMutator(" + strconv.FormatInt(int64(v), 10) + ")"
}
