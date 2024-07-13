package gointerface

import (
	"encoding/binary"
	"errors"
	"net"

	flatbuffers "github.com/google/flatbuffers/go"
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

func (self RLBotConnection) SendPacket(packet_obj interface{}) error {
	var packetPayload []byte
	var packetType uint16

	self.builder.Reset()

	switch v := packet_obj.(type) {
	case nil:
		packetPayload = []byte{0}
		packetType = 0
	case GameTickPacket:
		self.builder.Finish(v.Pack(&self.builder))
		packetPayload = self.builder.FinishedBytes()
		packetType = 1
	case FieldInfo:
		self.builder.Finish(v.Pack(&self.builder))
		packetPayload = self.builder.FinishedBytes()
		packetType = 2
	case StartCommand:
		self.builder.Finish(v.Pack(&self.builder))
		packetPayload = self.builder.FinishedBytes()
		packetType = 3
	case MatchSettings:
		self.builder.Finish(v.Pack(&self.builder))
		packetPayload = self.builder.FinishedBytes()
		packetType = 4
	case PlayerInput:
		self.builder.Finish(v.Pack(&self.builder))
		packetPayload = self.builder.FinishedBytes()
		packetType = 5
	case DesiredGameState:
		self.builder.Finish(v.Pack(&self.builder))
		packetPayload = self.builder.FinishedBytes()
		packetType = 6
	case RenderGroup:
		self.builder.Finish(v.Pack(&self.builder))
		packetPayload = self.builder.FinishedBytes()
		packetType = 7
	case RemoveRenderGroup:
		self.builder.Finish(v.Pack(&self.builder))
		packetPayload = self.builder.FinishedBytes()
		packetType = 8
	case MatchComm:
		self.builder.Finish(v.Pack(&self.builder))
		packetPayload = self.builder.FinishedBytes()
		packetType = 9
	case BallPrediction:
		self.builder.Finish(v.Pack(&self.builder))
		packetPayload = self.builder.FinishedBytes()
		packetType = 10
	case ReadyMessage:
		self.builder.Finish(v.Pack(&self.builder))
		packetPayload = self.builder.FinishedBytes()
		packetType = 12
	case MessagePacket:
		self.builder.Finish(v.Pack(&self.builder))
		packetPayload = self.builder.FinishedBytes()
		packetType = 12
	case StopCommand:
		self.builder.Finish(v.Pack(&self.builder))
		packetPayload = self.builder.FinishedBytes()
		packetType = 13
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
