// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flat

import "strconv"

/// Gravity mutator options.
type GravityMutator byte

const (
	GravityMutatorDefault   GravityMutator = 0
	GravityMutatorLow       GravityMutator = 1
	GravityMutatorHigh      GravityMutator = 2
	GravityMutatorSuperHigh GravityMutator = 3
	GravityMutatorReverse   GravityMutator = 4
)

var EnumNamesGravityMutator = map[GravityMutator]string{
	GravityMutatorDefault:   "Default",
	GravityMutatorLow:       "Low",
	GravityMutatorHigh:      "High",
	GravityMutatorSuperHigh: "SuperHigh",
	GravityMutatorReverse:   "Reverse",
}

var EnumValuesGravityMutator = map[string]GravityMutator{
	"Default":   GravityMutatorDefault,
	"Low":       GravityMutatorLow,
	"High":      GravityMutatorHigh,
	"SuperHigh": GravityMutatorSuperHigh,
	"Reverse":   GravityMutatorReverse,
}

func (v GravityMutator) String() string {
	if s, ok := EnumNamesGravityMutator[v]; ok {
		return s
	}
	return "GravityMutator(" + strconv.FormatInt(int64(v), 10) + ")"
}
