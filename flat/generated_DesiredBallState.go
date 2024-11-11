// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flat

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type DesiredBallStateT struct {
	Physics *DesiredPhysicsT `json:"physics"`
}

func (t *DesiredBallStateT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	physicsOffset := t.Physics.Pack(builder)
	DesiredBallStateStart(builder)
	DesiredBallStateAddPhysics(builder, physicsOffset)
	return DesiredBallStateEnd(builder)
}

func (rcv *DesiredBallState) UnPackTo(t *DesiredBallStateT) {
	t.Physics = rcv.Physics(nil).UnPack()
}

func (rcv *DesiredBallState) UnPack() *DesiredBallStateT {
	if rcv == nil {
		return nil
	}
	t := &DesiredBallStateT{}
	rcv.UnPackTo(t)
	return t
}

type DesiredBallState struct {
	_tab flatbuffers.Table
}

func GetRootAsDesiredBallState(buf []byte, offset flatbuffers.UOffsetT) *DesiredBallState {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &DesiredBallState{}
	x.Init(buf, n+offset)
	return x
}

func FinishDesiredBallStateBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsDesiredBallState(buf []byte, offset flatbuffers.UOffsetT) *DesiredBallState {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &DesiredBallState{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedDesiredBallStateBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *DesiredBallState) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *DesiredBallState) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *DesiredBallState) Physics(obj *DesiredPhysics) *DesiredPhysics {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(DesiredPhysics)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func DesiredBallStateStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func DesiredBallStateAddPhysics(builder *flatbuffers.Builder, physics flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(physics), 0)
}
func DesiredBallStateEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}