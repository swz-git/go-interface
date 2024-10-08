// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flat

import "strconv"

type MaxScore byte

const (
	MaxScoreDefault     MaxScore = 0
	MaxScoreOne_Goal    MaxScore = 1
	MaxScoreThree_Goals MaxScore = 2
	MaxScoreFive_Goals  MaxScore = 3
	MaxScoreSeven_Goals MaxScore = 4
	MaxScoreUnlimited   MaxScore = 5
)

var EnumNamesMaxScore = map[MaxScore]string{
	MaxScoreDefault:     "Default",
	MaxScoreOne_Goal:    "One_Goal",
	MaxScoreThree_Goals: "Three_Goals",
	MaxScoreFive_Goals:  "Five_Goals",
	MaxScoreSeven_Goals: "Seven_Goals",
	MaxScoreUnlimited:   "Unlimited",
}

var EnumValuesMaxScore = map[string]MaxScore{
	"Default":     MaxScoreDefault,
	"One_Goal":    MaxScoreOne_Goal,
	"Three_Goals": MaxScoreThree_Goals,
	"Five_Goals":  MaxScoreFive_Goals,
	"Seven_Goals": MaxScoreSeven_Goals,
	"Unlimited":   MaxScoreUnlimited,
}

func (v MaxScore) String() string {
	if s, ok := EnumNamesMaxScore[v]; ok {
		return s
	}
	return "MaxScore(" + strconv.FormatInt(int64(v), 10) + ")"
}
