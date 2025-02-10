// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flat

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// Match info with nullable components.
/// Used for game state setting to define which part of the match info should change.
type DesiredMatchInfoT struct {
	WorldGravityZ *FloatT `json:"world_gravity_z"`
	GameSpeed *FloatT `json:"game_speed"`
}

func (t *DesiredMatchInfoT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	DesiredMatchInfoStart(builder)
	worldGravityZOffset := t.WorldGravityZ.Pack(builder)
	DesiredMatchInfoAddWorldGravityZ(builder, worldGravityZOffset)
	gameSpeedOffset := t.GameSpeed.Pack(builder)
	DesiredMatchInfoAddGameSpeed(builder, gameSpeedOffset)
	return DesiredMatchInfoEnd(builder)
}

func (rcv *DesiredMatchInfo) UnPackTo(t *DesiredMatchInfoT) {
	t.WorldGravityZ = rcv.WorldGravityZ(nil).UnPack()
	t.GameSpeed = rcv.GameSpeed(nil).UnPack()
}

func (rcv *DesiredMatchInfo) UnPack() *DesiredMatchInfoT {
	if rcv == nil {
		return nil
	}
	t := &DesiredMatchInfoT{}
	rcv.UnPackTo(t)
	return t
}

type DesiredMatchInfo struct {
	_tab flatbuffers.Table
}

func GetRootAsDesiredMatchInfo(buf []byte, offset flatbuffers.UOffsetT) *DesiredMatchInfo {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &DesiredMatchInfo{}
	x.Init(buf, n+offset)
	return x
}

func FinishDesiredMatchInfoBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsDesiredMatchInfo(buf []byte, offset flatbuffers.UOffsetT) *DesiredMatchInfo {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &DesiredMatchInfo{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedDesiredMatchInfoBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *DesiredMatchInfo) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *DesiredMatchInfo) Table() flatbuffers.Table {
	return rcv._tab
}

/// The strength of gravity. Default is usually -650 depending on mutators.
/// To set gravity to 0, use 0.0000001 instead, as 0 will set gravity back to the default.
func (rcv *DesiredMatchInfo) WorldGravityZ(obj *Float) *Float {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(Float)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

/// The strength of gravity. Default is usually -650 depending on mutators.
/// To set gravity to 0, use 0.0000001 instead, as 0 will set gravity back to the default.
/// The game speed. Default is 1.0.
func (rcv *DesiredMatchInfo) GameSpeed(obj *Float) *Float {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(Float)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

/// The game speed. Default is 1.0.
func DesiredMatchInfoStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func DesiredMatchInfoAddWorldGravityZ(builder *flatbuffers.Builder, worldGravityZ flatbuffers.UOffsetT) {
	builder.PrependStructSlot(0, flatbuffers.UOffsetT(worldGravityZ), 0)
}
func DesiredMatchInfoAddGameSpeed(builder *flatbuffers.Builder, gameSpeed flatbuffers.UOffsetT) {
	builder.PrependStructSlot(1, flatbuffers.UOffsetT(gameSpeed), 0)
}
func DesiredMatchInfoEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
