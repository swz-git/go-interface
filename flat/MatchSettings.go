// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flat

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type MatchSettingsT struct {
	Launcher Launcher `json:"launcher"`
	GamePath string `json:"game_path"`
	AutoStartBots bool `json:"auto_start_bots"`
	GameMapUpk string `json:"game_map_upk"`
	PlayerConfigurations []*PlayerConfigurationT `json:"player_configurations"`
	ScriptConfigurations []*ScriptConfigurationT `json:"script_configurations"`
	GameMode GameMode `json:"game_mode"`
	SkipReplays bool `json:"skip_replays"`
	InstantStart bool `json:"instant_start"`
	MutatorSettings *MutatorSettingsT `json:"mutator_settings"`
	ExistingMatchBehavior ExistingMatchBehavior `json:"existing_match_behavior"`
	EnableRendering bool `json:"enable_rendering"`
	EnableStateSetting bool `json:"enable_state_setting"`
	AutoSaveReplay bool `json:"auto_save_replay"`
	Freeplay bool `json:"freeplay"`
}

func (t *MatchSettingsT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	gamePathOffset := flatbuffers.UOffsetT(0)
	if t.GamePath != "" {
		gamePathOffset = builder.CreateString(t.GamePath)
	}
	gameMapUpkOffset := flatbuffers.UOffsetT(0)
	if t.GameMapUpk != "" {
		gameMapUpkOffset = builder.CreateString(t.GameMapUpk)
	}
	playerConfigurationsOffset := flatbuffers.UOffsetT(0)
	if t.PlayerConfigurations != nil {
		playerConfigurationsLength := len(t.PlayerConfigurations)
		playerConfigurationsOffsets := make([]flatbuffers.UOffsetT, playerConfigurationsLength)
		for j := 0; j < playerConfigurationsLength; j++ {
			playerConfigurationsOffsets[j] = t.PlayerConfigurations[j].Pack(builder)
		}
		MatchSettingsStartPlayerConfigurationsVector(builder, playerConfigurationsLength)
		for j := playerConfigurationsLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(playerConfigurationsOffsets[j])
		}
		playerConfigurationsOffset = builder.EndVector(playerConfigurationsLength)
	}
	scriptConfigurationsOffset := flatbuffers.UOffsetT(0)
	if t.ScriptConfigurations != nil {
		scriptConfigurationsLength := len(t.ScriptConfigurations)
		scriptConfigurationsOffsets := make([]flatbuffers.UOffsetT, scriptConfigurationsLength)
		for j := 0; j < scriptConfigurationsLength; j++ {
			scriptConfigurationsOffsets[j] = t.ScriptConfigurations[j].Pack(builder)
		}
		MatchSettingsStartScriptConfigurationsVector(builder, scriptConfigurationsLength)
		for j := scriptConfigurationsLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(scriptConfigurationsOffsets[j])
		}
		scriptConfigurationsOffset = builder.EndVector(scriptConfigurationsLength)
	}
	mutatorSettingsOffset := t.MutatorSettings.Pack(builder)
	MatchSettingsStart(builder)
	MatchSettingsAddLauncher(builder, t.Launcher)
	MatchSettingsAddGamePath(builder, gamePathOffset)
	MatchSettingsAddAutoStartBots(builder, t.AutoStartBots)
	MatchSettingsAddGameMapUpk(builder, gameMapUpkOffset)
	MatchSettingsAddPlayerConfigurations(builder, playerConfigurationsOffset)
	MatchSettingsAddScriptConfigurations(builder, scriptConfigurationsOffset)
	MatchSettingsAddGameMode(builder, t.GameMode)
	MatchSettingsAddSkipReplays(builder, t.SkipReplays)
	MatchSettingsAddInstantStart(builder, t.InstantStart)
	MatchSettingsAddMutatorSettings(builder, mutatorSettingsOffset)
	MatchSettingsAddExistingMatchBehavior(builder, t.ExistingMatchBehavior)
	MatchSettingsAddEnableRendering(builder, t.EnableRendering)
	MatchSettingsAddEnableStateSetting(builder, t.EnableStateSetting)
	MatchSettingsAddAutoSaveReplay(builder, t.AutoSaveReplay)
	MatchSettingsAddFreeplay(builder, t.Freeplay)
	return MatchSettingsEnd(builder)
}

