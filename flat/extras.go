package flat

import flatbuffers "github.com/google/flatbuffers/go"

// There is no InitComplete in the schema but we still need to be able to send it
type InitCompleteT struct{}

// This function is purely here for InitComplete to implement PacketAbilities
func (InitCompleteT) Pack(*flatbuffers.Builder) flatbuffers.UOffsetT { return 0 }
