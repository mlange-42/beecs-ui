package game

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mlange-42/arche-model/model"
)

const TPS = 30

func Run() {
	game := NewGame(nil)
	game.Initialize()

	initGame(&game)

	if err := game.Run(); err != nil {
		log.Fatal(err)
	}
}

func initGame(g *Game) error {
	ebiten.SetVsyncEnabled(true)

	g.Model = model.New()

	g.Model.Initialize()

	return nil
}
