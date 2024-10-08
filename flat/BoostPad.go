// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flat

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type BoostPadT struct {
	Location *Vector3T `json:"location"`
	IsFullBoost bool `json:"is_full_boost"`
}

func (t *BoostPadT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	BoostPadStart(builder)
	locationOffset := t.Location.Pack(builder)
	BoostPadAddLocation(builder, locationOffset)
	BoostPadAddIsFullBoost(builder, t.IsFullBoost)
	return BoostPadEnd(builder)
}

func (rcv *BoostPad) UnPackTo(t *BoostPadT) {
	t.Location = rcv.Location(nil).UnPack()
	t.IsFullBoost = rcv.IsFullBoost()
}

func (rcv *BoostPad) UnPack() *BoostPadT {
	if rcv == nil {
		return nil
	}
	t := &BoostPadT{}
	rcv.UnPackTo(t)
	return t
}

type BoostPad struct {
	_tab flatbuffers.Table
}

func GetRootAsBoostPad(buf []byte, offset flatbuffers.UOffsetT) *BoostPad {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &BoostPad{}
	x.Init(buf, n+offset)
	return x
}

func FinishBoostPadBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsBoostPad(buf []byte, offset flatbuffers.UOffsetT) *BoostPad {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &BoostPad{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedBoostPadBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *BoostPad) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *BoostPad) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *BoostPad) Location(obj *Vector3) *Vector3 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(Vector3)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *BoostPad) IsFullBoost() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *BoostPad) MutateIsFullBoost(n bool) bool {
	return rcv._tab.MutateBoolSlot(6, n)
}

func BoostPadStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func BoostPadAddLocation(builder *flatbuffers.Builder, location flatbuffers.UOffsetT) {
	builder.PrependStructSlot(0, flatbuffers.UOffsetT(location), 0)
}
func BoostPadAddIsFullBoost(builder *flatbuffers.Builder, isFullBoost bool) {
	builder.PrependBoolSlot(1, isFullBoost, false)
}
func BoostPadEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
