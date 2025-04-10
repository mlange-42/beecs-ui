package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mlange-42/ark-tools/app"
	"github.com/mlange-42/ark/ecs"
	"github.com/mlange-42/beecs-ui/internal/res"
)

// Game container
type Game struct {
	Model  *app.App
	Screen res.Screen
	Mouse  res.Mouse

	Systems   []app.System
	gameSpeed *res.GameSpeed

	canvasHelper *canvasHelper
}

// NewGame returns a new game
func NewGame(mod *app.App) Game {
	return Game{
		Model:        mod,
		Screen:       res.Screen{Image: nil, Width: 0, Height: 0},
		canvasHelper: newCanvasHelper(),
	}
}

// Initialize the game.
func (g *Game) Initialize() {
	ebiten.SetWindowSize(1080, 720)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("beecs-ui")
}

// Initialize a new run.
func (g *Game) InitializeRun() {
	g.gameSpeed = ecs.GetResource[res.GameSpeed](&g.Model.World)
	for _, sys := range g.Systems {
		sys.Initialize(&g.Model.World)
	}
}

// Run the game.
func (g *Game) Run() error {
	if err := ebiten.RunGame(g); err != nil {
		return err
	}
	return nil
}

// Layout the game.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	s := ebiten.DeviceScaleFactor()
	return int(float64(outsideWidth) * s), int(float64(outsideHeight) * s)
}

// Update the game.
func (g *Game) Update() error {
	g.updateMouse()

	if !g.gameSpeed.Pause {
		g.Model.Update()
	}

	for _, sys := range g.Systems {
		sys.Update(&g.Model.World)
	}

	return nil
}

// Draw the game.
func (g *Game) Draw(screen *ebiten.Image) {
	g.Screen.Image = screen
	g.Screen.Width = screen.Bounds().Dx()
	g.Screen.Height = screen.Bounds().Dy()
	g.Model.UpdateUI()
}

// Reset the game.
func (g *Game) Reset() {
	g.Systems = g.Systems[:0]
}

func (g *Game) updateMouse() {
	g.Mouse.IsInside = g.canvasHelper.isMouseInside(g.Screen.Width, g.Screen.Height)
}
