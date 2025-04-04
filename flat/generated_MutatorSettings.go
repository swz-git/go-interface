// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flat

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// All mutators options.
type MutatorSettingsT struct {
	MatchLength MatchLengthMutator `json:"match_length"`
	MaxScore MaxScoreMutator `json:"max_score"`
	MultiBall MultiBallMutator `json:"multi_ball"`
	Overtime OvertimeMutator `json:"overtime"`
	SeriesLength SeriesLengthMutator `json:"series_length"`
	GameSpeed GameSpeedMutator `json:"game_speed"`
	BallMaxSpeed BallMaxSpeedMutator `json:"ball_max_speed"`
	BallType BallTypeMutator `json:"ball_type"`
	BallWeight BallWeightMutator `json:"ball_weight"`
	BallSize BallSizeMutator `json:"ball_size"`
	BallBounciness BallBouncinessMutator `json:"ball_bounciness"`
	BoostAmount BoostAmountMutator `json:"boost_amount"`
	Rumble RumbleMutator `json:"rumble"`
	BoostStrength BoostStrengthMutator `json:"boost_strength"`
	Gravity GravityMutator `json:"gravity"`
	Demolish DemolishMutator `json:"demolish"`
	RespawnTime RespawnTimeMutator `json:"respawn_time"`
	MaxTime MaxTimeMutator `json:"max_time"`
	GameEvent GameEventMutator `json:"game_event"`
	Audio AudioMutator `json:"audio"`
	BallGravity BallGravityMutator `json:"ball_gravity"`
	Territory TerritoryMutator `json:"territory"`
	StaleBall StaleBallMutator `json:"stale_ball"`
	Jump JumpMutator `json:"jump"`
	DodgeTimer DodgeTimerMutator `json:"dodge_timer"`
	PossessionScore PossessionScoreMutator `json:"possession_score"`
	DemolishScore DemolishScoreMutator `json:"demolish_score"`
	NormalGoalScore NormalGoalScoreMutator `json:"normal_goal_score"`
	AerialGoalScore AerialGoalScoreMutator `json:"aerial_goal_score"`
	AssistGoalScore AssistGoalScoreMutator `json:"assist_goal_score"`
	InputRestriction InputRestrictionMutator `json:"input_restriction"`
}

func (t *MutatorSettingsT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	MutatorSettingsStart(builder)
	MutatorSettingsAddMatchLength(builder, t.MatchLength)
	MutatorSettingsAddMaxScore(builder, t.MaxScore)
	MutatorSettingsAddMultiBall(builder, t.MultiBall)
	MutatorSettingsAddOvertime(builder, t.Overtime)
	MutatorSettingsAddSeriesLength(builder, t.SeriesLength)
	MutatorSettingsAddGameSpeed(builder, t.GameSpeed)
	MutatorSettingsAddBallMaxSpeed(builder, t.BallMaxSpeed)
	MutatorSettingsAddBallType(builder, t.BallType)
	MutatorSettingsAddBallWeight(builder, t.BallWeight)
	MutatorSettingsAddBallSize(builder, t.BallSize)
	MutatorSettingsAddBallBounciness(builder, t.BallBounciness)
	MutatorSettingsAddBoostAmount(builder, t.BoostAmount)
	MutatorSettingsAddRumble(builder, t.Rumble)
	MutatorSettingsAddBoostStrength(builder, t.BoostStrength)
	MutatorSettingsAddGravity(builder, t.Gravity)
	MutatorSettingsAddDemolish(builder, t.Demolish)
	MutatorSettingsAddRespawnTime(builder, t.RespawnTime)
	MutatorSettingsAddMaxTime(builder, t.MaxTime)
	MutatorSettingsAddGameEvent(builder, t.GameEvent)
	MutatorSettingsAddAudio(builder, t.Audio)
	MutatorSettingsAddBallGravity(builder, t.BallGravity)
	MutatorSettingsAddTerritory(builder, t.Territory)
	MutatorSettingsAddStaleBall(builder, t.StaleBall)
	MutatorSettingsAddJump(builder, t.Jump)
	MutatorSettingsAddDodgeTimer(builder, t.DodgeTimer)
	MutatorSettingsAddPossessionScore(builder, t.PossessionScore)
	MutatorSettingsAddDemolishScore(builder, t.DemolishScore)
	MutatorSettingsAddNormalGoalScore(builder, t.NormalGoalScore)
	MutatorSettingsAddAerialGoalScore(builder, t.AerialGoalScore)
	MutatorSettingsAddAssistGoalScore(builder, t.AssistGoalScore)
	MutatorSettingsAddInputRestriction(builder, t.InputRestriction)
	return MutatorSettingsEnd(builder)
}

