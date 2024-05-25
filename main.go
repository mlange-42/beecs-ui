package main

import (
	"github.com/mlange-42/beecs-ui/game"
)

func main() {
	//stop := profile.Start(profile.CPUProfile, profile.ProfilePath("."))
	game.Run(GameData, "parameters.json")
	//stop.Stop()
}
