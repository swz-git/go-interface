// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flat

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// The state of a boost pad.
/// Note, static properties of boost pads, such as their location and size, are found in the field info.
type BoostPadStateT struct {
	IsActive bool `json:"is_active"`
	Timer float32 `json:"timer"`
}

func (t *BoostPadStateT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	return CreateBoostPadState(builder, t.IsActive, t.Timer)
}
func (rcv *BoostPadState) UnPackTo(t *BoostPadStateT) {
	t.IsActive = rcv.IsActive()
	t.Timer = rcv.Timer()
}

func (rcv *BoostPadState) UnPack() *BoostPadStateT {
	if rcv == nil {
		return nil
	}
	t := &BoostPadStateT{}
	rcv.UnPackTo(t)
	return t
}

type BoostPadState struct {
	_tab flatbuffers.Struct
}

func (rcv *BoostPadState) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *BoostPadState) Table() flatbuffers.Table {
	return rcv._tab.Table
}

/// True if the boost can be picked up right now.
func (rcv *BoostPadState) IsActive() bool {
	return rcv._tab.GetBool(rcv._tab.Pos + flatbuffers.UOffsetT(0))
}
/// True if the boost can be picked up right now.
func (rcv *BoostPadState) MutateIsActive(n bool) bool {
	return rcv._tab.MutateBool(rcv._tab.Pos+flatbuffers.UOffsetT(0), n)
}

/// The number of seconds since the boost has been picked up, or 0 if the boost is active.
/// A big boost pad becomes active again after 10 seconds.
/// A small boost pad becomes active again after 4 seconds.
func (rcv *BoostPadState) Timer() float32 {
	return rcv._tab.GetFloat32(rcv._tab.Pos + flatbuffers.UOffsetT(4))
}
/// The number of seconds since the boost has been picked up, or 0 if the boost is active.
/// A big boost pad becomes active again after 10 seconds.
/// A small boost pad becomes active again after 4 seconds.
func (rcv *BoostPadState) MutateTimer(n float32) bool {
	return rcv._tab.MutateFloat32(rcv._tab.Pos+flatbuffers.UOffsetT(4), n)
}

func CreateBoostPadState(builder *flatbuffers.Builder, isActive bool, timer float32) flatbuffers.UOffsetT {
	builder.Prep(4, 8)
	builder.PrependFloat32(timer)
	builder.Pad(3)
	builder.PrependBool(isActive)
	return builder.Offset()
}
