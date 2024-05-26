package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mlange-42/beecs-ui/game"
)

func main() {
	var parFile string

	switch len(os.Args) {
	case 1:
		fmt.Println("Optionally, provide a parameter file to load as argument")
	case 2:
		parFile = os.Args[1]
	default:
		log.Fatal(fmt.Errorf("expects zero or one arguments, got %d", len(os.Args)-1))
	}

	//stop := profile.Start(profile.CPUProfile, profile.ProfilePath("."))
	game.Run(GameData, parFile)
	//stop.Stop()
}
