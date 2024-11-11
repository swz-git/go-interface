// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flat

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type GoalInfoT struct {
	TeamNum int32 `json:"team_num"`
	Location *Vector3T `json:"location"`
	Direction *Vector3T `json:"direction"`
	Width float32 `json:"width"`
	Height float32 `json:"height"`
}

func (t *GoalInfoT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	GoalInfoStart(builder)
	GoalInfoAddTeamNum(builder, t.TeamNum)
	locationOffset := t.Location.Pack(builder)
	GoalInfoAddLocation(builder, locationOffset)
	directionOffset := t.Direction.Pack(builder)
	GoalInfoAddDirection(builder, directionOffset)
	GoalInfoAddWidth(builder, t.Width)
	GoalInfoAddHeight(builder, t.Height)
	return GoalInfoEnd(builder)
}

func (rcv *GoalInfo) UnPackTo(t *GoalInfoT) {
	t.TeamNum = rcv.TeamNum()
	t.Location = rcv.Location(nil).UnPack()
	t.Direction = rcv.Direction(nil).UnPack()
	t.Width = rcv.Width()
	t.Height = rcv.Height()
}

func (rcv *GoalInfo) UnPack() *GoalInfoT {
	if rcv == nil {
		return nil
	}
	t := &GoalInfoT{}
	rcv.UnPackTo(t)
	return t
}

type GoalInfo struct {
	_tab flatbuffers.Table
}

func GetRootAsGoalInfo(buf []byte, offset flatbuffers.UOffsetT) *GoalInfo {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &GoalInfo{}
	x.Init(buf, n+offset)
	return x
}

func FinishGoalInfoBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsGoalInfo(buf []byte, offset flatbuffers.UOffsetT) *GoalInfo {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &GoalInfo{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedGoalInfoBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *GoalInfo) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *GoalInfo) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *GoalInfo) TeamNum() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *GoalInfo) MutateTeamNum(n int32) bool {
	return rcv._tab.MutateInt32Slot(4, n)
}

func (rcv *GoalInfo) Location(obj *Vector3) *Vector3 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
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

func (rcv *GoalInfo) Direction(obj *Vector3) *Vector3 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
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

func (rcv *GoalInfo) Width() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *GoalInfo) MutateWidth(n float32) bool {
	return rcv._tab.MutateFloat32Slot(10, n)
}

func (rcv *GoalInfo) Height() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *GoalInfo) MutateHeight(n float32) bool {
	return rcv._tab.MutateFloat32Slot(12, n)
}

func GoalInfoStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func GoalInfoAddTeamNum(builder *flatbuffers.Builder, teamNum int32) {
	builder.PrependInt32Slot(0, teamNum, 0)
}
func GoalInfoAddLocation(builder *flatbuffers.Builder, location flatbuffers.UOffsetT) {
	builder.PrependStructSlot(1, flatbuffers.UOffsetT(location), 0)
}
func GoalInfoAddDirection(builder *flatbuffers.Builder, direction flatbuffers.UOffsetT) {
	builder.PrependStructSlot(2, flatbuffers.UOffsetT(direction), 0)
}
func GoalInfoAddWidth(builder *flatbuffers.Builder, width float32) {
	builder.PrependFloat32Slot(3, width, 0.0)
}
func GoalInfoAddHeight(builder *flatbuffers.Builder, height float32) {
	builder.PrependFloat32Slot(4, height, 0.0)
}
func GoalInfoEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}