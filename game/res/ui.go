package res

import (
	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mlange-42/arche/ecs"
)

// UI resource.Represents the complete game UI.
type UI struct {
	ui    *ebitenui.UI
	fonts *Fonts
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

func NewUI(world *ecs.World, fonts *Fonts) UI {
	ui := UI{
		fonts: fonts,
	}

	rootContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewStackedLayout()),
	)

	eui := ebitenui.UI{
		Container: rootContainer,
	}
	ui.ui = &eui

	return ui
}
