// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flat

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// A client message to change the loadout of a car.
/// If sent before the ready message, this simply sets the loadout of the car.
/// If sent after the ready message and if game state setting is enabled, this will respawn the car with the new loadout.
/// Bots can only set the loadout of their own car(s).
type SetLoadoutT struct {
	Index uint32 `json:"index"`
	Loadout *PlayerLoadoutT `json:"loadout"`
}

func (t *SetLoadoutT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	loadoutOffset := t.Loadout.Pack(builder)
	SetLoadoutStart(builder)
	SetLoadoutAddIndex(builder, t.Index)
	SetLoadoutAddLoadout(builder, loadoutOffset)
	return SetLoadoutEnd(builder)
}

func (rcv *SetLoadout) UnPackTo(t *SetLoadoutT) {
	t.Index = rcv.Index()
	t.Loadout = rcv.Loadout(nil).UnPack()
}

func (rcv *SetLoadout) UnPack() *SetLoadoutT {
	if rcv == nil {
		return nil
	}
	t := &SetLoadoutT{}
	rcv.UnPackTo(t)
	return t
}

type SetLoadout struct {
	_tab flatbuffers.Table
}

func GetRootAsSetLoadout(buf []byte, offset flatbuffers.UOffsetT) *SetLoadout {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &SetLoadout{}
	x.Init(buf, n+offset)
	return x
}

func FinishSetLoadoutBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsSetLoadout(buf []byte, offset flatbuffers.UOffsetT) *SetLoadout {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &SetLoadout{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedSetLoadoutBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *SetLoadout) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *SetLoadout) Table() flatbuffers.Table {
	return rcv._tab
}

/// The index of the car to change loadout off.
func (rcv *SetLoadout) Index() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

/// The index of the car to change loadout off.
func (rcv *SetLoadout) MutateIndex(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

/// The new loadout of the car.
func (rcv *SetLoadout) Loadout(obj *PlayerLoadout) *PlayerLoadout {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(PlayerLoadout)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

/// The new loadout of the car.
func SetLoadoutStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func SetLoadoutAddIndex(builder *flatbuffers.Builder, index uint32) {
	builder.PrependUint32Slot(0, index, 0)
}
func SetLoadoutAddLoadout(builder *flatbuffers.Builder, loadout flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(loadout), 0)
}
func SetLoadoutEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
