// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flat

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// Information about the current match such as time and gravity.
type MatchInfoT struct {
	SecondsElapsed float32 `json:"seconds_elapsed"`
	GameTimeRemaining float32 `json:"game_time_remaining"`
	IsOvertime bool `json:"is_overtime"`
	IsUnlimitedTime bool `json:"is_unlimited_time"`
	MatchPhase MatchPhase `json:"match_phase"`
	WorldGravityZ float32 `json:"world_gravity_z"`
	GameSpeed float32 `json:"game_speed"`
	LastSpectated uint32 `json:"last_spectated"`
	FrameNum uint32 `json:"frame_num"`
}

func (t *MatchInfoT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	MatchInfoStart(builder)
	MatchInfoAddSecondsElapsed(builder, t.SecondsElapsed)
	MatchInfoAddGameTimeRemaining(builder, t.GameTimeRemaining)
	MatchInfoAddIsOvertime(builder, t.IsOvertime)
	MatchInfoAddIsUnlimitedTime(builder, t.IsUnlimitedTime)
	MatchInfoAddMatchPhase(builder, t.MatchPhase)
	MatchInfoAddWorldGravityZ(builder, t.WorldGravityZ)
	MatchInfoAddGameSpeed(builder, t.GameSpeed)
	MatchInfoAddLastSpectated(builder, t.LastSpectated)
	MatchInfoAddFrameNum(builder, t.FrameNum)
	return MatchInfoEnd(builder)
}

func (rcv *MatchInfo) UnPackTo(t *MatchInfoT) {
	t.SecondsElapsed = rcv.SecondsElapsed()
	t.GameTimeRemaining = rcv.GameTimeRemaining()
	t.IsOvertime = rcv.IsOvertime()
	t.IsUnlimitedTime = rcv.IsUnlimitedTime()
	t.MatchPhase = rcv.MatchPhase()
	t.WorldGravityZ = rcv.WorldGravityZ()
	t.GameSpeed = rcv.GameSpeed()
	t.LastSpectated = rcv.LastSpectated()
	t.FrameNum = rcv.FrameNum()
}

func (rcv *MatchInfo) UnPack() *MatchInfoT {
	if rcv == nil {
		return nil
	}
	t := &MatchInfoT{}
	rcv.UnPackTo(t)
	return t
}

type MatchInfo struct {
	_tab flatbuffers.Table
}

func GetRootAsMatchInfo(buf []byte, offset flatbuffers.UOffsetT) *MatchInfo {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &MatchInfo{}
	x.Init(buf, n+offset)
	return x
}

func FinishMatchInfoBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsMatchInfo(buf []byte, offset flatbuffers.UOffsetT) *MatchInfo {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &MatchInfo{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedMatchInfoBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *MatchInfo) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *MatchInfo) Table() flatbuffers.Table {
	return rcv._tab
}

/// How many seconds have elapsed since the first game packet of the match.
/// This value ticks up even during kickoffs, replays, pause, etc.
func (rcv *MatchInfo) SecondsElapsed() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// How many seconds have elapsed since the first game packet of the match.
/// This value ticks up even during kickoffs, replays, pause, etc.
func (rcv *MatchInfo) MutateSecondsElapsed(n float32) bool {
	return rcv._tab.MutateFloat32Slot(4, n)
}

/// Seconds remaining of the match.
/// This value ticks up instead of down during overtime or when the game duration mutator is set to Unlimited.
/// I.e. it matches the in-game timer at the top.
func (rcv *MatchInfo) GameTimeRemaining() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Seconds remaining of the match.
/// This value ticks up instead of down during overtime or when the game duration mutator is set to Unlimited.
/// I.e. it matches the in-game timer at the top.
func (rcv *MatchInfo) MutateGameTimeRemaining(n float32) bool {
	return rcv._tab.MutateFloat32Slot(6, n)
}

/// True if the game is in overtime.
func (rcv *MatchInfo) IsOvertime() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

