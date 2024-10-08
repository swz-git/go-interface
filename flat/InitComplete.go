// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flat

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type InitCompleteT struct {
	SpawnId int32 `json:"spawn_id"`
}

func (t *InitCompleteT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	InitCompleteStart(builder)
	InitCompleteAddSpawnId(builder, t.SpawnId)
	return InitCompleteEnd(builder)
}

func (rcv *InitComplete) UnPackTo(t *InitCompleteT) {
	t.SpawnId = rcv.SpawnId()
}

func (rcv *InitComplete) UnPack() *InitCompleteT {
	if rcv == nil {
		return nil
	}
	t := &InitCompleteT{}
	rcv.UnPackTo(t)
	return t
}

type InitComplete struct {
	_tab flatbuffers.Table
}

func GetRootAsInitComplete(buf []byte, offset flatbuffers.UOffsetT) *InitComplete {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &InitComplete{}
	x.Init(buf, n+offset)
	return x
}

func FinishInitCompleteBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsInitComplete(buf []byte, offset flatbuffers.UOffsetT) *InitComplete {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &InitComplete{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedInitCompleteBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *InitComplete) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *InitComplete) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *InitComplete) SpawnId() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *InitComplete) MutateSpawnId(n int32) bool {
	return rcv._tab.MutateInt32Slot(4, n)
}

func InitCompleteStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func InitCompleteAddSpawnId(builder *flatbuffers.Builder, spawnId int32) {
	builder.PrependInt32Slot(0, spawnId, 0)
}
func InitCompleteEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
