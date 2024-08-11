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

	spawnId, err := strconv.Atoi(os.Getenv("RLBOT_SPAWN_IDS"))
	if err != nil {
		panic("ERR: RLBOT_SPAWN_IDS wasn't an integer")
	}

	err = conn.SendPacket(&RLBotFlat.ConnectionSettingsT{
		WantsBallPredictions: true,
		WantsComms:           true,
		CloseAfterMatch:      true,
	})
	if err != nil {
		panic(err)
	}

	err = conn.SendPacket(&RLBotFlat.InitCompleteT{
		SpawnId: int32(spawnId),
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

		var botIndex int
		for i, player := range gameTickPacket.Players {
			if player.SpawnId == int32(spawnId) {
				botIndex = i
			}
		}
		if botIndex == 0 {
			// If we aren't in the game, don't do anything
			continue
		}

		if len(gameTickPacket.Balls) < 1 {
			continue // no ball to chase :(
		}

		target := gameTickPacket.Balls[0].Physics
		car := gameTickPacket.Players[spawnId].Physics

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
			PlayerIndex:     uint32(spawnId),
			ControllerState: &controller,
		})
	}
}
