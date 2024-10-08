// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flat

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type RenderAnchorT struct {
	World *Vector3T `json:"world"`
	Relative *RelativeAnchorT `json:"relative"`
}

func (t *RenderAnchorT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	relativeOffset := t.Relative.Pack(builder)

	RenderAnchorStart(builder)
	worldOffset := t.World.Pack(builder)
	RenderAnchorAddWorld(builder, worldOffset)
	if t.Relative != nil {
		RenderAnchorAddRelativeType(builder, t.Relative.Type)
	}
	RenderAnchorAddRelative(builder, relativeOffset)
	return RenderAnchorEnd(builder)
}

func (rcv *RenderAnchor) UnPackTo(t *RenderAnchorT) {
	t.World = rcv.World(nil).UnPack()
	relativeTable := flatbuffers.Table{}
	if rcv.Relative(&relativeTable) {
		t.Relative = rcv.RelativeType().UnPack(relativeTable)
	}
}

func (rcv *RenderAnchor) UnPack() *RenderAnchorT {
	if rcv == nil {
		return nil
	}
	t := &RenderAnchorT{}
	rcv.UnPackTo(t)
	return t
}

type RenderAnchor struct {
	_tab flatbuffers.Table
}

func GetRootAsRenderAnchor(buf []byte, offset flatbuffers.UOffsetT) *RenderAnchor {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &RenderAnchor{}
	x.Init(buf, n+offset)
	return x
}

func FinishRenderAnchorBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsRenderAnchor(buf []byte, offset flatbuffers.UOffsetT) *RenderAnchor {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &RenderAnchor{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedRenderAnchorBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *RenderAnchor) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *RenderAnchor) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *RenderAnchor) World(obj *Vector3) *Vector3 {
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

func (rcv *RenderAnchor) RelativeType() RelativeAnchor {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return RelativeAnchor(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *RenderAnchor) MutateRelativeType(n RelativeAnchor) bool {
	return rcv._tab.MutateByteSlot(6, byte(n))
}

func (rcv *RenderAnchor) Relative(obj *flatbuffers.Table) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		rcv._tab.Union(obj, o)
		return true
	}
	return false
}

func RenderAnchorStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func RenderAnchorAddWorld(builder *flatbuffers.Builder, world flatbuffers.UOffsetT) {
	builder.PrependStructSlot(0, flatbuffers.UOffsetT(world), 0)
}
func RenderAnchorAddRelativeType(builder *flatbuffers.Builder, relativeType RelativeAnchor) {
	builder.PrependByteSlot(1, byte(relativeType), 0)
}
func RenderAnchorAddRelative(builder *flatbuffers.Builder, relative flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(relative), 0)
}
func RenderAnchorEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
