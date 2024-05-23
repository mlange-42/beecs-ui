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
	fonts   *res.Fonts
	sprites *res.Sprites

	InfoLabel *widget.Text
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

func New(world *ecs.World, fonts *res.Fonts, sprites *res.Sprites) UI {
	ui := UI{
		fonts:   fonts,
		sprites: sprites,
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
		widget.ContainerOpts.BackgroundImage(ui.sprites.BackgroundNineSlice),
		rowLayout(widget.DirectionVertical),
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

func (ui *UI) createTopBar() *widget.Container {
	root := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(ui.sprites.BackgroundNineSlice),
		rowLayout(widget.DirectionHorizontal),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position:  widget.RowLayoutPositionStart,
				Stretch:   true,
				MaxHeight: 40,
			}),
			widget.WidgetOpts.MinSize(40, 10),
		),
	)

	labels := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(ui.sprites.BackgroundNineSlice),
		rowLayout(widget.DirectionHorizontal),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionStart,
				Stretch:  true,
			}),
			widget.WidgetOpts.MinSize(200, 10),
		),
	)

	buttons := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(ui.sprites.BackgroundNineSlice),
		rowLayout(widget.DirectionHorizontal),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionStart,
				Stretch:  true,
			}),
			widget.WidgetOpts.MinSize(40, 10),
		),
	)

	ui.InfoLabel = widget.NewText(
		widget.TextOpts.Text("", ui.fonts.Default, ui.sprites.TextColor),
		widget.TextOpts.Position(widget.TextPositionStart, widget.TextPositionCenter),
	)

	labels.AddChild(ui.InfoLabel)

	root.AddChild(labels)
	root.AddChild(buttons)

	return root
}

func rowLayout(d widget.Direction) widget.ContainerOpt {
	return widget.ContainerOpts.Layout(
		widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(d),
			widget.RowLayoutOpts.Padding(widget.NewInsetsSimple(4)),
			widget.RowLayoutOpts.Spacing(6),
		),
	)
}

func gridLayout(c int) widget.ContainerOpt {
	return widget.ContainerOpts.Layout(
		widget.NewGridLayout(
			widget.GridLayoutOpts.Columns(c),
			widget.GridLayoutOpts.Padding(widget.NewInsetsSimple(4)),
			widget.GridLayoutOpts.Spacing(6, 6),
		),
	)
}
