package game

import (
	"embed"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	arche "github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/beecs-ui/game/res"
	"github.com/mlange-42/beecs-ui/game/sys"
	"github.com/mlange-42/beecs/model"
	"github.com/mlange-42/beecs/params"
)

const TPS = 30

var GameData embed.FS

func Run(data embed.FS) {
	GameData = data

	game := NewGame(nil)
	game.Initialize()

	initGame(&game)

	if err := game.Run(); err != nil {
		log.Fatal(err)
	}
}

func run(g *Game) {
	if err := initGame(g); err != nil {
		panic(err)
	}
}

func initGame(g *Game) error {
	ebiten.SetVsyncEnabled(true)

	g.Model = arche.New()

	p := params.Default()
	model.Default(&p, g.Model)

	ecs.AddResource(&g.Model.World, &res.GameSpeed{
		MinSpeed: -2,
		MaxSpeed: 3,
	})

	ecs.AddResource(&g.Model.World, &res.GameTick{})

	ecs.AddResource(&g.Model.World, &res.UpdateInterval{
		Interval: TPS,
	})

	ecs.AddResource(&g.Model.World, &g.Screen)
	ecs.AddResource(&g.Model.World, &g.Mouse)

	sprites := res.NewSprites(GameData, "data/images")
	ecs.AddResource(&g.Model.World, &sprites)

	fonts := res.NewFonts(GameData, "data/fonts")
	ecs.AddResource(&g.Model.World, &fonts)

	g.Model.AddSystem(&sys.InitUI{
		ResetFn: func() {
			run(g)
		},
	})
	g.Model.AddSystem(&sys.UpdateUI{})

	g.Model.AddSystem(&sys.GameControls{
		PauseKey:      ebiten.KeySpace,
		StepKey:       ebiten.KeyArrowRight,
		SlowerKey:     '[',
		FasterKey:     ']',
		FullscreenKey: ebiten.KeyF11,
	})
	g.Model.AddUISystem(&sys.RenderUI{})
	g.Model.AddSystem(&sys.Tick{})

	g.Model.Initialize()

	return nil
}
