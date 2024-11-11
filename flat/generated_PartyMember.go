// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flat

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// A player that Rocket League treats as human, e.g. has a dedicated camera and can do training mode,
/// but is actually controlled by a bot.
type PartyMemberT struct {
}

func (t *PartyMemberT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	PartyMemberStart(builder)
	return PartyMemberEnd(builder)
}

func (rcv *PartyMember) UnPackTo(t *PartyMemberT) {
}

func (rcv *PartyMember) UnPack() *PartyMemberT {
	if rcv == nil {
		return nil
	}
	t := &PartyMemberT{}
	rcv.UnPackTo(t)
	return t
}

type PartyMember struct {
	_tab flatbuffers.Table
}

func GetRootAsPartyMember(buf []byte, offset flatbuffers.UOffsetT) *PartyMember {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &PartyMember{}
	x.Init(buf, n+offset)
	return x
}

func FinishPartyMemberBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsPartyMember(buf []byte, offset flatbuffers.UOffsetT) *PartyMember {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &PartyMember{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedPartyMemberBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *PartyMember) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *PartyMember) Table() flatbuffers.Table {
	return rcv._tab
}

func PartyMemberStart(builder *flatbuffers.Builder) {
	builder.StartObject(0)
}
func PartyMemberEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}