func (rcv *MatchSettings) UnPackTo(t *MatchSettingsT) {
	t.Launcher = rcv.Launcher()
	t.GamePath = string(rcv.GamePath())
	t.AutoStartBots = rcv.AutoStartBots()
	t.GameMapUpk = string(rcv.GameMapUpk())
	playerConfigurationsLength := rcv.PlayerConfigurationsLength()
	t.PlayerConfigurations = make([]*PlayerConfigurationT, playerConfigurationsLength)
	for j := 0; j < playerConfigurationsLength; j++ {
		x := PlayerConfiguration{}
		rcv.PlayerConfigurations(&x, j)
		t.PlayerConfigurations[j] = x.UnPack()
	}
	scriptConfigurationsLength := rcv.ScriptConfigurationsLength()
	t.ScriptConfigurations = make([]*ScriptConfigurationT, scriptConfigurationsLength)
	for j := 0; j < scriptConfigurationsLength; j++ {
		x := ScriptConfiguration{}
		rcv.ScriptConfigurations(&x, j)
		t.ScriptConfigurations[j] = x.UnPack()
	}
	t.GameMode = rcv.GameMode()
	t.SkipReplays = rcv.SkipReplays()
	t.InstantStart = rcv.InstantStart()
	t.MutatorSettings = rcv.MutatorSettings(nil).UnPack()
	t.ExistingMatchBehavior = rcv.ExistingMatchBehavior()
	t.EnableRendering = rcv.EnableRendering()
	t.EnableStateSetting = rcv.EnableStateSetting()
	t.AutoSaveReplay = rcv.AutoSaveReplay()
	t.Freeplay = rcv.Freeplay()
}

func (rcv *MatchSettings) UnPack() *MatchSettingsT {
	if rcv == nil {
		return nil
	}
	t := &MatchSettingsT{}
	rcv.UnPackTo(t)
	return t
}

type MatchSettings struct {
	_tab flatbuffers.Table
}

func GetRootAsMatchSettings(buf []byte, offset flatbuffers.UOffsetT) *MatchSettings {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &MatchSettings{}
	x.Init(buf, n+offset)
	return x
}

func FinishMatchSettingsBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsMatchSettings(buf []byte, offset flatbuffers.UOffsetT) *MatchSettings {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &MatchSettings{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedMatchSettingsBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *MatchSettings) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *MatchSettings) Table() flatbuffers.Table {
	return rcv._tab
}

/// Leave unset to tell RLBot to not launch the game
func (rcv *MatchSettings) Launcher() Launcher {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return Launcher(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

/// Leave unset to tell RLBot to not launch the game
func (rcv *MatchSettings) MutateLauncher(n Launcher) bool {
	return rcv._tab.MutateByteSlot(4, byte(n))
}

/// The path to the main Rocket League binary
func (rcv *MatchSettings) GamePath() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// The path to the main Rocket League binary
func (rcv *MatchSettings) AutoStartBots() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *MatchSettings) MutateAutoStartBots(n bool) bool {
	return rcv._tab.MutateBoolSlot(8, n)
}

/// The name of a upk file, like UtopiaStadium_P, which should be loaded.
/// On Steam version of Rocket League this can be used to load custom map files,
/// but on Epic version it only works on the Psyonix maps.
func (rcv *MatchSettings) GameMapUpk() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// The name of a upk file, like UtopiaStadium_P, which should be loaded.
/// On Steam version of Rocket League this can be used to load custom map files,
/// but on Epic version it only works on the Psyonix maps.
func (rcv *MatchSettings) PlayerConfigurations(obj *PlayerConfiguration, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *MatchSettings) PlayerConfigurationsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *MatchSettings) ScriptConfigurations(obj *ScriptConfiguration, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *MatchSettings) ScriptConfigurationsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *MatchSettings) GameMode() GameMode {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return GameMode(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *MatchSettings) MutateGameMode(n GameMode) bool {
	return rcv._tab.MutateByteSlot(16, byte(n))
}

func (rcv *MatchSettings) SkipReplays() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *MatchSettings) MutateSkipReplays(n bool) bool {
	return rcv._tab.MutateBoolSlot(18, n)
}

