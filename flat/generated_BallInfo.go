// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flat

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type BallInfoT struct {
	Physics *PhysicsT `json:"physics"`
	Shape *CollisionShapeT `json:"shape"`
}

func (t *BallInfoT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	shapeOffset := t.Shape.Pack(builder)

	BallInfoStart(builder)
	physicsOffset := t.Physics.Pack(builder)
	BallInfoAddPhysics(builder, physicsOffset)
	if t.Shape != nil {
		BallInfoAddShapeType(builder, t.Shape.Type)
	}
	BallInfoAddShape(builder, shapeOffset)
	return BallInfoEnd(builder)
}

func (rcv *BallInfo) UnPackTo(t *BallInfoT) {
	t.Physics = rcv.Physics(nil).UnPack()
	shapeTable := flatbuffers.Table{}
	if rcv.Shape(&shapeTable) {
		t.Shape = rcv.ShapeType().UnPack(shapeTable)
	}
}

func (rcv *BallInfo) UnPack() *BallInfoT {
	if rcv == nil {
		return nil
	}
	t := &BallInfoT{}
	rcv.UnPackTo(t)
	return t
}

type BallInfo struct {
	_tab flatbuffers.Table
}

func GetRootAsBallInfo(buf []byte, offset flatbuffers.UOffsetT) *BallInfo {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &BallInfo{}
	x.Init(buf, n+offset)
	return x
}

func FinishBallInfoBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsBallInfo(buf []byte, offset flatbuffers.UOffsetT) *BallInfo {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &BallInfo{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedBallInfoBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *BallInfo) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *BallInfo) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *BallInfo) Physics(obj *Physics) *Physics {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(Physics)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *BallInfo) ShapeType() CollisionShape {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return CollisionShape(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *BallInfo) MutateShapeType(n CollisionShape) bool {
	return rcv._tab.MutateByteSlot(6, byte(n))
}

func (rcv *BallInfo) Shape(obj *flatbuffers.Table) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		rcv._tab.Union(obj, o)
		return true
	}
	return false
}

func BallInfoStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func BallInfoAddPhysics(builder *flatbuffers.Builder, physics flatbuffers.UOffsetT) {
	builder.PrependStructSlot(0, flatbuffers.UOffsetT(physics), 0)
}
func BallInfoAddShapeType(builder *flatbuffers.Builder, shapeType CollisionShape) {
	builder.PrependByteSlot(1, byte(shapeType), 0)
}
func BallInfoAddShape(builder *flatbuffers.Builder, shape flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(shape), 0)
}
func BallInfoEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}