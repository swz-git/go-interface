package main

import (
	"os"

	RLBot "github.com/swz-git/go-interface"
	RLBotFlat "github.com/swz-git/go-interface/flat"
)

func main() {
	println("Connecting...")
	conn, err := RLBot.Connect("127.0.0.1:23234")
	if err != nil {
		panic("could not connect to rlbot core")
	}

	dir, err := os.Getwd()
	if err != nil {
		panic("could not get current working directory")
	}

	println("Sending MatchConfiguration packet...")
	conn.SendPacket(&RLBotFlat.MatchConfigurationT{
		PlayerConfigurations: []*RLBotFlat.PlayerConfigurationT{{
			Variety: &RLBotFlat.PlayerClassT{
				Type:  RLBotFlat.PlayerClassCustomBot,
				Value: &RLBotFlat.CustomBotT{},
			},
			RootDir: dir + "/examples/atba",
			RunCommand: "go run main.go",
			AgentId: "rlbot/go-example-bot",
			Name:    "Go Example",
			Team:    0,
			Loadout: &RLBotFlat.PlayerLoadoutT{},
		}, {
			Variety: &RLBotFlat.PlayerClassT{
				Type:  RLBotFlat.PlayerClassHuman,
				Value: &RLBotFlat.HumanT{},
			},
			Name:    "",
			Team:    1,
			Loadout: &RLBotFlat.PlayerLoadoutT{},
		}},
		GameMode:   RLBotFlat.GameModeSoccer,
		GameMapUpk: "UtopiaStadium_P",
		Mutators: &RLBotFlat.MutatorSettingsT{
			MatchLength: RLBotFlat.MatchLengthMutatorUnlimited,
		},
		ExistingMatchBehavior: RLBotFlat.ExistingMatchBehaviorRestart,
		EnableRendering:       true,
		AutoStartBots:         true,
	})
}
