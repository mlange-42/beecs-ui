package game

import (
	"log"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2"
	arche "github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/beecs-ui/data"
	"github.com/mlange-42/beecs-ui/internal/res"
	"github.com/mlange-42/beecs-ui/internal/sys"
	"github.com/mlange-42/beecs/model"
	"github.com/mlange-42/beecs/params"
)

const TPS = 30

func Run(paramsFile string) {
	game := NewGame(nil)
	game.Initialize()

	if err := initGame(&game, paramsFile, map[string]any{}); err != nil {
		log.Fatal(err)
	}

	if err := game.Run(); err != nil {
		log.Fatal(err)
	}
}

func run(g *Game, paramsFile string, overwriteParams map[string]any) {
	if err := initGame(g, paramsFile, overwriteParams); err != nil {
		panic(err)
	}
}

func initGame(g *Game, paramsFile string, overwriteParams map[string]any) error {
	ebiten.SetVsyncEnabled(true)
	ebiten.SetTPS(TPS)

	g.Reset()
	g.Model = arche.New()

	p := params.CustomParams{
		Parameters: params.Default(),
	}
	if paramsFile != "" {
		err := p.FromJSON(paramsFile)
		if err != nil {
			return err
		}
		p.Parameters.WorkingDirectory.Path = filepath.Dir(paramsFile)
	}

	p.Parameters.Termination.MaxTicks = 0

	model.Default(&p, g.Model)

	ecs.AddResource(&g.Model.World, &res.GameSpeed{
		Speeds:     []uint16{5, 7, 10, 15, 30, 60, 120, 240, 480, 1000, 9999},
		SpeedIndex: 4,
		Pause:      true,
	})

	ecs.AddResource(&g.Model.World, &res.GameTick{})

	ecs.AddResource(&g.Model.World, &g.Screen)
	ecs.AddResource(&g.Model.World, &g.Mouse)

	sprites := res.NewSprites(data.Images, "images")
	ecs.AddResource(&g.Model.World, &sprites)

	fonts := res.NewFonts(data.Fonts, "fonts")
	ecs.AddResource(&g.Model.World, &fonts)

	for name, value := range overwriteParams {
		err := model.SetParameter(&g.Model.World, name, value)
		if err != nil {
			return err
		}
	}

	g.Model.AddSystem(&sys.InitUI{
		ResetFn: func(parameters map[string]any) {
			run(g, paramsFile, parameters)
		},
		LayoutData: data.Layouts,
		Layout:     "default",
	})

	g.Systems = append(g.Systems, &sys.UpdateUI{})

	g.Systems = append(g.Systems, &sys.Tick{})

	g.Systems = append(g.Systems, &sys.GameControls{
		PauseKey:      ebiten.KeySpace,
		StepKey:       ebiten.KeyArrowRight,
		SlowerKey:     '[',
		FasterKey:     ']',
		FullscreenKey: ebiten.KeyF11,
	})

	g.Model.AddUISystem(&sys.RenderUI{})

	g.Model.Initialize()
	g.InitializeRun()

	return nil
}
