// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flat

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// A RenderMessage for text in 3D space.
type String3DT struct {
	Text string `json:"text"`
	Anchor *RenderAnchorT `json:"anchor"`
	Scale float32 `json:"scale"`
	Foreground *ColorT `json:"foreground"`
	Background *ColorT `json:"background"`
	HAlign TextHAlign `json:"h_align"`
	VAlign TextVAlign `json:"v_align"`
}

func (t *String3DT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	textOffset := flatbuffers.UOffsetT(0)
	if t.Text != "" {
		textOffset = builder.CreateString(t.Text)
	}
	anchorOffset := t.Anchor.Pack(builder)
	String3DStart(builder)
	String3DAddText(builder, textOffset)
	String3DAddAnchor(builder, anchorOffset)
	String3DAddScale(builder, t.Scale)
	foregroundOffset := t.Foreground.Pack(builder)
	String3DAddForeground(builder, foregroundOffset)
	backgroundOffset := t.Background.Pack(builder)
	String3DAddBackground(builder, backgroundOffset)
	String3DAddHAlign(builder, t.HAlign)
	String3DAddVAlign(builder, t.VAlign)
	return String3DEnd(builder)
}

func (rcv *String3D) UnPackTo(t *String3DT) {
	t.Text = string(rcv.Text())
	t.Anchor = rcv.Anchor(nil).UnPack()
	t.Scale = rcv.Scale()
	t.Foreground = rcv.Foreground(nil).UnPack()
	t.Background = rcv.Background(nil).UnPack()
	t.HAlign = rcv.HAlign()
	t.VAlign = rcv.VAlign()
}

func (rcv *String3D) UnPack() *String3DT {
	if rcv == nil {
		return nil
	}
	t := &String3DT{}
	rcv.UnPackTo(t)
	return t
}

type String3D struct {
	_tab flatbuffers.Table
}

func GetRootAsString3D(buf []byte, offset flatbuffers.UOffsetT) *String3D {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &String3D{}
	x.Init(buf, n+offset)
	return x
}

func FinishString3DBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsString3D(buf []byte, offset flatbuffers.UOffsetT) *String3D {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &String3D{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedString3DBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *String3D) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *String3D) Table() flatbuffers.Table {
	return rcv._tab
}

/// The text to be displayed.
func (rcv *String3D) Text() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// The text to be displayed.
/// The position of the text.
func (rcv *String3D) Anchor(obj *RenderAnchor) *RenderAnchor {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(RenderAnchor)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

/// The position of the text.
/// The scale of the text.
/// When scale is 1, the characters are 20 pixels tall and 10 pixels wide.
func (rcv *String3D) Scale() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// The scale of the text.
/// When scale is 1, the characters are 20 pixels tall and 10 pixels wide.
func (rcv *String3D) MutateScale(n float32) bool {
	return rcv._tab.MutateFloat32Slot(8, n)
}

/// The color of the text.
func (rcv *String3D) Foreground(obj *Color) *Color {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(Color)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

/// The color of the text.
/// The color of the background for the text.
func (rcv *String3D) Background(obj *Color) *Color {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(Color)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

/// The color of the background for the text.
/// The horizontal alignment of the text.
func (rcv *String3D) HAlign() TextHAlign {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return TextHAlign(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

/// The horizontal alignment of the text.
func (rcv *String3D) MutateHAlign(n TextHAlign) bool {
	return rcv._tab.MutateByteSlot(14, byte(n))
}

/// The vertical alignment of the text.
func (rcv *String3D) VAlign() TextVAlign {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return TextVAlign(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

/// The vertical alignment of the text.
func (rcv *String3D) MutateVAlign(n TextVAlign) bool {
	return rcv._tab.MutateByteSlot(16, byte(n))
}

func String3DStart(builder *flatbuffers.Builder) {
	builder.StartObject(7)
}
func String3DAddText(builder *flatbuffers.Builder, text flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(text), 0)
}
func String3DAddAnchor(builder *flatbuffers.Builder, anchor flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(anchor), 0)
}
func String3DAddScale(builder *flatbuffers.Builder, scale float32) {
	builder.PrependFloat32Slot(2, scale, 0.0)
}
func String3DAddForeground(builder *flatbuffers.Builder, foreground flatbuffers.UOffsetT) {
	builder.PrependStructSlot(3, flatbuffers.UOffsetT(foreground), 0)
}
func String3DAddBackground(builder *flatbuffers.Builder, background flatbuffers.UOffsetT) {
	builder.PrependStructSlot(4, flatbuffers.UOffsetT(background), 0)
}
func String3DAddHAlign(builder *flatbuffers.Builder, hAlign TextHAlign) {
	builder.PrependByteSlot(5, byte(hAlign), 0)
}
func String3DAddVAlign(builder *flatbuffers.Builder, vAlign TextVAlign) {
	builder.PrependByteSlot(6, byte(vAlign), 0)
}
func String3DEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
