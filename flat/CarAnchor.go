// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flat

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type CarAnchorT struct {
	Index uint32 `json:"index"`
	Local *Vector3T `json:"local"`
}

func (t *CarAnchorT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	CarAnchorStart(builder)
	CarAnchorAddIndex(builder, t.Index)
	localOffset := t.Local.Pack(builder)
	CarAnchorAddLocal(builder, localOffset)
	return CarAnchorEnd(builder)
}

func (rcv *CarAnchor) UnPackTo(t *CarAnchorT) {
	t.Index = rcv.Index()
	t.Local = rcv.Local(nil).UnPack()
}

func (rcv *CarAnchor) UnPack() *CarAnchorT {
	if rcv == nil {
		return nil
	}
	t := &CarAnchorT{}
	rcv.UnPackTo(t)
	return t
}

type CarAnchor struct {
	_tab flatbuffers.Table
}

func GetRootAsCarAnchor(buf []byte, offset flatbuffers.UOffsetT) *CarAnchor {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &CarAnchor{}
	x.Init(buf, n+offset)
	return x
}

func FinishCarAnchorBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsCarAnchor(buf []byte, offset flatbuffers.UOffsetT) *CarAnchor {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &CarAnchor{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedCarAnchorBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *CarAnchor) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *CarAnchor) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *CarAnchor) Index() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *CarAnchor) MutateIndex(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

func (rcv *CarAnchor) Local(obj *Vector3) *Vector3 {
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

func CarAnchorStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func CarAnchorAddIndex(builder *flatbuffers.Builder, index uint32) {
	builder.PrependUint32Slot(0, index, 0)
}
func CarAnchorAddLocal(builder *flatbuffers.Builder, local flatbuffers.UOffsetT) {
	builder.PrependStructSlot(1, flatbuffers.UOffsetT(local), 0)
}
func CarAnchorEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}