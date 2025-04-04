// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flat

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// Defines the car type, color, and other aspects of the player's appearance.
/// See https://wiki.rlbot.org/botmaking/bot-customization/
type PlayerLoadoutT struct {
	TeamColorId uint32 `json:"team_color_id"`
	CustomColorId uint32 `json:"custom_color_id"`
	CarId uint32 `json:"car_id"`
	DecalId uint32 `json:"decal_id"`
	WheelsId uint32 `json:"wheels_id"`
	BoostId uint32 `json:"boost_id"`
	AntennaId uint32 `json:"antenna_id"`
	HatId uint32 `json:"hat_id"`
	PaintFinishId uint32 `json:"paint_finish_id"`
	CustomFinishId uint32 `json:"custom_finish_id"`
	EngineAudioId uint32 `json:"engine_audio_id"`
	TrailsId uint32 `json:"trails_id"`
	GoalExplosionId uint32 `json:"goal_explosion_id"`
	LoadoutPaint *LoadoutPaintT `json:"loadout_paint"`
	PrimaryColorLookup *ColorT `json:"primary_color_lookup"`
	SecondaryColorLookup *ColorT `json:"secondary_color_lookup"`
}

func (t *PlayerLoadoutT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	loadoutPaintOffset := t.LoadoutPaint.Pack(builder)
	PlayerLoadoutStart(builder)
	PlayerLoadoutAddTeamColorId(builder, t.TeamColorId)
	PlayerLoadoutAddCustomColorId(builder, t.CustomColorId)
	PlayerLoadoutAddCarId(builder, t.CarId)
	PlayerLoadoutAddDecalId(builder, t.DecalId)
	PlayerLoadoutAddWheelsId(builder, t.WheelsId)
	PlayerLoadoutAddBoostId(builder, t.BoostId)
	PlayerLoadoutAddAntennaId(builder, t.AntennaId)
	PlayerLoadoutAddHatId(builder, t.HatId)
	PlayerLoadoutAddPaintFinishId(builder, t.PaintFinishId)
	PlayerLoadoutAddCustomFinishId(builder, t.CustomFinishId)
	PlayerLoadoutAddEngineAudioId(builder, t.EngineAudioId)
	PlayerLoadoutAddTrailsId(builder, t.TrailsId)
	PlayerLoadoutAddGoalExplosionId(builder, t.GoalExplosionId)
	PlayerLoadoutAddLoadoutPaint(builder, loadoutPaintOffset)
	primaryColorLookupOffset := t.PrimaryColorLookup.Pack(builder)
	PlayerLoadoutAddPrimaryColorLookup(builder, primaryColorLookupOffset)
	secondaryColorLookupOffset := t.SecondaryColorLookup.Pack(builder)
	PlayerLoadoutAddSecondaryColorLookup(builder, secondaryColorLookupOffset)
	return PlayerLoadoutEnd(builder)
}

func (rcv *PlayerLoadout) UnPackTo(t *PlayerLoadoutT) {
	t.TeamColorId = rcv.TeamColorId()
	t.CustomColorId = rcv.CustomColorId()
	t.CarId = rcv.CarId()
	t.DecalId = rcv.DecalId()
	t.WheelsId = rcv.WheelsId()
	t.BoostId = rcv.BoostId()
	t.AntennaId = rcv.AntennaId()
	t.HatId = rcv.HatId()
	t.PaintFinishId = rcv.PaintFinishId()
	t.CustomFinishId = rcv.CustomFinishId()
	t.EngineAudioId = rcv.EngineAudioId()
	t.TrailsId = rcv.TrailsId()
	t.GoalExplosionId = rcv.GoalExplosionId()
	t.LoadoutPaint = rcv.LoadoutPaint(nil).UnPack()
	t.PrimaryColorLookup = rcv.PrimaryColorLookup(nil).UnPack()
	t.SecondaryColorLookup = rcv.SecondaryColorLookup(nil).UnPack()
}

func (rcv *PlayerLoadout) UnPack() *PlayerLoadoutT {
	if rcv == nil {
		return nil
	}
	t := &PlayerLoadoutT{}
	rcv.UnPackTo(t)
	return t
}

type PlayerLoadout struct {
	_tab flatbuffers.Table
}

func GetRootAsPlayerLoadout(buf []byte, offset flatbuffers.UOffsetT) *PlayerLoadout {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &PlayerLoadout{}
	x.Init(buf, n+offset)
	return x
}

func FinishPlayerLoadoutBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsPlayerLoadout(buf []byte, offset flatbuffers.UOffsetT) *PlayerLoadout {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &PlayerLoadout{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedPlayerLoadoutBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *PlayerLoadout) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *PlayerLoadout) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *PlayerLoadout) TeamColorId() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *PlayerLoadout) MutateTeamColorId(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

func (rcv *PlayerLoadout) CustomColorId() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *PlayerLoadout) MutateCustomColorId(n uint32) bool {
	return rcv._tab.MutateUint32Slot(6, n)
}

func (rcv *PlayerLoadout) CarId() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *PlayerLoadout) MutateCarId(n uint32) bool {
	return rcv._tab.MutateUint32Slot(8, n)
}

func (rcv *PlayerLoadout) DecalId() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *PlayerLoadout) MutateDecalId(n uint32) bool {
	return rcv._tab.MutateUint32Slot(10, n)
}

func (rcv *PlayerLoadout) WheelsId() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *PlayerLoadout) MutateWheelsId(n uint32) bool {
	return rcv._tab.MutateUint32Slot(12, n)
}

