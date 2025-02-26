// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flat

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// An entry in the ball prediction describing where a ball will be at some future time.
type PredictionSliceT struct {
	GameSeconds float32 `json:"game_seconds"`
	Physics *PhysicsT `json:"physics"`
}

func (t *PredictionSliceT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	return CreatePredictionSlice(builder, t.GameSeconds, t.Physics.Location.X, t.Physics.Location.Y, t.Physics.Location.Z, t.Physics.Rotation.Pitch, t.Physics.Rotation.Yaw, t.Physics.Rotation.Roll, t.Physics.Velocity.X, t.Physics.Velocity.Y, t.Physics.Velocity.Z, t.Physics.AngularVelocity.X, t.Physics.AngularVelocity.Y, t.Physics.AngularVelocity.Z)
}
func (rcv *PredictionSlice) UnPackTo(t *PredictionSliceT) {
	t.GameSeconds = rcv.GameSeconds()
	t.Physics = rcv.Physics(nil).UnPack()
}

func (rcv *PredictionSlice) UnPack() *PredictionSliceT {
	if rcv == nil {
		return nil
	}
	t := &PredictionSliceT{}
	rcv.UnPackTo(t)
	return t
}

type PredictionSlice struct {
	_tab flatbuffers.Struct
}

func (rcv *PredictionSlice) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *PredictionSlice) Table() flatbuffers.Table {
	return rcv._tab.Table
}

/// The moment in game time that this prediction corresponds to.
/// This corresponds to 'seconds_elapsed' in the MatchInfo.
func (rcv *PredictionSlice) GameSeconds() float32 {
	return rcv._tab.GetFloat32(rcv._tab.Pos + flatbuffers.UOffsetT(0))
}
/// The moment in game time that this prediction corresponds to.
/// This corresponds to 'seconds_elapsed' in the MatchInfo.
func (rcv *PredictionSlice) MutateGameSeconds(n float32) bool {
	return rcv._tab.MutateFloat32(rcv._tab.Pos+flatbuffers.UOffsetT(0), n)
}

/// The predicted location and motion of the object.
func (rcv *PredictionSlice) Physics(obj *Physics) *Physics {
	if obj == nil {
		obj = new(Physics)
	}
	obj.Init(rcv._tab.Bytes, rcv._tab.Pos+4)
	return obj
}
/// The predicted location and motion of the object.

func CreatePredictionSlice(builder *flatbuffers.Builder, gameSeconds float32, physics_location_x float32, physics_location_y float32, physics_location_z float32, physics_rotation_pitch float32, physics_rotation_yaw float32, physics_rotation_roll float32, physics_velocity_x float32, physics_velocity_y float32, physics_velocity_z float32, physics_angular_velocity_x float32, physics_angular_velocity_y float32, physics_angular_velocity_z float32) flatbuffers.UOffsetT {
	builder.Prep(4, 52)
	builder.Prep(4, 48)
	builder.Prep(4, 12)
	builder.PrependFloat32(physics_angular_velocity_z)
	builder.PrependFloat32(physics_angular_velocity_y)
	builder.PrependFloat32(physics_angular_velocity_x)
	builder.Prep(4, 12)
	builder.PrependFloat32(physics_velocity_z)
	builder.PrependFloat32(physics_velocity_y)
	builder.PrependFloat32(physics_velocity_x)
	builder.Prep(4, 12)
	builder.PrependFloat32(physics_rotation_roll)
	builder.PrependFloat32(physics_rotation_yaw)
	builder.PrependFloat32(physics_rotation_pitch)
	builder.Prep(4, 12)
	builder.PrependFloat32(physics_location_z)
	builder.PrependFloat32(physics_location_y)
	builder.PrependFloat32(physics_location_x)
	builder.PrependFloat32(gameSeconds)
	return builder.Offset()
}
