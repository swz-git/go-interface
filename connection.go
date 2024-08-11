package gointerface

import (
	"encoding/binary"
	"errors"
	"io"
	"net"
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

type PacketAblilities interface {
	Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT
}

// TODO: Disallow sending and recving packets that shouldn't be

func (self RLBotConnection) SendPacket(packet_obj PacketAblilities) error {
	var packetPayload []byte
	var packetType uint16

	self.builder.Reset()

	switch v := packet_obj.(type) {
	case nil:
		packetPayload = []byte{0}
		packetType = 0
	case *flat.GameTickPacketT:
		self.builder.Finish(v.Pack(&self.builder))
		packetPayload = self.builder.FinishedBytes()
		packetType = 1
	case *flat.FieldInfoT:
		self.builder.Finish(v.Pack(&self.builder))
		packetPayload = self.builder.FinishedBytes()
		packetType = 2
	case *flat.StartCommandT:
		self.builder.Finish(v.Pack(&self.builder))
		packetPayload = self.builder.FinishedBytes()
		packetType = 3
	case *flat.MatchSettingsT:
		self.builder.Finish(v.Pack(&self.builder))
		packetPayload = self.builder.FinishedBytes()
		packetType = 4
	case *flat.PlayerInputT:
		self.builder.Finish(v.Pack(&self.builder))
		packetPayload = self.builder.FinishedBytes()
		packetType = 5
	case *flat.DesiredGameStateT:
		self.builder.Finish(v.Pack(&self.builder))
		packetPayload = self.builder.FinishedBytes()
		packetType = 6
	case *flat.RenderGroupT:
		self.builder.Finish(v.Pack(&self.builder))
		packetPayload = self.builder.FinishedBytes()
		packetType = 7
	case *flat.RemoveRenderGroupT:
		self.builder.Finish(v.Pack(&self.builder))
		packetPayload = self.builder.FinishedBytes()
		packetType = 8
	case *flat.MatchCommT:
		self.builder.Finish(v.Pack(&self.builder))
		packetPayload = self.builder.FinishedBytes()
		packetType = 9
	case *flat.BallPredictionT:
		self.builder.Finish(v.Pack(&self.builder))
		packetPayload = self.builder.FinishedBytes()
		packetType = 10
	case *flat.ConnectionSettingsT:
		self.builder.Finish(v.Pack(&self.builder))
		packetPayload = self.builder.FinishedBytes()
		packetType = 11
	case *flat.StopCommandT:
		self.builder.Finish(v.Pack(&self.builder))
		packetPayload = self.builder.FinishedBytes()
		packetType = 12
	case *flat.SetLoadoutT:
		self.builder.Finish(v.Pack(&self.builder))
		packetPayload = self.builder.FinishedBytes()
		packetType = 13
	case *flat.InitCompleteT:
		self.builder.Finish(v.Pack(&self.builder))
		packetPayload = self.builder.FinishedBytes()
		packetType = 14
	default:
		return errors.New("Unsupported packet type")
	}

	finalBuf := []byte{}

	tempBuf := make([]byte, 2)
	binary.BigEndian.PutUint16(tempBuf, packetType)
	finalBuf = append(finalBuf, tempBuf...)

	tempBuf = make([]byte, 2)
	binary.BigEndian.PutUint16(tempBuf, uint16(len(packetPayload)))
	finalBuf = append(finalBuf, tempBuf...)

	finalBuf = append(finalBuf, packetPayload...)

	_, err := self.conn.Write(finalBuf)
	if err != nil {
		return err
	}

	return nil
}

func (self RLBotConnection) RecvPacket() (PacketAblilities, error) {
	buffer := make([]byte, 2)

	// Read packetType
	_, err := io.ReadFull(self.conn, buffer)
	if err != nil {
		println("1")
		return nil, err
	}
	packetType := binary.BigEndian.Uint16(buffer)

	// Read packetLen
	_, err = io.ReadFull(self.conn, buffer)
	if err != nil {
		println("2")
		return nil, err
	}
	packetLen := binary.BigEndian.Uint16(buffer)

	// Read packetPayload
	buffer = make([]byte, packetLen)
	_, err = io.ReadFull(self.conn, buffer)
	if err != nil {
		println("3")
		return nil, err
	}

	switch packetType {
	case 0:
		return nil, nil
	case 1: //flat.GameTickPacketT:
		return flat.GetRootAsGameTickPacket(buffer, 0).UnPack(), nil
	case 2: //flat.FieldInfoT:
		return flat.GetRootAsFieldInfo(buffer, 0).UnPack(), nil
	case 3: //flat.StartCommandT:
		return flat.GetRootAsStartCommand(buffer, 0).UnPack(), nil
	case 4: //flat.MatchSettingsT:
		return flat.GetRootAsMatchSettings(buffer, 0).UnPack(), nil
	case 5: //flat.PlayerInputT:
		return flat.GetRootAsPlayerInput(buffer, 0).UnPack(), nil
	case 6: //flat.DesiredGameStateT:
		return flat.GetRootAsDesiredGameState(buffer, 0).UnPack(), nil
	case 7: //flat.RenderGroupT:
		return flat.GetRootAsRenderGroup(buffer, 0).UnPack(), nil
	case 8: //flat.RemoveRenderGroupT:
		return flat.GetRootAsRemoveRenderGroup(buffer, 0).UnPack(), nil
	case 9: //flat.MatchCommT:
		return flat.GetRootAsMatchComm(buffer, 0).UnPack(), nil
	case 10: //flat.BallPredictionT:
		return flat.GetRootAsBallPrediction(buffer, 0).UnPack(), nil
	case 11: //flat.ConnectionSettingsT:
		return flat.GetRootAsConnectionSettings(buffer, 0).UnPack(), nil
	case 12: //flat.StopCommandT:
		return flat.GetRootAsStopCommand(buffer, 0).UnPack(), nil
	case 13: //flat.SetLoadoutT:
		return flat.GetRootAsSetLoadout(buffer, 0).UnPack(), nil
	case 14: //flat.InitCompleteT:
		return flat.GetRootAsInitComplete(buffer, 0).UnPack(), nil
	default:
		return nil, errors.New("Unknown packet type: " + strconv.Itoa(int(packetType)))
	}
}