/// True if the game is in overtime.
func (rcv *MatchInfo) MutateIsOvertime(n bool) bool {
	return rcv._tab.MutateBoolSlot(8, n)
}

/// True if the game duration is set to Unlimited.
func (rcv *MatchInfo) IsUnlimitedTime() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

/// True if the game duration is set to Unlimited.
func (rcv *MatchInfo) MutateIsUnlimitedTime(n bool) bool {
	return rcv._tab.MutateBoolSlot(10, n)
}

/// The current phase of the match, i.e. kickoff, replay, active, etc.
func (rcv *MatchInfo) MatchPhase() MatchPhase {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return MatchPhase(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

/// The current phase of the match, i.e. kickoff, replay, active, etc.
func (rcv *MatchInfo) MutateMatchPhase(n MatchPhase) bool {
	return rcv._tab.MutateByteSlot(12, byte(n))
}

/// The current strength of gravity. Default is -650.
func (rcv *MatchInfo) WorldGravityZ() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// The current strength of gravity. Default is -650.
func (rcv *MatchInfo) MutateWorldGravityZ(n float32) bool {
	return rcv._tab.MutateFloat32Slot(14, n)
}

/// Game speed multiplier. Regular game speed is 1.0.
func (rcv *MatchInfo) GameSpeed() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Game speed multiplier. Regular game speed is 1.0.
func (rcv *MatchInfo) MutateGameSpeed(n float32) bool {
	return rcv._tab.MutateFloat32Slot(16, n)
}

/// Index of the player who was most recently a spectated by the host.
func (rcv *MatchInfo) LastSpectated() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

/// Index of the player who was most recently a spectated by the host.
func (rcv *MatchInfo) MutateLastSpectated(n uint32) bool {
	return rcv._tab.MutateUint32Slot(18, n)
}

/// Tracks the number of physics frames the game has computed.
/// May increase by more than one across consecutive packets.
/// Data type will roll over after 414 days at 120Hz.
func (rcv *MatchInfo) FrameNum() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

/// Tracks the number of physics frames the game has computed.
/// May increase by more than one across consecutive packets.
/// Data type will roll over after 414 days at 120Hz.
func (rcv *MatchInfo) MutateFrameNum(n uint32) bool {
	return rcv._tab.MutateUint32Slot(20, n)
}

func MatchInfoStart(builder *flatbuffers.Builder) {
	builder.StartObject(9)
}
func MatchInfoAddSecondsElapsed(builder *flatbuffers.Builder, secondsElapsed float32) {
	builder.PrependFloat32Slot(0, secondsElapsed, 0.0)
}
func MatchInfoAddGameTimeRemaining(builder *flatbuffers.Builder, gameTimeRemaining float32) {
	builder.PrependFloat32Slot(1, gameTimeRemaining, 0.0)
}
func MatchInfoAddIsOvertime(builder *flatbuffers.Builder, isOvertime bool) {
	builder.PrependBoolSlot(2, isOvertime, false)
}
func MatchInfoAddIsUnlimitedTime(builder *flatbuffers.Builder, isUnlimitedTime bool) {
	builder.PrependBoolSlot(3, isUnlimitedTime, false)
}
func MatchInfoAddMatchPhase(builder *flatbuffers.Builder, matchPhase MatchPhase) {
	builder.PrependByteSlot(4, byte(matchPhase), 0)
}
func MatchInfoAddWorldGravityZ(builder *flatbuffers.Builder, worldGravityZ float32) {
	builder.PrependFloat32Slot(5, worldGravityZ, 0.0)
}
func MatchInfoAddGameSpeed(builder *flatbuffers.Builder, gameSpeed float32) {
	builder.PrependFloat32Slot(6, gameSpeed, 0.0)
}
func MatchInfoAddLastSpectated(builder *flatbuffers.Builder, lastSpectated uint32) {
	builder.PrependUint32Slot(7, lastSpectated, 0)
}
func MatchInfoAddFrameNum(builder *flatbuffers.Builder, frameNum uint32) {
	builder.PrependUint32Slot(8, frameNum, 0)
}
func MatchInfoEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
