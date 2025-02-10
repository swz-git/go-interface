// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flat

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// A combination of button presses and analog steering values like those produced by a physical controller or keyboard.
/// This is sent by bots each tick to RLBot to indicate what they want to do that tick.
/// For example, if you want to hold the jump button for 20 ticks, then you must send 20 controller states where jump is true.
/// Remember to send controller states with jump set to false to let go of the jump button afterwards.
type ControllerStateT struct {
	Throttle float32 `json:"throttle"`
	Steer float32 `json:"steer"`
	Pitch float32 `json:"pitch"`
	Yaw float32 `json:"yaw"`
	Roll float32 `json:"roll"`
	Jump bool `json:"jump"`
	Boost bool `json:"boost"`
	Handbrake bool `json:"handbrake"`
	UseItem bool `json:"use_item"`
}

func (t *ControllerStateT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	return CreateControllerState(builder, t.Throttle, t.Steer, t.Pitch, t.Yaw, t.Roll, t.Jump, t.Boost, t.Handbrake, t.UseItem)
}
func (rcv *ControllerState) UnPackTo(t *ControllerStateT) {
	t.Throttle = rcv.Throttle()
	t.Steer = rcv.Steer()
	t.Pitch = rcv.Pitch()
	t.Yaw = rcv.Yaw()
	t.Roll = rcv.Roll()
	t.Jump = rcv.Jump()
	t.Boost = rcv.Boost()
	t.Handbrake = rcv.Handbrake()
	t.UseItem = rcv.UseItem()
}

func (rcv *ControllerState) UnPack() *ControllerStateT {
	if rcv == nil {
		return nil
	}
	t := &ControllerStateT{}
	rcv.UnPackTo(t)
	return t
}

type ControllerState struct {
	_tab flatbuffers.Struct
}

func (rcv *ControllerState) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ControllerState) Table() flatbuffers.Table {
	return rcv._tab.Table
}

/// -1 for full reverse, 1 for full forward.
func (rcv *ControllerState) Throttle() float32 {
	return rcv._tab.GetFloat32(rcv._tab.Pos + flatbuffers.UOffsetT(0))
}
/// -1 for full reverse, 1 for full forward.
func (rcv *ControllerState) MutateThrottle(n float32) bool {
	return rcv._tab.MutateFloat32(rcv._tab.Pos+flatbuffers.UOffsetT(0), n)
}

/// -1 for full left, 1 for full right.
func (rcv *ControllerState) Steer() float32 {
	return rcv._tab.GetFloat32(rcv._tab.Pos + flatbuffers.UOffsetT(4))
}
/// -1 for full left, 1 for full right.
func (rcv *ControllerState) MutateSteer(n float32) bool {
	return rcv._tab.MutateFloat32(rcv._tab.Pos+flatbuffers.UOffsetT(4), n)
}

/// -1 for nose down, 1 for nose up.
func (rcv *ControllerState) Pitch() float32 {
	return rcv._tab.GetFloat32(rcv._tab.Pos + flatbuffers.UOffsetT(8))
}
/// -1 for nose down, 1 for nose up.
func (rcv *ControllerState) MutatePitch(n float32) bool {
	return rcv._tab.MutateFloat32(rcv._tab.Pos+flatbuffers.UOffsetT(8), n)
}

/// -1 for full left, 1 for full right.
func (rcv *ControllerState) Yaw() float32 {
	return rcv._tab.GetFloat32(rcv._tab.Pos + flatbuffers.UOffsetT(12))
}
/// -1 for full left, 1 for full right.
func (rcv *ControllerState) MutateYaw(n float32) bool {
	return rcv._tab.MutateFloat32(rcv._tab.Pos+flatbuffers.UOffsetT(12), n)
}

/// -1 for roll left, 1 for roll right.
func (rcv *ControllerState) Roll() float32 {
	return rcv._tab.GetFloat32(rcv._tab.Pos + flatbuffers.UOffsetT(16))
}
/// -1 for roll left, 1 for roll right.
func (rcv *ControllerState) MutateRoll(n float32) bool {
	return rcv._tab.MutateFloat32(rcv._tab.Pos+flatbuffers.UOffsetT(16), n)
}

/// True if you want to press the jump button.
func (rcv *ControllerState) Jump() bool {
	return rcv._tab.GetBool(rcv._tab.Pos + flatbuffers.UOffsetT(20))
}
/// True if you want to press the jump button.
func (rcv *ControllerState) MutateJump(n bool) bool {
	return rcv._tab.MutateBool(rcv._tab.Pos+flatbuffers.UOffsetT(20), n)
}

/// True if you want to press the boost button.
func (rcv *ControllerState) Boost() bool {
	return rcv._tab.GetBool(rcv._tab.Pos + flatbuffers.UOffsetT(21))
}
/// True if you want to press the boost button.
func (rcv *ControllerState) MutateBoost(n bool) bool {
	return rcv._tab.MutateBool(rcv._tab.Pos+flatbuffers.UOffsetT(21), n)
}

/// True if you want to press the handbrake button.
func (rcv *ControllerState) Handbrake() bool {
	return rcv._tab.GetBool(rcv._tab.Pos + flatbuffers.UOffsetT(22))
}
/// True if you want to press the handbrake button.
func (rcv *ControllerState) MutateHandbrake(n bool) bool {
	return rcv._tab.MutateBool(rcv._tab.Pos+flatbuffers.UOffsetT(22), n)
}

/// True if you want to press the 'use item' button. Used in Rumble and other game modes.
func (rcv *ControllerState) UseItem() bool {
	return rcv._tab.GetBool(rcv._tab.Pos + flatbuffers.UOffsetT(23))
}
/// True if you want to press the 'use item' button. Used in Rumble and other game modes.
func (rcv *ControllerState) MutateUseItem(n bool) bool {
	return rcv._tab.MutateBool(rcv._tab.Pos+flatbuffers.UOffsetT(23), n)
}

func CreateControllerState(builder *flatbuffers.Builder, throttle float32, steer float32, pitch float32, yaw float32, roll float32, jump bool, boost bool, handbrake bool, useItem bool) flatbuffers.UOffsetT {
	builder.Prep(4, 24)
	builder.PrependBool(useItem)
	builder.PrependBool(handbrake)
	builder.PrependBool(boost)
	builder.PrependBool(jump)
	builder.PrependFloat32(roll)
	builder.PrependFloat32(yaw)
	builder.PrependFloat32(pitch)
	builder.PrependFloat32(steer)
	builder.PrependFloat32(throttle)
	return builder.Offset()
}
