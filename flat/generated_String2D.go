// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flat

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type String2DT struct {
	Text string `json:"text"`
	X float32 `json:"x"`
	Y float32 `json:"y"`
	Scale float32 `json:"scale"`
	Foreground *ColorT `json:"foreground"`
	Background *ColorT `json:"background"`
	HAlign TextHAlign `json:"h_align"`
	VAlign TextVAlign `json:"v_align"`
}

func (t *String2DT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	textOffset := flatbuffers.UOffsetT(0)
	if t.Text != "" {
		textOffset = builder.CreateString(t.Text)
	}
	String2DStart(builder)
	String2DAddText(builder, textOffset)
	String2DAddX(builder, t.X)
	String2DAddY(builder, t.Y)
	String2DAddScale(builder, t.Scale)
	foregroundOffset := t.Foreground.Pack(builder)
	String2DAddForeground(builder, foregroundOffset)
	backgroundOffset := t.Background.Pack(builder)
	String2DAddBackground(builder, backgroundOffset)
	String2DAddHAlign(builder, t.HAlign)
	String2DAddVAlign(builder, t.VAlign)
	return String2DEnd(builder)
}

func (rcv *String2D) UnPackTo(t *String2DT) {
	t.Text = string(rcv.Text())
	t.X = rcv.X()
	t.Y = rcv.Y()
	t.Scale = rcv.Scale()
	t.Foreground = rcv.Foreground(nil).UnPack()
	t.Background = rcv.Background(nil).UnPack()
	t.HAlign = rcv.HAlign()
	t.VAlign = rcv.VAlign()
}

func (rcv *String2D) UnPack() *String2DT {
	if rcv == nil {
		return nil
	}
	t := &String2DT{}
	rcv.UnPackTo(t)
	return t
}

type String2D struct {
	_tab flatbuffers.Table
}

func GetRootAsString2D(buf []byte, offset flatbuffers.UOffsetT) *String2D {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &String2D{}
	x.Init(buf, n+offset)
	return x
}

func FinishString2DBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsString2D(buf []byte, offset flatbuffers.UOffsetT) *String2D {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &String2D{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedString2DBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *String2D) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *String2D) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *String2D) Text() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Screen-space coordinates such that x=0 is left edge and x=1 is right edge of window.
func (rcv *String2D) X() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Screen-space coordinates such that x=0 is left edge and x=1 is right edge of window.
func (rcv *String2D) MutateX(n float32) bool {
	return rcv._tab.MutateFloat32Slot(6, n)
}

/// Screen-space coordinates such that y=0 is top edge and y=1 is bottom edge of window.
func (rcv *String2D) Y() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Screen-space coordinates such that y=0 is top edge and y=1 is bottom edge of window.
func (rcv *String2D) MutateY(n float32) bool {
	return rcv._tab.MutateFloat32Slot(8, n)
}

func (rcv *String2D) Scale() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *String2D) MutateScale(n float32) bool {
	return rcv._tab.MutateFloat32Slot(10, n)
}

func (rcv *String2D) Foreground(obj *Color) *Color {
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

func (rcv *String2D) Background(obj *Color) *Color {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
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

func (rcv *String2D) HAlign() TextHAlign {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return TextHAlign(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *String2D) MutateHAlign(n TextHAlign) bool {
	return rcv._tab.MutateByteSlot(16, byte(n))
}

func (rcv *String2D) VAlign() TextVAlign {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return TextVAlign(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *String2D) MutateVAlign(n TextVAlign) bool {
	return rcv._tab.MutateByteSlot(18, byte(n))
}

func String2DStart(builder *flatbuffers.Builder) {
	builder.StartObject(8)
}
func String2DAddText(builder *flatbuffers.Builder, text flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(text), 0)
}
func String2DAddX(builder *flatbuffers.Builder, x float32) {
	builder.PrependFloat32Slot(1, x, 0.0)
}
func String2DAddY(builder *flatbuffers.Builder, y float32) {
	builder.PrependFloat32Slot(2, y, 0.0)
}
func String2DAddScale(builder *flatbuffers.Builder, scale float32) {
	builder.PrependFloat32Slot(3, scale, 0.0)
}
func String2DAddForeground(builder *flatbuffers.Builder, foreground flatbuffers.UOffsetT) {
	builder.PrependStructSlot(4, flatbuffers.UOffsetT(foreground), 0)
}
func String2DAddBackground(builder *flatbuffers.Builder, background flatbuffers.UOffsetT) {
	builder.PrependStructSlot(5, flatbuffers.UOffsetT(background), 0)
}
func String2DAddHAlign(builder *flatbuffers.Builder, hAlign TextHAlign) {
	builder.PrependByteSlot(6, byte(hAlign), 0)
}
func String2DAddVAlign(builder *flatbuffers.Builder, vAlign TextVAlign) {
	builder.PrependByteSlot(7, byte(vAlign), 0)
}
func String2DEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}