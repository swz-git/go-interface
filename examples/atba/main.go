package main

import (
	"math"
	"os"
	"strconv"

	RLBot "github.com/swz-git/go-interface"
	RLBotFlat "github.com/swz-git/go-interface/flat"
)

func main() {
	println("Connecting...")

	addr := os.Getenv("RLBOT_CORE_ADDR")
	if addr == "" {
		addr = "127.0.0.1:23234"
	}

	conn, err := RLBot.Connect(addr)
	if err != nil {
		panic("could not connect to rlbot core")
	}

	println("Running!")

	carIndex, err := strconv.Atoi(os.Getenv("RLBOT_INDEX"))
	if err != nil {
		println("WARN: No RLBOT_INDEX found, assuming 0")
		carIndex = 0
	}

	err = conn.SendPacket(&RLBotFlat.ReadyMessageT{
		WantsBallPredictions: true,
		WantsComms:           true,
		CloseAfterMatch:      true,
	})
	if err != nil {
		panic(err)
	}

	for {
		packet, err := conn.RecvPacket()
		if err != nil {
			panic(err)
		}
		gameTickPacket, ok := packet.(*RLBotFlat.GameTickPacketT)
		if !ok { // if not gametickpacket
			continue
		}
		if len(gameTickPacket.Balls) < 1 {
			continue // no ball to chase :(
		}

		target := gameTickPacket.Balls[0].Physics
		car := gameTickPacket.Players[carIndex].Physics

		botToTargetAngle := math.Atan2(float64(target.Location.Y-car.Location.Y), float64(target.Location.X-car.Location.X))
		botFrontToTargetAngle := botToTargetAngle - float64(car.Rotation.Yaw)

		if botFrontToTargetAngle > 3.14 {
			botFrontToTargetAngle -= 2. * 3.14
		}
		if botFrontToTargetAngle < -3.14 {
			botFrontToTargetAngle += 2. * 3.14
		}

		controller := RLBotFlat.ControllerStateT{}

		if botFrontToTargetAngle > 0. {
			controller.Steer = 1
		} else {
			controller.Steer = -1
		}

		controller.Throttle = 1

		conn.SendPacket(&RLBotFlat.PlayerInputT{
			PlayerIndex:     uint32(carIndex),
			ControllerState: &controller,
		})
	}
}