func (rcv *MutatorSettings) UnPackTo(t *MutatorSettingsT) {
	t.MatchLength = rcv.MatchLength()
	t.MaxScore = rcv.MaxScore()
	t.MultiBall = rcv.MultiBall()
	t.Overtime = rcv.Overtime()
	t.SeriesLength = rcv.SeriesLength()
	t.GameSpeed = rcv.GameSpeed()
	t.BallMaxSpeed = rcv.BallMaxSpeed()
	t.BallType = rcv.BallType()
	t.BallWeight = rcv.BallWeight()
	t.BallSize = rcv.BallSize()
	t.BallBounciness = rcv.BallBounciness()
	t.BoostAmount = rcv.BoostAmount()
	t.Rumble = rcv.Rumble()
	t.BoostStrength = rcv.BoostStrength()
	t.Gravity = rcv.Gravity()
	t.Demolish = rcv.Demolish()
	t.RespawnTime = rcv.RespawnTime()
	t.MaxTime = rcv.MaxTime()
	t.GameEvent = rcv.GameEvent()
	t.Audio = rcv.Audio()
	t.BallGravity = rcv.BallGravity()
	t.Territory = rcv.Territory()
	t.StaleBall = rcv.StaleBall()
	t.Jump = rcv.Jump()
	t.DodgeTimer = rcv.DodgeTimer()
	t.PossessionScore = rcv.PossessionScore()
	t.DemolishScore = rcv.DemolishScore()
	t.NormalGoalScore = rcv.NormalGoalScore()
	t.AerialGoalScore = rcv.AerialGoalScore()
	t.AssistGoalScore = rcv.AssistGoalScore()
	t.InputRestriction = rcv.InputRestriction()
}

func (rcv *MutatorSettings) UnPack() *MutatorSettingsT {
	if rcv == nil {
		return nil
	}
	t := &MutatorSettingsT{}
	rcv.UnPackTo(t)
	return t
}

type MutatorSettings struct {
	_tab flatbuffers.Table
}

func GetRootAsMutatorSettings(buf []byte, offset flatbuffers.UOffsetT) *MutatorSettings {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &MutatorSettings{}
	x.Init(buf, n+offset)
	return x
}

func FinishMutatorSettingsBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsMutatorSettings(buf []byte, offset flatbuffers.UOffsetT) *MutatorSettings {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &MutatorSettings{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedMutatorSettingsBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *MutatorSettings) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *MutatorSettings) Table() flatbuffers.Table {
	return rcv._tab
}

