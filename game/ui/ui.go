package ui

import (
	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/beecs-ui/game/res"
)

// UI resource.Represents the complete game UI.
type UI struct {
	ui      *ebitenui.UI
	time    *res.GameTick
	fonts   *res.Fonts
	sprites *res.Sprites
	speed   *res.GameSpeed

	InfoLabel   *widget.Text
	SpeedLabel  *widget.Text
	PauseButton *widget.Button

	resetFn func()
}

func (ui *UI) UI() *ebitenui.UI {
	return ui.ui
}

func (ui *UI) Update() {
	ui.UI().Update()
}

func (ui *UI) Draw(screen *ebiten.Image) {
	ui.UI().Draw(screen)
}

func New(world *ecs.World, time *res.GameTick, fonts *res.Fonts, sprites *res.Sprites, speed *res.GameSpeed, resetFn func()) UI {
	ui := UI{
		time:    time,
		fonts:   fonts,
		sprites: sprites,
		speed:   speed,
		resetFn: resetFn,
	}

	rootContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)

	rootContainer.AddChild(ui.createUI())

	eui := ebitenui.UI{
		Container: rootContainer,
	}
	ui.ui = &eui

	return ui
}

func (ui *UI) createUI() *widget.Container {
	root := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(ui.sprites.Background),
		gridLayout([]bool{true}, []bool{false, true}, 4, 0),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionEnd,
				VerticalPosition:   widget.AnchorLayoutPositionStart,
				StretchHorizontal:  true,
				StretchVertical:    true,
			}),
			widget.WidgetOpts.MinSize(40, 10),
		),
	)

	root.AddChild(ui.createTopBar())

	return root
}
