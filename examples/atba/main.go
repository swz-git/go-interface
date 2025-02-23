package main

import (
	"fmt"
	"math"
	"os"

	RLBot "github.com/swz-git/go-interface"
	RLBotFlat "github.com/swz-git/go-interface/flat"
)

const DEFAULT_AGENT_ID = "rlbot/go-example-bot"

func main() {
	println("Connecting...")

	rlbot_ip := os.Getenv("RLBOT_SERVER_IP")
	if rlbot_ip == "" {
		rlbot_ip = "127.0.0.1"
	}

	rlbot_port := os.Getenv("RLBOT_SERVER_PORT")
	if rlbot_port == "" {
		rlbot_port = "23234"
	}

	addr := fmt.Sprintf("%s:%s", rlbot_ip, rlbot_port)
	conn, err := RLBot.Connect(addr)
	if err != nil {
		panic(fmt.Sprintf("could not connect to rlbot core at %s: %s", addr, err))
	}

	println("Connected, initializing...")

	match_config, _, controllables, err := conn.Initialize(DEFAULT_AGENT_ID, false, false)
	if err != nil {
		panic(err)
	}

	team := controllables.Team
	index := controllables.Controllables[0].Index
	spawnId := controllables.Controllables[0].SpawnId

	var name string
	for _, player := range match_config.PlayerConfigurations {
		if player.SpawnId == spawnId {
			name = player.Name
			break
		}
	}

	fmt.Printf("Initialized, running as \"%s\" on team %d\n", name, team)

	for {
		packet, err := conn.RecvPacket()
		if err != nil {
			panic(err)
		}
		gameTickPacket, ok := packet.(*RLBotFlat.GamePacketT)
		if !ok { // if not gametickpacket
			continue
		}

		if len(gameTickPacket.Players) <= int(index) {
			continue // we haven't been spawned in yet
		}

		if len(gameTickPacket.Balls) < 1 {
			continue // no ball to chase :(
		}

		target := gameTickPacket.Balls[0].Physics
		car := gameTickPacket.Players[index].Physics

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
			PlayerIndex:     index,
			ControllerState: &controller,
		})
	}
}