/// Duration of the match.
func (rcv *MutatorSettings) MatchLength() MatchLengthMutator {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return MatchLengthMutator(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

/// Duration of the match.
func (rcv *MutatorSettings) MutateMatchLength(n MatchLengthMutator) bool {
	return rcv._tab.MutateByteSlot(4, byte(n))
}

/// Max score of match. If this score is reached, the team immediately wins.
func (rcv *MutatorSettings) MaxScore() MaxScoreMutator {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return MaxScoreMutator(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

/// Max score of match. If this score is reached, the team immediately wins.
func (rcv *MutatorSettings) MutateMaxScore(n MaxScoreMutator) bool {
	return rcv._tab.MutateByteSlot(6, byte(n))
}

/// The number of balls.
func (rcv *MutatorSettings) MultiBall() MultiBallMutator {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return MultiBallMutator(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

/// The number of balls.
func (rcv *MutatorSettings) MutateMultiBall(n MultiBallMutator) bool {
	return rcv._tab.MutateByteSlot(8, byte(n))
}

/// The overtime rules and tiebreaker.
func (rcv *MutatorSettings) Overtime() OvertimeMutator {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return OvertimeMutator(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

/// The overtime rules and tiebreaker.
func (rcv *MutatorSettings) MutateOvertime(n OvertimeMutator) bool {
	return rcv._tab.MutateByteSlot(10, byte(n))
}

/// The series length.
func (rcv *MutatorSettings) SeriesLength() SeriesLengthMutator {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return SeriesLengthMutator(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

/// The series length.
func (rcv *MutatorSettings) MutateSeriesLength(n SeriesLengthMutator) bool {
	return rcv._tab.MutateByteSlot(12, byte(n))
}

/// A game speed multiplier.
func (rcv *MutatorSettings) GameSpeed() GameSpeedMutator {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return GameSpeedMutator(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

/// A game speed multiplier.
func (rcv *MutatorSettings) MutateGameSpeed(n GameSpeedMutator) bool {
	return rcv._tab.MutateByteSlot(14, byte(n))
}

/// Ball max speed.
func (rcv *MutatorSettings) BallMaxSpeed() BallMaxSpeedMutator {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return BallMaxSpeedMutator(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

/// Ball max speed.
func (rcv *MutatorSettings) MutateBallMaxSpeed(n BallMaxSpeedMutator) bool {
	return rcv._tab.MutateByteSlot(16, byte(n))
}

/// Ball type and shape.
func (rcv *MutatorSettings) BallType() BallTypeMutator {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return BallTypeMutator(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

/// Ball type and shape.
func (rcv *MutatorSettings) MutateBallType(n BallTypeMutator) bool {
	return rcv._tab.MutateByteSlot(18, byte(n))
}

/// Ball weight and how much is curves.
func (rcv *MutatorSettings) BallWeight() BallWeightMutator {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return BallWeightMutator(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

/// Ball weight and how much is curves.
func (rcv *MutatorSettings) MutateBallWeight(n BallWeightMutator) bool {
	return rcv._tab.MutateByteSlot(20, byte(n))
}

/// Ball size.
func (rcv *MutatorSettings) BallSize() BallSizeMutator {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return BallSizeMutator(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

/// Ball size.
func (rcv *MutatorSettings) MutateBallSize(n BallSizeMutator) bool {
	return rcv._tab.MutateByteSlot(22, byte(n))
}

/// Ball bounciness.
func (rcv *MutatorSettings) BallBounciness() BallBouncinessMutator {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		return BallBouncinessMutator(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

/// Ball bounciness.
func (rcv *MutatorSettings) MutateBallBounciness(n BallBouncinessMutator) bool {
	return rcv._tab.MutateByteSlot(24, byte(n))
}

/// Boost amount/recharge.
func (rcv *MutatorSettings) BoostAmount() BoostAmountMutator {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(26))
	if o != 0 {
		return BoostAmountMutator(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

/// Boost amount/recharge.
func (rcv *MutatorSettings) MutateBoostAmount(n BoostAmountMutator) bool {
	return rcv._tab.MutateByteSlot(26, byte(n))
}

/// Rumble item rules.
func (rcv *MutatorSettings) Rumble() RumbleMutator {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(28))
	if o != 0 {
		return RumbleMutator(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

/// Rumble item rules.
func (rcv *MutatorSettings) MutateRumble(n RumbleMutator) bool {
	return rcv._tab.MutateByteSlot(28, byte(n))
}

/// Boost strength multiplier.
func (rcv *MutatorSettings) BoostStrength() BoostStrengthMutator {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(30))
	if o != 0 {
		return BoostStrengthMutator(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

/// Boost strength multiplier.
func (rcv *MutatorSettings) MutateBoostStrength(n BoostStrengthMutator) bool {
	return rcv._tab.MutateByteSlot(30, byte(n))
}

/// Strength of gravity.
func (rcv *MutatorSettings) Gravity() GravityMutator {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(32))
	if o != 0 {
		return GravityMutator(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

/// Strength of gravity.
func (rcv *MutatorSettings) MutateGravity(n GravityMutator) bool {
	return rcv._tab.MutateByteSlot(32, byte(n))
}

/// Demolition conditions.
func (rcv *MutatorSettings) Demolish() DemolishMutator {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(34))
	if o != 0 {
		return DemolishMutator(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

/// Demolition conditions.
func (rcv *MutatorSettings) MutateDemolish(n DemolishMutator) bool {
	return rcv._tab.MutateByteSlot(34, byte(n))
}

/// Demolition respawn time.
func (rcv *MutatorSettings) RespawnTime() RespawnTimeMutator {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(36))
	if o != 0 {
		return RespawnTimeMutator(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

/// Demolition respawn time.
func (rcv *MutatorSettings) MutateRespawnTime(n RespawnTimeMutator) bool {
	return rcv._tab.MutateByteSlot(36, byte(n))
}

/// Max real-time duration of match including kickoff, replays, and more.
/// If the score is tied upon time-out, the number of shots determine the winner.
func (rcv *MutatorSettings) MaxTime() MaxTimeMutator {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(38))
	if o != 0 {
		return MaxTimeMutator(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

/// Max real-time duration of match including kickoff, replays, and more.
/// If the score is tied upon time-out, the number of shots determine the winner.
func (rcv *MutatorSettings) MutateMaxTime(n MaxTimeMutator) bool {
	return rcv._tab.MutateByteSlot(38, byte(n))
}

/// Additional game behaviour for custom modes.
func (rcv *MutatorSettings) GameEvent() GameEventMutator {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(40))
	if o != 0 {
		return GameEventMutator(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

/// Additional game behaviour for custom modes.
func (rcv *MutatorSettings) MutateGameEvent(n GameEventMutator) bool {
	return rcv._tab.MutateByteSlot(40, byte(n))
}

/// Additional audio options for custom modes.
func (rcv *MutatorSettings) Audio() AudioMutator {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(42))
	if o != 0 {
		return AudioMutator(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

/// Additional audio options for custom modes.
func (rcv *MutatorSettings) MutateAudio(n AudioMutator) bool {
	return rcv._tab.MutateByteSlot(42, byte(n))
}

/// Ball gravity.
func (rcv *MutatorSettings) BallGravity() BallGravityMutator {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(44))
	if o != 0 {
		return BallGravityMutator(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

/// Ball gravity.
func (rcv *MutatorSettings) MutateBallGravity(n BallGravityMutator) bool {
	return rcv._tab.MutateByteSlot(44, byte(n))
}

/// Territory mutator.
func (rcv *MutatorSettings) Territory() TerritoryMutator {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(46))
	if o != 0 {
		return TerritoryMutator(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

/// Territory mutator.
func (rcv *MutatorSettings) MutateTerritory(n TerritoryMutator) bool {
	return rcv._tab.MutateByteSlot(46, byte(n))
}

/// Stale ball mutator.
func (rcv *MutatorSettings) StaleBall() StaleBallMutator {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(48))
	if o != 0 {
		return StaleBallMutator(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

/// Stale ball mutator.
func (rcv *MutatorSettings) MutateStaleBall(n StaleBallMutator) bool {
	return rcv._tab.MutateByteSlot(48, byte(n))
}

/// Jumps mutator.
func (rcv *MutatorSettings) Jump() JumpMutator {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(50))
	if o != 0 {
		return JumpMutator(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

/// Jumps mutator.
func (rcv *MutatorSettings) MutateJump(n JumpMutator) bool {
	return rcv._tab.MutateByteSlot(50, byte(n))
}

/// Dodge timer mutator.
func (rcv *MutatorSettings) DodgeTimer() DodgeTimerMutator {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(52))
	if o != 0 {
		return DodgeTimerMutator(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

/// Dodge timer mutator.
func (rcv *MutatorSettings) MutateDodgeTimer(n DodgeTimerMutator) bool {
	return rcv._tab.MutateByteSlot(52, byte(n))
}

/// Possession score mutator.
func (rcv *MutatorSettings) PossessionScore() PossessionScoreMutator {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(54))
	if o != 0 {
		return PossessionScoreMutator(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

/// Possession score mutator.
func (rcv *MutatorSettings) MutatePossessionScore(n PossessionScoreMutator) bool {
	return rcv._tab.MutateByteSlot(54, byte(n))
}

/// Demolish score mutator.
func (rcv *MutatorSettings) DemolishScore() DemolishScoreMutator {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(56))
	if o != 0 {
		return DemolishScoreMutator(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

/// Demolish score mutator.
func (rcv *MutatorSettings) MutateDemolishScore(n DemolishScoreMutator) bool {
	return rcv._tab.MutateByteSlot(56, byte(n))
}

/// Normal goal score mutator.
func (rcv *MutatorSettings) NormalGoalScore() NormalGoalScoreMutator {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(58))
	if o != 0 {
		return NormalGoalScoreMutator(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

/// Normal goal score mutator.
func (rcv *MutatorSettings) MutateNormalGoalScore(n NormalGoalScoreMutator) bool {
	return rcv._tab.MutateByteSlot(58, byte(n))
}

/// Aerial goal score mutator.
func (rcv *MutatorSettings) AerialGoalScore() AerialGoalScoreMutator {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(60))
	if o != 0 {
		return AerialGoalScoreMutator(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

/// Aerial goal score mutator.
func (rcv *MutatorSettings) MutateAerialGoalScore(n AerialGoalScoreMutator) bool {
	return rcv._tab.MutateByteSlot(60, byte(n))
}

/// Assist goal score mutator.
func (rcv *MutatorSettings) AssistGoalScore() AssistGoalScoreMutator {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(62))
	if o != 0 {
		return AssistGoalScoreMutator(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

/// Assist goal score mutator.
func (rcv *MutatorSettings) MutateAssistGoalScore(n AssistGoalScoreMutator) bool {
	return rcv._tab.MutateByteSlot(62, byte(n))
}

/// Player input restriction mutator.
func (rcv *MutatorSettings) InputRestriction() InputRestrictionMutator {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(64))
	if o != 0 {
		return InputRestrictionMutator(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

/// Player input restriction mutator.
func (rcv *MutatorSettings) MutateInputRestriction(n InputRestrictionMutator) bool {
	return rcv._tab.MutateByteSlot(64, byte(n))
}

func MutatorSettingsStart(builder *flatbuffers.Builder) {
	builder.StartObject(31)
}
func MutatorSettingsAddMatchLength(builder *flatbuffers.Builder, matchLength MatchLengthMutator) {
	builder.PrependByteSlot(0, byte(matchLength), 0)
}
func MutatorSettingsAddMaxScore(builder *flatbuffers.Builder, maxScore MaxScoreMutator) {
	builder.PrependByteSlot(1, byte(maxScore), 0)
}
func MutatorSettingsAddMultiBall(builder *flatbuffers.Builder, multiBall MultiBallMutator) {
	builder.PrependByteSlot(2, byte(multiBall), 0)
}
func MutatorSettingsAddOvertime(builder *flatbuffers.Builder, overtime OvertimeMutator) {
	builder.PrependByteSlot(3, byte(overtime), 0)
}
func MutatorSettingsAddSeriesLength(builder *flatbuffers.Builder, seriesLength SeriesLengthMutator) {
	builder.PrependByteSlot(4, byte(seriesLength), 0)
}
func MutatorSettingsAddGameSpeed(builder *flatbuffers.Builder, gameSpeed GameSpeedMutator) {
	builder.PrependByteSlot(5, byte(gameSpeed), 0)
}
func MutatorSettingsAddBallMaxSpeed(builder *flatbuffers.Builder, ballMaxSpeed BallMaxSpeedMutator) {
	builder.PrependByteSlot(6, byte(ballMaxSpeed), 0)
}
func MutatorSettingsAddBallType(builder *flatbuffers.Builder, ballType BallTypeMutator) {
	builder.PrependByteSlot(7, byte(ballType), 0)
}
func MutatorSettingsAddBallWeight(builder *flatbuffers.Builder, ballWeight BallWeightMutator) {
	builder.PrependByteSlot(8, byte(ballWeight), 0)
}
func MutatorSettingsAddBallSize(builder *flatbuffers.Builder, ballSize BallSizeMutator) {
	builder.PrependByteSlot(9, byte(ballSize), 0)
}
func MutatorSettingsAddBallBounciness(builder *flatbuffers.Builder, ballBounciness BallBouncinessMutator) {
	builder.PrependByteSlot(10, byte(ballBounciness), 0)
}
func MutatorSettingsAddBoostAmount(builder *flatbuffers.Builder, boostAmount BoostAmountMutator) {
	builder.PrependByteSlot(11, byte(boostAmount), 0)
}
func MutatorSettingsAddRumble(builder *flatbuffers.Builder, rumble RumbleMutator) {
	builder.PrependByteSlot(12, byte(rumble), 0)
}
func MutatorSettingsAddBoostStrength(builder *flatbuffers.Builder, boostStrength BoostStrengthMutator) {
	builder.PrependByteSlot(13, byte(boostStrength), 0)
}
func MutatorSettingsAddGravity(builder *flatbuffers.Builder, gravity GravityMutator) {
	builder.PrependByteSlot(14, byte(gravity), 0)
}
func MutatorSettingsAddDemolish(builder *flatbuffers.Builder, demolish DemolishMutator) {
	builder.PrependByteSlot(15, byte(demolish), 0)
}
func MutatorSettingsAddRespawnTime(builder *flatbuffers.Builder, respawnTime RespawnTimeMutator) {
	builder.PrependByteSlot(16, byte(respawnTime), 0)
}
func MutatorSettingsAddMaxTime(builder *flatbuffers.Builder, maxTime MaxTimeMutator) {
	builder.PrependByteSlot(17, byte(maxTime), 0)
}
func MutatorSettingsAddGameEvent(builder *flatbuffers.Builder, gameEvent GameEventMutator) {
	builder.PrependByteSlot(18, byte(gameEvent), 0)
}
func MutatorSettingsAddAudio(builder *flatbuffers.Builder, audio AudioMutator) {
	builder.PrependByteSlot(19, byte(audio), 0)
}
func MutatorSettingsAddBallGravity(builder *flatbuffers.Builder, ballGravity BallGravityMutator) {
	builder.PrependByteSlot(20, byte(ballGravity), 0)
}
func MutatorSettingsAddTerritory(builder *flatbuffers.Builder, territory TerritoryMutator) {
	builder.PrependByteSlot(21, byte(territory), 0)
}
func MutatorSettingsAddStaleBall(builder *flatbuffers.Builder, staleBall StaleBallMutator) {
	builder.PrependByteSlot(22, byte(staleBall), 0)
}
func MutatorSettingsAddJump(builder *flatbuffers.Builder, jump JumpMutator) {
	builder.PrependByteSlot(23, byte(jump), 0)
}
func MutatorSettingsAddDodgeTimer(builder *flatbuffers.Builder, dodgeTimer DodgeTimerMutator) {
	builder.PrependByteSlot(24, byte(dodgeTimer), 0)
}
func MutatorSettingsAddPossessionScore(builder *flatbuffers.Builder, possessionScore PossessionScoreMutator) {
	builder.PrependByteSlot(25, byte(possessionScore), 0)
}
func MutatorSettingsAddDemolishScore(builder *flatbuffers.Builder, demolishScore DemolishScoreMutator) {
	builder.PrependByteSlot(26, byte(demolishScore), 0)
}
func MutatorSettingsAddNormalGoalScore(builder *flatbuffers.Builder, normalGoalScore NormalGoalScoreMutator) {
	builder.PrependByteSlot(27, byte(normalGoalScore), 0)
}
func MutatorSettingsAddAerialGoalScore(builder *flatbuffers.Builder, aerialGoalScore AerialGoalScoreMutator) {
	builder.PrependByteSlot(28, byte(aerialGoalScore), 0)
}
func MutatorSettingsAddAssistGoalScore(builder *flatbuffers.Builder, assistGoalScore AssistGoalScoreMutator) {
	builder.PrependByteSlot(29, byte(assistGoalScore), 0)
}
func MutatorSettingsAddInputRestriction(builder *flatbuffers.Builder, inputRestriction InputRestrictionMutator) {
	builder.PrependByteSlot(30, byte(inputRestriction), 0)
}
func MutatorSettingsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
