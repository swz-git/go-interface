// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flat

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// A group of RenderMessages that are drawn and cleared together.
/// A RenderGroup will stay rendered until it is overriden or cleared.
/// The group is identified by a unique id.
/// A client can only clear its own RenderGroups.
type RenderGroupT struct {
	RenderMessages []*RenderMessageT `json:"render_messages"`
	Id int32 `json:"id"`
}

func (t *RenderGroupT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	renderMessagesOffset := flatbuffers.UOffsetT(0)
	if t.RenderMessages != nil {
		renderMessagesLength := len(t.RenderMessages)
		renderMessagesOffsets := make([]flatbuffers.UOffsetT, renderMessagesLength)
		for j := 0; j < renderMessagesLength; j++ {
			renderMessagesOffsets[j] = t.RenderMessages[j].Pack(builder)
		}
		RenderGroupStartRenderMessagesVector(builder, renderMessagesLength)
		for j := renderMessagesLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(renderMessagesOffsets[j])
		}
		renderMessagesOffset = builder.EndVector(renderMessagesLength)
	}
	RenderGroupStart(builder)
	RenderGroupAddRenderMessages(builder, renderMessagesOffset)
	RenderGroupAddId(builder, t.Id)
	return RenderGroupEnd(builder)
}

func (rcv *RenderGroup) UnPackTo(t *RenderGroupT) {
	renderMessagesLength := rcv.RenderMessagesLength()
	t.RenderMessages = make([]*RenderMessageT, renderMessagesLength)
	for j := 0; j < renderMessagesLength; j++ {
		x := RenderMessage{}
		rcv.RenderMessages(&x, j)
		t.RenderMessages[j] = x.UnPack()
	}
	t.Id = rcv.Id()
}

func (rcv *RenderGroup) UnPack() *RenderGroupT {
	if rcv == nil {
		return nil
	}
	t := &RenderGroupT{}
	rcv.UnPackTo(t)
	return t
}

type RenderGroup struct {
	_tab flatbuffers.Table
}

func GetRootAsRenderGroup(buf []byte, offset flatbuffers.UOffsetT) *RenderGroup {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &RenderGroup{}
	x.Init(buf, n+offset)
	return x
}

func FinishRenderGroupBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsRenderGroup(buf []byte, offset flatbuffers.UOffsetT) *RenderGroup {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &RenderGroup{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedRenderGroupBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *RenderGroup) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *RenderGroup) Table() flatbuffers.Table {
	return rcv._tab
}

/// The content of the RenderGroup.
func (rcv *RenderGroup) RenderMessages(obj *RenderMessage, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *RenderGroup) RenderMessagesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// The content of the RenderGroup.
/// The id of the RenderGroup.
func (rcv *RenderGroup) Id() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

/// The id of the RenderGroup.
func (rcv *RenderGroup) MutateId(n int32) bool {
	return rcv._tab.MutateInt32Slot(6, n)
}

func RenderGroupStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func RenderGroupAddRenderMessages(builder *flatbuffers.Builder, renderMessages flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(renderMessages), 0)
}
func RenderGroupStartRenderMessagesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func RenderGroupAddId(builder *flatbuffers.Builder, id int32) {
	builder.PrependInt32Slot(1, id, 0)
}
func RenderGroupEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