func (rcv *PlayerLoadout) BoostId() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *PlayerLoadout) MutateBoostId(n uint32) bool {
	return rcv._tab.MutateUint32Slot(14, n)
}

func (rcv *PlayerLoadout) AntennaId() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *PlayerLoadout) MutateAntennaId(n uint32) bool {
	return rcv._tab.MutateUint32Slot(16, n)
}

func (rcv *PlayerLoadout) HatId() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *PlayerLoadout) MutateHatId(n uint32) bool {
	return rcv._tab.MutateUint32Slot(18, n)
}

func (rcv *PlayerLoadout) PaintFinishId() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *PlayerLoadout) MutatePaintFinishId(n uint32) bool {
	return rcv._tab.MutateUint32Slot(20, n)
}

func (rcv *PlayerLoadout) CustomFinishId() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *PlayerLoadout) MutateCustomFinishId(n uint32) bool {
	return rcv._tab.MutateUint32Slot(22, n)
}

func (rcv *PlayerLoadout) EngineAudioId() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *PlayerLoadout) MutateEngineAudioId(n uint32) bool {
	return rcv._tab.MutateUint32Slot(24, n)
}

func (rcv *PlayerLoadout) TrailsId() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(26))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *PlayerLoadout) MutateTrailsId(n uint32) bool {
	return rcv._tab.MutateUint32Slot(26, n)
}

func (rcv *PlayerLoadout) GoalExplosionId() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(28))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *PlayerLoadout) MutateGoalExplosionId(n uint32) bool {
	return rcv._tab.MutateUint32Slot(28, n)
}

func (rcv *PlayerLoadout) LoadoutPaint(obj *LoadoutPaint) *LoadoutPaint {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(30))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(LoadoutPaint)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

/// Sets the primary color of the car to the swatch that most closely matches the provided
/// RGB color value. If set, this overrides teamColorId.
func (rcv *PlayerLoadout) PrimaryColorLookup(obj *Color) *Color {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(32))
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

/// Sets the primary color of the car to the swatch that most closely matches the provided
/// RGB color value. If set, this overrides teamColorId.
/// Sets the secondary color of the car to the swatch that most closely matches the provided
/// RGB color value. If set, this overrides customColorId.
func (rcv *PlayerLoadout) SecondaryColorLookup(obj *Color) *Color {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(34))
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

/// Sets the secondary color of the car to the swatch that most closely matches the provided
/// RGB color value. If set, this overrides customColorId.
func PlayerLoadoutStart(builder *flatbuffers.Builder) {
	builder.StartObject(16)
}
func PlayerLoadoutAddTeamColorId(builder *flatbuffers.Builder, teamColorId uint32) {
	builder.PrependUint32Slot(0, teamColorId, 0)
}
func PlayerLoadoutAddCustomColorId(builder *flatbuffers.Builder, customColorId uint32) {
	builder.PrependUint32Slot(1, customColorId, 0)
}
func PlayerLoadoutAddCarId(builder *flatbuffers.Builder, carId uint32) {
	builder.PrependUint32Slot(2, carId, 0)
}
func PlayerLoadoutAddDecalId(builder *flatbuffers.Builder, decalId uint32) {
	builder.PrependUint32Slot(3, decalId, 0)
}
func PlayerLoadoutAddWheelsId(builder *flatbuffers.Builder, wheelsId uint32) {
	builder.PrependUint32Slot(4, wheelsId, 0)
}
func PlayerLoadoutAddBoostId(builder *flatbuffers.Builder, boostId uint32) {
	builder.PrependUint32Slot(5, boostId, 0)
}
func PlayerLoadoutAddAntennaId(builder *flatbuffers.Builder, antennaId uint32) {
	builder.PrependUint32Slot(6, antennaId, 0)
}
func PlayerLoadoutAddHatId(builder *flatbuffers.Builder, hatId uint32) {
	builder.PrependUint32Slot(7, hatId, 0)
}
func PlayerLoadoutAddPaintFinishId(builder *flatbuffers.Builder, paintFinishId uint32) {
	builder.PrependUint32Slot(8, paintFinishId, 0)
}
func PlayerLoadoutAddCustomFinishId(builder *flatbuffers.Builder, customFinishId uint32) {
	builder.PrependUint32Slot(9, customFinishId, 0)
}
func PlayerLoadoutAddEngineAudioId(builder *flatbuffers.Builder, engineAudioId uint32) {
	builder.PrependUint32Slot(10, engineAudioId, 0)
}
func PlayerLoadoutAddTrailsId(builder *flatbuffers.Builder, trailsId uint32) {
	builder.PrependUint32Slot(11, trailsId, 0)
}
func PlayerLoadoutAddGoalExplosionId(builder *flatbuffers.Builder, goalExplosionId uint32) {
	builder.PrependUint32Slot(12, goalExplosionId, 0)
}
func PlayerLoadoutAddLoadoutPaint(builder *flatbuffers.Builder, loadoutPaint flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(13, flatbuffers.UOffsetT(loadoutPaint), 0)
}
func PlayerLoadoutAddPrimaryColorLookup(builder *flatbuffers.Builder, primaryColorLookup flatbuffers.UOffsetT) {
	builder.PrependStructSlot(14, flatbuffers.UOffsetT(primaryColorLookup), 0)
}
func PlayerLoadoutAddSecondaryColorLookup(builder *flatbuffers.Builder, secondaryColorLookup flatbuffers.UOffsetT) {
	builder.PrependStructSlot(15, flatbuffers.UOffsetT(secondaryColorLookup), 0)
}
func PlayerLoadoutEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
