// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flat

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// Definition of a match.
/// Can be sent to RLBot to request the start of a match.
type MatchConfigurationT struct {
	Launcher Launcher `json:"launcher"`
	LauncherArg string `json:"launcher_arg"`
	AutoStartBots bool `json:"auto_start_bots"`
	GameMapUpk string `json:"game_map_upk"`
	PlayerConfigurations []*PlayerConfigurationT `json:"player_configurations"`
	ScriptConfigurations []*ScriptConfigurationT `json:"script_configurations"`
	GameMode GameMode `json:"game_mode"`
	SkipReplays bool `json:"skip_replays"`
	InstantStart bool `json:"instant_start"`
	Mutators *MutatorSettingsT `json:"mutators"`
	ExistingMatchBehavior ExistingMatchBehavior `json:"existing_match_behavior"`
	EnableRendering bool `json:"enable_rendering"`
	EnableStateSetting bool `json:"enable_state_setting"`
	AutoSaveReplay bool `json:"auto_save_replay"`
	Freeplay bool `json:"freeplay"`
}

func (t *MatchConfigurationT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	launcherArgOffset := flatbuffers.UOffsetT(0)
	if t.LauncherArg != "" {
		launcherArgOffset = builder.CreateString(t.LauncherArg)
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
		MatchConfigurationStartPlayerConfigurationsVector(builder, playerConfigurationsLength)
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
		MatchConfigurationStartScriptConfigurationsVector(builder, scriptConfigurationsLength)
		for j := scriptConfigurationsLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(scriptConfigurationsOffsets[j])
		}
		scriptConfigurationsOffset = builder.EndVector(scriptConfigurationsLength)
	}
	mutatorsOffset := t.Mutators.Pack(builder)
	MatchConfigurationStart(builder)
	MatchConfigurationAddLauncher(builder, t.Launcher)
	MatchConfigurationAddLauncherArg(builder, launcherArgOffset)
	MatchConfigurationAddAutoStartBots(builder, t.AutoStartBots)
	MatchConfigurationAddGameMapUpk(builder, gameMapUpkOffset)
	MatchConfigurationAddPlayerConfigurations(builder, playerConfigurationsOffset)
	MatchConfigurationAddScriptConfigurations(builder, scriptConfigurationsOffset)
	MatchConfigurationAddGameMode(builder, t.GameMode)
	MatchConfigurationAddSkipReplays(builder, t.SkipReplays)
	MatchConfigurationAddInstantStart(builder, t.InstantStart)
	MatchConfigurationAddMutators(builder, mutatorsOffset)
	MatchConfigurationAddExistingMatchBehavior(builder, t.ExistingMatchBehavior)
	MatchConfigurationAddEnableRendering(builder, t.EnableRendering)
	MatchConfigurationAddEnableStateSetting(builder, t.EnableStateSetting)
	MatchConfigurationAddAutoSaveReplay(builder, t.AutoSaveReplay)
	MatchConfigurationAddFreeplay(builder, t.Freeplay)
	return MatchConfigurationEnd(builder)
}

func (rcv *MatchConfiguration) UnPackTo(t *MatchConfigurationT) {
	t.Launcher = rcv.Launcher()
	t.LauncherArg = string(rcv.LauncherArg())
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
	t.Mutators = rcv.Mutators(nil).UnPack()
	t.ExistingMatchBehavior = rcv.ExistingMatchBehavior()
	t.EnableRendering = rcv.EnableRendering()
	t.EnableStateSetting = rcv.EnableStateSetting()
	t.AutoSaveReplay = rcv.AutoSaveReplay()
	t.Freeplay = rcv.Freeplay()
}

func (rcv *MatchConfiguration) UnPack() *MatchConfigurationT {
	if rcv == nil {
		return nil
	}
	t := &MatchConfigurationT{}
	rcv.UnPackTo(t)
	return t
}

type MatchConfiguration struct {
	_tab flatbuffers.Table
}

func GetRootAsMatchConfiguration(buf []byte, offset flatbuffers.UOffsetT) *MatchConfiguration {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &MatchConfiguration{}
	x.Init(buf, n+offset)
	return x
}

func FinishMatchConfigurationBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsMatchConfiguration(buf []byte, offset flatbuffers.UOffsetT) *MatchConfiguration {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &MatchConfiguration{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedMatchConfigurationBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *MatchConfiguration) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *MatchConfiguration) Table() flatbuffers.Table {
	return rcv._tab
}

/// How to launch Rocket League.
/// If left unset, RLBot will not launch the game.
/// To use Legendary, use Custom and set launcher_arg="legendary".
func (rcv *MatchConfiguration) Launcher() Launcher {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return Launcher(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

/// How to launch Rocket League.
/// If left unset, RLBot will not launch the game.
/// To use Legendary, use Custom and set launcher_arg="legendary".
func (rcv *MatchConfiguration) MutateLauncher(n Launcher) bool {
	return rcv._tab.MutateByteSlot(4, byte(n))
}

/// Additional configuration for the launching method.
/// See launcher.
func (rcv *MatchConfiguration) LauncherArg() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Additional configuration for the launching method.
/// See launcher.
/// If true, RLBot will start the bots with a non-empty run command in their player configuration.
func (rcv *MatchConfiguration) AutoStartBots() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

/// If true, RLBot will start the bots with a non-empty run command in their player configuration.
func (rcv *MatchConfiguration) MutateAutoStartBots(n bool) bool {
	return rcv._tab.MutateBoolSlot(8, n)
}

/// The name of a upk file, like UtopiaStadium_P, which should be loaded.
/// On Steam version of Rocket League this can be used to load custom map files,
/// but on Epic version it only works on the Psyonix maps.
/// Available maps can be found here: https://github.com/VirxEC/python-interface/blob/master/rlbot/utils/maps.py
func (rcv *MatchConfiguration) GameMapUpk() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// The name of a upk file, like UtopiaStadium_P, which should be loaded.
/// On Steam version of Rocket League this can be used to load custom map files,
/// but on Epic version it only works on the Psyonix maps.
/// Available maps can be found here: https://github.com/VirxEC/python-interface/blob/master/rlbot/utils/maps.py
/// The players in the match.
func (rcv *MatchConfiguration) PlayerConfigurations(obj *PlayerConfiguration, j int) bool {
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

func (rcv *MatchConfiguration) PlayerConfigurationsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// The players in the match.
/// The custom scripts used in the match.
func (rcv *MatchConfiguration) ScriptConfigurations(obj *ScriptConfiguration, j int) bool {
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

func (rcv *MatchConfiguration) ScriptConfigurationsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// The custom scripts used in the match.
/// The game mode.
/// This affects a few of the game rules although many game modes can also be recreated solely from mutators.
/// See what mutators and game mode combinations make up the official modes at https://github.com/VirxEC/python-interface/tree/master/tests/gamemodes
func (rcv *MatchConfiguration) GameMode() GameMode {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return GameMode(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

/// The game mode.
/// This affects a few of the game rules although many game modes can also be recreated solely from mutators.
/// See what mutators and game mode combinations make up the official modes at https://github.com/VirxEC/python-interface/tree/master/tests/gamemodes
func (rcv *MatchConfiguration) MutateGameMode(n GameMode) bool {
	return rcv._tab.MutateByteSlot(16, byte(n))
}

/// Whether to skip goal replays.
func (rcv *MatchConfiguration) SkipReplays() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

/// Whether to skip goal replays.
func (rcv *MatchConfiguration) MutateSkipReplays(n bool) bool {
	return rcv._tab.MutateBoolSlot(18, n)
}

/// Whether to start without a kickoff countdown.
func (rcv *MatchConfiguration) InstantStart() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

/// Whether to start without a kickoff countdown.
func (rcv *MatchConfiguration) MutateInstantStart(n bool) bool {
	return rcv._tab.MutateBoolSlot(20, n)
}

/// Mutator settings.
func (rcv *MatchConfiguration) Mutators(obj *MutatorSettings) *MutatorSettings {
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

/// Mutator settings.
/// How to handle any ongoing match.
func (rcv *MatchConfiguration) ExistingMatchBehavior() ExistingMatchBehavior {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		return ExistingMatchBehavior(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

/// How to handle any ongoing match.
func (rcv *MatchConfiguration) MutateExistingMatchBehavior(n ExistingMatchBehavior) bool {
	return rcv._tab.MutateByteSlot(24, byte(n))
}

/// Whether debug rendering is displayed.
func (rcv *MatchConfiguration) EnableRendering() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(26))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

/// Whether debug rendering is displayed.
func (rcv *MatchConfiguration) MutateEnableRendering(n bool) bool {
	return rcv._tab.MutateBoolSlot(26, n)
}

/// Whether clients are allowed to manipulate the game state, e.g. teleporting cars and ball.
func (rcv *MatchConfiguration) EnableStateSetting() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(28))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

/// Whether clients are allowed to manipulate the game state, e.g. teleporting cars and ball.
func (rcv *MatchConfiguration) MutateEnableStateSetting(n bool) bool {
	return rcv._tab.MutateBoolSlot(28, n)
}

/// Whether the match replay should be saved.
func (rcv *MatchConfiguration) AutoSaveReplay() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(30))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

/// Whether the match replay should be saved.
func (rcv *MatchConfiguration) MutateAutoSaveReplay(n bool) bool {
	return rcv._tab.MutateBoolSlot(30, n)
}

/// If set to true, a free play match is launched instead of an exhibition match.
/// This allows the players to use training keybinds, Bakkesmod plugins, and other features that are only allowed in free play.
func (rcv *MatchConfiguration) Freeplay() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(32))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

/// If set to true, a free play match is launched instead of an exhibition match.
/// This allows the players to use training keybinds, Bakkesmod plugins, and other features that are only allowed in free play.
func (rcv *MatchConfiguration) MutateFreeplay(n bool) bool {
	return rcv._tab.MutateBoolSlot(32, n)
}

func MatchConfigurationStart(builder *flatbuffers.Builder) {
	builder.StartObject(15)
}
func MatchConfigurationAddLauncher(builder *flatbuffers.Builder, launcher Launcher) {
	builder.PrependByteSlot(0, byte(launcher), 0)
}
func MatchConfigurationAddLauncherArg(builder *flatbuffers.Builder, launcherArg flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(launcherArg), 0)
}
func MatchConfigurationAddAutoStartBots(builder *flatbuffers.Builder, autoStartBots bool) {
	builder.PrependBoolSlot(2, autoStartBots, false)
}
func MatchConfigurationAddGameMapUpk(builder *flatbuffers.Builder, gameMapUpk flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(gameMapUpk), 0)
}
func MatchConfigurationAddPlayerConfigurations(builder *flatbuffers.Builder, playerConfigurations flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(playerConfigurations), 0)
}
func MatchConfigurationStartPlayerConfigurationsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func MatchConfigurationAddScriptConfigurations(builder *flatbuffers.Builder, scriptConfigurations flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(scriptConfigurations), 0)
}
func MatchConfigurationStartScriptConfigurationsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func MatchConfigurationAddGameMode(builder *flatbuffers.Builder, gameMode GameMode) {
	builder.PrependByteSlot(6, byte(gameMode), 0)
}
func MatchConfigurationAddSkipReplays(builder *flatbuffers.Builder, skipReplays bool) {
	builder.PrependBoolSlot(7, skipReplays, false)
}
func MatchConfigurationAddInstantStart(builder *flatbuffers.Builder, instantStart bool) {
	builder.PrependBoolSlot(8, instantStart, false)
}
func MatchConfigurationAddMutators(builder *flatbuffers.Builder, mutators flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(9, flatbuffers.UOffsetT(mutators), 0)
}
func MatchConfigurationAddExistingMatchBehavior(builder *flatbuffers.Builder, existingMatchBehavior ExistingMatchBehavior) {
	builder.PrependByteSlot(10, byte(existingMatchBehavior), 0)
}
func MatchConfigurationAddEnableRendering(builder *flatbuffers.Builder, enableRendering bool) {
	builder.PrependBoolSlot(11, enableRendering, false)
}
func MatchConfigurationAddEnableStateSetting(builder *flatbuffers.Builder, enableStateSetting bool) {
	builder.PrependBoolSlot(12, enableStateSetting, false)
}
func MatchConfigurationAddAutoSaveReplay(builder *flatbuffers.Builder, autoSaveReplay bool) {
	builder.PrependBoolSlot(13, autoSaveReplay, false)
}
func MatchConfigurationAddFreeplay(builder *flatbuffers.Builder, freeplay bool) {
	builder.PrependBoolSlot(14, freeplay, false)
}
func MatchConfigurationEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
