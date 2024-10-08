// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flat

import (
	flatbuffers "github.com/google/flatbuffers/go"
	"strconv"
)

type CollisionShape byte

const (
	CollisionShapeNONE          CollisionShape = 0
	CollisionShapeBoxShape      CollisionShape = 1
	CollisionShapeSphereShape   CollisionShape = 2
	CollisionShapeCylinderShape CollisionShape = 3
)

var EnumNamesCollisionShape = map[CollisionShape]string{
	CollisionShapeNONE:          "NONE",
	CollisionShapeBoxShape:      "BoxShape",
	CollisionShapeSphereShape:   "SphereShape",
	CollisionShapeCylinderShape: "CylinderShape",
}

var EnumValuesCollisionShape = map[string]CollisionShape{
	"NONE":          CollisionShapeNONE,
	"BoxShape":      CollisionShapeBoxShape,
	"SphereShape":   CollisionShapeSphereShape,
	"CylinderShape": CollisionShapeCylinderShape,
}

func (v CollisionShape) String() string {
	if s, ok := EnumNamesCollisionShape[v]; ok {
		return s
	}
	return "CollisionShape(" + strconv.FormatInt(int64(v), 10) + ")"
}

type CollisionShapeT struct {
	Type CollisionShape
	Value interface{}
}

func (t *CollisionShapeT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	switch t.Type {
	case CollisionShapeBoxShape:
		return t.Value.(*BoxShapeT).Pack(builder)
	case CollisionShapeSphereShape:
		return t.Value.(*SphereShapeT).Pack(builder)
	case CollisionShapeCylinderShape:
		return t.Value.(*CylinderShapeT).Pack(builder)
	}
	return 0
}

func (rcv CollisionShape) UnPack(table flatbuffers.Table) *CollisionShapeT {
	switch rcv {
	case CollisionShapeBoxShape:
		var x BoxShape
		x.Init(table.Bytes, table.Pos)
		return &CollisionShapeT{Type: CollisionShapeBoxShape, Value: x.UnPack()}
	case CollisionShapeSphereShape:
		var x SphereShape
		x.Init(table.Bytes, table.Pos)
		return &CollisionShapeT{Type: CollisionShapeSphereShape, Value: x.UnPack()}
	case CollisionShapeCylinderShape:
		var x CylinderShape
		x.Init(table.Bytes, table.Pos)
		return &CollisionShapeT{Type: CollisionShapeCylinderShape, Value: x.UnPack()}
	}
	return nil
}
