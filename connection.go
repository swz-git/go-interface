package gointerface

import (
	"encoding/binary"
	"errors"
	"io"
	"net"
	"os"
	"strconv"

	flatbuffers "github.com/google/flatbuffers/go"
	flat "github.com/swz-git/go-interface/flat"
)

type RLBotConnection struct {
	conn    net.Conn
	builder flatbuffers.Builder
}

func Connect(addr string) (RLBotConnection, error) {
	conn, err := net.Dial("tcp", addr)
	if tcpConn, ok := conn.(*net.TCPConn); ok {
		tcpConn.SetNoDelay(true)
	}
	return RLBotConnection{conn, *flatbuffers.NewBuilder(65536)}, err
}

func (conn RLBotConnection) Initialize(default_agent_id string, wants_ball_predictions bool, wants_comms bool) (*flat.MatchConfigurationT, *flat.FieldInfoT, *flat.ControllableTeamInfoT, error) {
	agent_id := os.Getenv("RLBOT_AGENT_ID")
	if agent_id == "" {
		agent_id = default_agent_id
	}

	err := conn.SendPacket(&flat.ConnectionSettingsT{
		AgentId:              agent_id,
		WantsBallPredictions: wants_ball_predictions,
		WantsComms:           wants_comms,
		CloseBetweenMatches:  true,
	})
	if err != nil {
		return nil, nil, nil, err
	}

	err = conn.SendPacket(&flat.InitCompleteT{})
	if err != nil {
		return nil, nil, nil, err
	}

	var match_config *flat.MatchConfigurationT
	var field_info *flat.FieldInfoT
	var controllables *flat.ControllableTeamInfoT

	for {
		packet, err := conn.RecvPacket()
		if err != nil {
			return nil, nil, nil, err
		}

		switch packet := packet.(type) {
		case *flat.MatchConfigurationT:
			match_config = packet
		case *flat.FieldInfoT:
			field_info = packet
		case *flat.ControllableTeamInfoT:
			controllables = packet
		}

		if match_config != nil && field_info != nil && controllables != nil {
			break
		}
	}

	return match_config, field_info, controllables, nil
}

type PacketAbilities interface {
	Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT
}

func (conn RLBotConnection) SendPacket(packet_obj PacketAbilities) error {
	var packetPayload []byte
	var packetType uint16

	conn.builder.Reset()

	switch v := packet_obj.(type) {
	case nil:
		packetPayload = []byte{0}
		packetType = 0
	case *flat.StartCommandT:
		conn.builder.Finish(v.Pack(&conn.builder))
		packetPayload = conn.builder.FinishedBytes()
		packetType = 3
	case *flat.MatchConfigurationT:
		conn.builder.Finish(v.Pack(&conn.builder))
		packetPayload = conn.builder.FinishedBytes()
		packetType = 4
	case *flat.PlayerInputT:
		conn.builder.Finish(v.Pack(&conn.builder))
		packetPayload = conn.builder.FinishedBytes()
		packetType = 5
	case *flat.DesiredGameStateT:
		conn.builder.Finish(v.Pack(&conn.builder))
		packetPayload = conn.builder.FinishedBytes()
		packetType = 6
	case *flat.RenderGroupT:
		conn.builder.Finish(v.Pack(&conn.builder))
		packetPayload = conn.builder.FinishedBytes()
		packetType = 7
	case *flat.RemoveRenderGroupT:
		conn.builder.Finish(v.Pack(&conn.builder))
		packetPayload = conn.builder.FinishedBytes()
		packetType = 8
	case *flat.MatchCommT:
		conn.builder.Finish(v.Pack(&conn.builder))
		packetPayload = conn.builder.FinishedBytes()
		packetType = 9
	case *flat.ConnectionSettingsT:
		conn.builder.Finish(v.Pack(&conn.builder))
		packetPayload = conn.builder.FinishedBytes()
		packetType = 11
	case *flat.StopCommandT:
		conn.builder.Finish(v.Pack(&conn.builder))
		packetPayload = conn.builder.FinishedBytes()
		packetType = 12
	case *flat.SetLoadoutT:
		conn.builder.Finish(v.Pack(&conn.builder))
		packetPayload = conn.builder.FinishedBytes()
		packetType = 13
	case *flat.InitCompleteT:
		packetPayload = []byte{0}
		packetType = 14
	default:
		return errors.New("unsupported packet type")
	}

	finalBuf := []byte{}

	tempBuf := make([]byte, 2)
	binary.BigEndian.PutUint16(tempBuf, packetType)
	finalBuf = append(finalBuf, tempBuf...)

	tempBuf = make([]byte, 2)
	binary.BigEndian.PutUint16(tempBuf, uint16(len(packetPayload)))
	finalBuf = append(finalBuf, tempBuf...)

	finalBuf = append(finalBuf, packetPayload...)

	_, err := conn.conn.Write(finalBuf)
	if err != nil {
		return err
	}

	return nil
}

func (conn RLBotConnection) RecvPacket() (PacketAbilities, error) {
	buffer := make([]byte, 2)

	// Read packetType
	_, err := io.ReadFull(conn.conn, buffer)
	if err != nil {
		return nil, err
	}
	packetType := binary.BigEndian.Uint16(buffer)

	// Read packetLen
	_, err = io.ReadFull(conn.conn, buffer)
	if err != nil {
		return nil, err
	}
	packetLen := binary.BigEndian.Uint16(buffer)

	// Read packetPayload
	buffer = make([]byte, packetLen)
	_, err = io.ReadFull(conn.conn, buffer)
	if err != nil {
		return nil, err
	}

	switch packetType {
	case 0:
		return nil, nil
	case 1: //flat.GamePacketT:
		return flat.GetRootAsGamePacket(buffer, 0).UnPack(), nil
	case 2: //flat.FieldInfoT:
		return flat.GetRootAsFieldInfo(buffer, 0).UnPack(), nil
	case 4: //flat.MatchConfigurationT:
		return flat.GetRootAsMatchConfiguration(buffer, 0).UnPack(), nil
	case 5: //flat.PlayerInputT:
		return flat.GetRootAsPlayerInput(buffer, 0).UnPack(), nil
	case 9: //flat.MatchCommT:
		return flat.GetRootAsMatchComm(buffer, 0).UnPack(), nil
	case 10: //flat.BallPredictionT:
		return flat.GetRootAsBallPrediction(buffer, 0).UnPack(), nil
	case 15: //flat.ControllableTeamInfoT:
		return flat.GetRootAsControllableTeamInfo(buffer, 0).UnPack(), nil
	default:
		return nil, errors.New("unknown packet type: " + strconv.Itoa(int(packetType)))
	}
}
