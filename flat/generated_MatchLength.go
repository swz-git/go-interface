// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flat

import "strconv"

type MatchLength byte

const (
	MatchLengthFive_Minutes   MatchLength = 0
	MatchLengthTen_Minutes    MatchLength = 1
	MatchLengthTwenty_Minutes MatchLength = 2
	MatchLengthUnlimited      MatchLength = 3
)

var EnumNamesMatchLength = map[MatchLength]string{
	MatchLengthFive_Minutes:   "Five_Minutes",
	MatchLengthTen_Minutes:    "Ten_Minutes",
	MatchLengthTwenty_Minutes: "Twenty_Minutes",
	MatchLengthUnlimited:      "Unlimited",
}

var EnumValuesMatchLength = map[string]MatchLength{
	"Five_Minutes":   MatchLengthFive_Minutes,
	"Ten_Minutes":    MatchLengthTen_Minutes,
	"Twenty_Minutes": MatchLengthTwenty_Minutes,
	"Unlimited":      MatchLengthUnlimited,
}

func (v MatchLength) String() string {
	if s, ok := EnumNamesMatchLength[v]; ok {
		return s
	}
	return "MatchLength(" + strconv.FormatInt(int64(v), 10) + ")"
}