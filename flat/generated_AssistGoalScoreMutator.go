// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flat

import "strconv"

type AssistGoalScoreMutator byte

const (
	AssistGoalScoreMutatorZero  AssistGoalScoreMutator = 0
	AssistGoalScoreMutatorOne   AssistGoalScoreMutator = 1
	AssistGoalScoreMutatorTwo   AssistGoalScoreMutator = 2
	AssistGoalScoreMutatorThree AssistGoalScoreMutator = 3
)

var EnumNamesAssistGoalScoreMutator = map[AssistGoalScoreMutator]string{
	AssistGoalScoreMutatorZero:  "Zero",
	AssistGoalScoreMutatorOne:   "One",
	AssistGoalScoreMutatorTwo:   "Two",
	AssistGoalScoreMutatorThree: "Three",
}

var EnumValuesAssistGoalScoreMutator = map[string]AssistGoalScoreMutator{
	"Zero":  AssistGoalScoreMutatorZero,
	"One":   AssistGoalScoreMutatorOne,
	"Two":   AssistGoalScoreMutatorTwo,
	"Three": AssistGoalScoreMutatorThree,
}

func (v AssistGoalScoreMutator) String() string {
	if s, ok := EnumNamesAssistGoalScoreMutator[v]; ok {
		return s
	}
	return "AssistGoalScoreMutator(" + strconv.FormatInt(int64(v), 10) + ")"
}
