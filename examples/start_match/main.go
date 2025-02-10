package main

import (
	RLBot "github.com/swz-git/go-interface"
	RLBotFlat "github.com/swz-git/go-interface/flat"
)

func main() {
	println("Connecting...")
	conn, err := RLBot.Connect("127.0.0.1:23234")
	if err != nil {
		panic("could not connect to rlbot core")
	}

	println("Sending MatchSettings packet...")
	conn.SendPacket(&RLBotFlat.MatchConfigurationT{
		PlayerConfigurations: []*RLBotFlat.PlayerConfigurationT{{
			Variety: &RLBotFlat.PlayerClassT{
				Type:  RLBotFlat.PlayerClassCustomBot,
				Value: &RLBotFlat.CustomBotT{},
			},
			Name:    "BOT1",
			Team:    0,
			Loadout: &RLBotFlat.PlayerLoadoutT{},
			SpawnId: 1,
		}, {
			Variety: &RLBotFlat.PlayerClassT{
				Type:  RLBotFlat.PlayerClassHuman,
				Value: &RLBotFlat.HumanT{},
			},
			Name:    "Human", // Cannot be "" for some reason?
			Team:    1,
			Loadout: &RLBotFlat.PlayerLoadoutT{},
			SpawnId: 2,
		}},
		GameMode:   RLBotFlat.GameModeSoccer,
		GameMapUpk: "UtopiaStadium_P",
		Mutators: &RLBotFlat.MutatorSettingsT{
			MatchLength: RLBotFlat.MatchLengthMutatorUnlimited,
		},
		ExistingMatchBehavior: RLBotFlat.ExistingMatchBehaviorRestart,
		EnableRendering:       true,
		// AutoStartBots:         false,
	})
}