func (rcv *MatchSettings) InstantStart() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *MatchSettings) MutateInstantStart(n bool) bool {
	return rcv._tab.MutateBoolSlot(20, n)
}

func (rcv *MatchSettings) MutatorSettings(obj *MutatorSettings) *MutatorSettings {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(MutatorSettings)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *MatchSettings) ExistingMatchBehavior() ExistingMatchBehavior {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		return ExistingMatchBehavior(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *MatchSettings) MutateExistingMatchBehavior(n ExistingMatchBehavior) bool {
	return rcv._tab.MutateByteSlot(24, byte(n))
}

func (rcv *MatchSettings) EnableRendering() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(26))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *MatchSettings) MutateEnableRendering(n bool) bool {
	return rcv._tab.MutateBoolSlot(26, n)
}

func (rcv *MatchSettings) EnableStateSetting() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(28))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *MatchSettings) MutateEnableStateSetting(n bool) bool {
	return rcv._tab.MutateBoolSlot(28, n)
}

func (rcv *MatchSettings) AutoSaveReplay() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(30))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *MatchSettings) MutateAutoSaveReplay(n bool) bool {
	return rcv._tab.MutateBoolSlot(30, n)
}

func (rcv *MatchSettings) Freeplay() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(32))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *MatchSettings) MutateFreeplay(n bool) bool {
	return rcv._tab.MutateBoolSlot(32, n)
}

func MatchSettingsStart(builder *flatbuffers.Builder) {
	builder.StartObject(15)
}
func MatchSettingsAddLauncher(builder *flatbuffers.Builder, launcher Launcher) {
	builder.PrependByteSlot(0, byte(launcher), 0)
}
func MatchSettingsAddGamePath(builder *flatbuffers.Builder, gamePath flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(gamePath), 0)
}
func MatchSettingsAddAutoStartBots(builder *flatbuffers.Builder, autoStartBots bool) {
	builder.PrependBoolSlot(2, autoStartBots, false)
}
func MatchSettingsAddGameMapUpk(builder *flatbuffers.Builder, gameMapUpk flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(gameMapUpk), 0)
}
func MatchSettingsAddPlayerConfigurations(builder *flatbuffers.Builder, playerConfigurations flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(playerConfigurations), 0)
}
func MatchSettingsStartPlayerConfigurationsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func MatchSettingsAddScriptConfigurations(builder *flatbuffers.Builder, scriptConfigurations flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(scriptConfigurations), 0)
}
func MatchSettingsStartScriptConfigurationsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func MatchSettingsAddGameMode(builder *flatbuffers.Builder, gameMode GameMode) {
	builder.PrependByteSlot(6, byte(gameMode), 0)
}
func MatchSettingsAddSkipReplays(builder *flatbuffers.Builder, skipReplays bool) {
	builder.PrependBoolSlot(7, skipReplays, false)
}
func MatchSettingsAddInstantStart(builder *flatbuffers.Builder, instantStart bool) {
	builder.PrependBoolSlot(8, instantStart, false)
}
func MatchSettingsAddMutatorSettings(builder *flatbuffers.Builder, mutatorSettings flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(9, flatbuffers.UOffsetT(mutatorSettings), 0)
}
func MatchSettingsAddExistingMatchBehavior(builder *flatbuffers.Builder, existingMatchBehavior ExistingMatchBehavior) {
	builder.PrependByteSlot(10, byte(existingMatchBehavior), 0)
}
func MatchSettingsAddEnableRendering(builder *flatbuffers.Builder, enableRendering bool) {
	builder.PrependBoolSlot(11, enableRendering, false)
}
func MatchSettingsAddEnableStateSetting(builder *flatbuffers.Builder, enableStateSetting bool) {
	builder.PrependBoolSlot(12, enableStateSetting, false)
}
func MatchSettingsAddAutoSaveReplay(builder *flatbuffers.Builder, autoSaveReplay bool) {
	builder.PrependBoolSlot(13, autoSaveReplay, false)
}
func MatchSettingsAddFreeplay(builder *flatbuffers.Builder, freeplay bool) {
	builder.PrependBoolSlot(14, freeplay, false)
}
func MatchSettingsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
