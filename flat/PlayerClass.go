// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flat

import (
	flatbuffers "github.com/google/flatbuffers/go"
	"strconv"
)

type PlayerClass byte

const (
	PlayerClassNONE        PlayerClass = 0
	PlayerClassRLBot       PlayerClass = 1
	PlayerClassHuman       PlayerClass = 2
	PlayerClassPsyonix     PlayerClass = 3
	PlayerClassPartyMember PlayerClass = 4
)

var EnumNamesPlayerClass = map[PlayerClass]string{
	PlayerClassNONE:        "NONE",
	PlayerClassRLBot:       "RLBot",
	PlayerClassHuman:       "Human",
	PlayerClassPsyonix:     "Psyonix",
	PlayerClassPartyMember: "PartyMember",
}

var EnumValuesPlayerClass = map[string]PlayerClass{
	"NONE":        PlayerClassNONE,
	"RLBot":       PlayerClassRLBot,
	"Human":       PlayerClassHuman,
	"Psyonix":     PlayerClassPsyonix,
	"PartyMember": PlayerClassPartyMember,
}

func (v PlayerClass) String() string {
	if s, ok := EnumNamesPlayerClass[v]; ok {
		return s
	}
	return "PlayerClass(" + strconv.FormatInt(int64(v), 10) + ")"
}

type PlayerClassT struct {
	Type PlayerClass
	Value interface{}
}

func (t *PlayerClassT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	switch t.Type {
	case PlayerClassRLBot:
		return t.Value.(*RLBotT).Pack(builder)
	case PlayerClassHuman:
		return t.Value.(*HumanT).Pack(builder)
	case PlayerClassPsyonix:
		return t.Value.(*PsyonixT).Pack(builder)
	case PlayerClassPartyMember:
		return t.Value.(*PartyMemberT).Pack(builder)
	}
	return 0
}

func (rcv PlayerClass) UnPack(table flatbuffers.Table) *PlayerClassT {
	switch rcv {
	case PlayerClassRLBot:
		var x RLBot
		x.Init(table.Bytes, table.Pos)
		return &PlayerClassT{Type: PlayerClassRLBot, Value: x.UnPack()}
	case PlayerClassHuman:
		var x Human
		x.Init(table.Bytes, table.Pos)
		return &PlayerClassT{Type: PlayerClassHuman, Value: x.UnPack()}
	case PlayerClassPsyonix:
		var x Psyonix
		x.Init(table.Bytes, table.Pos)
		return &PlayerClassT{Type: PlayerClassPsyonix, Value: x.UnPack()}
	case PlayerClassPartyMember:
		var x PartyMember
		x.Init(table.Bytes, table.Pos)
		return &PlayerClassT{Type: PlayerClassPartyMember, Value: x.UnPack()}
	}
	return nil
}
