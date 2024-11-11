// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flat

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ScriptConfigurationT struct {
	Name string `json:"name"`
	Location string `json:"location"`
	RunCommand string `json:"run_command"`
	SpawnId int32 `json:"spawn_id"`
	AgentId string `json:"agent_id"`
}

func (t *ScriptConfigurationT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	nameOffset := flatbuffers.UOffsetT(0)
	if t.Name != "" {
		nameOffset = builder.CreateString(t.Name)
	}
	locationOffset := flatbuffers.UOffsetT(0)
	if t.Location != "" {
		locationOffset = builder.CreateString(t.Location)
	}
	runCommandOffset := flatbuffers.UOffsetT(0)
	if t.RunCommand != "" {
		runCommandOffset = builder.CreateString(t.RunCommand)
	}
	agentIdOffset := flatbuffers.UOffsetT(0)
	if t.AgentId != "" {
		agentIdOffset = builder.CreateString(t.AgentId)
	}
	ScriptConfigurationStart(builder)
	ScriptConfigurationAddName(builder, nameOffset)
	ScriptConfigurationAddLocation(builder, locationOffset)
	ScriptConfigurationAddRunCommand(builder, runCommandOffset)
	ScriptConfigurationAddSpawnId(builder, t.SpawnId)
	ScriptConfigurationAddAgentId(builder, agentIdOffset)
	return ScriptConfigurationEnd(builder)
}

func (rcv *ScriptConfiguration) UnPackTo(t *ScriptConfigurationT) {
	t.Name = string(rcv.Name())
	t.Location = string(rcv.Location())
	t.RunCommand = string(rcv.RunCommand())
	t.SpawnId = rcv.SpawnId()
	t.AgentId = string(rcv.AgentId())
}

func (rcv *ScriptConfiguration) UnPack() *ScriptConfigurationT {
	if rcv == nil {
		return nil
	}
	t := &ScriptConfigurationT{}
	rcv.UnPackTo(t)
	return t
}

type ScriptConfiguration struct {
	_tab flatbuffers.Table
}

func GetRootAsScriptConfiguration(buf []byte, offset flatbuffers.UOffsetT) *ScriptConfiguration {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ScriptConfiguration{}
	x.Init(buf, n+offset)
	return x
}

func FinishScriptConfigurationBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsScriptConfiguration(buf []byte, offset flatbuffers.UOffsetT) *ScriptConfiguration {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &ScriptConfiguration{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedScriptConfigurationBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *ScriptConfiguration) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ScriptConfiguration) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ScriptConfiguration) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *ScriptConfiguration) Location() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *ScriptConfiguration) RunCommand() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *ScriptConfiguration) SpawnId() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *ScriptConfiguration) MutateSpawnId(n int32) bool {
	return rcv._tab.MutateInt32Slot(10, n)
}

func (rcv *ScriptConfiguration) AgentId() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func ScriptConfigurationStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func ScriptConfigurationAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(name), 0)
}
func ScriptConfigurationAddLocation(builder *flatbuffers.Builder, location flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(location), 0)
}
func ScriptConfigurationAddRunCommand(builder *flatbuffers.Builder, runCommand flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(runCommand), 0)
}
func ScriptConfigurationAddSpawnId(builder *flatbuffers.Builder, spawnId int32) {
	builder.PrependInt32Slot(3, spawnId, 0)
}
func ScriptConfigurationAddAgentId(builder *flatbuffers.Builder, agentId flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(agentId), 0)
}
func ScriptConfigurationEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}