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

func (ui *UI) createTopBar() *widget.Container {
	root := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(ui.sprites.Background),
		gridLayout([]bool{true, false}, []bool{true}, 4, 0),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.GridLayoutData{}),
			widget.WidgetOpts.MinSize(40, 10),
		),
	)

	root.AddChild(ui.createTopBarLabels())
	root.AddChild(ui.createTopBarButtons())

	return root
}

func (ui *UI) createTopBarLabels() *widget.Container {
	labels := widget.NewContainer(
		//widget.ContainerOpts.BackgroundImage(ui.sprites.Background),
		rowLayout(widget.DirectionHorizontal, 4),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.GridLayoutData{}),
			widget.WidgetOpts.MinSize(200, 10),
		),
	)
	ui.InfoLabel = widget.NewText(
		widget.TextOpts.Text("", ui.fonts.Default, ui.sprites.TextColor),
		widget.TextOpts.Position(widget.TextPositionStart, widget.TextPositionCenter),
		widget.TextOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.RowLayoutData{
			Position: widget.RowLayoutPositionCenter,
		})),
	)

	labels.AddChild(ui.InfoLabel)
	return labels
}

func (ui *UI) createTopBarButtons() *widget.Container {
	buttons := widget.NewContainer(
		//widget.ContainerOpts.BackgroundImage(ui.sprites.Background),
		rowLayout(widget.DirectionHorizontal, 4),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.GridLayoutData{}),
			widget.WidgetOpts.MinSize(40, 10),
		),
	)

	resetButton := ui.button("<<", func(args *widget.ButtonClickedEventArgs) {
		ui.resetFn()
	})

	ui.PauseButton = ui.button(">>", func(args *widget.ButtonClickedEventArgs) {
		ui.speed.Pause = !ui.speed.Pause
		ui.speed.NextPause = -1
	})

	stepButton := ui.button(">", func(args *widget.ButtonClickedEventArgs) {
		ui.speed.Pause = false
		ui.speed.NextPause = ui.time.Tick + 1
	})

	stepMonthButton := ui.button(">M", func(args *widget.ButtonClickedEventArgs) {
		ui.speed.Pause = false
		ui.speed.NextPause = ui.time.Tick + 30
	})

	stepYearButton := ui.button(">Y", func(args *widget.ButtonClickedEventArgs) {
		ui.speed.Pause = false
		ui.speed.NextPause = ui.time.Tick + 365
	})

	buttons.AddChild(resetButton)
	buttons.AddChild(ui.PauseButton)
	buttons.AddChild(stepButton)
	buttons.AddChild(stepMonthButton)
	buttons.AddChild(stepYearButton)

	return buttons
}

func (ui *UI) defaultButtonImage() *widget.ButtonImage {
	return &widget.ButtonImage{
		Idle:    ui.sprites.Background,
		Hover:   ui.sprites.BackgroundHover,
		Pressed: ui.sprites.BackgroundPressed,
	}
}

func rowLayout(d widget.Direction, space int) widget.ContainerOpt {
	return widget.ContainerOpts.Layout(
		widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(d),
			widget.RowLayoutOpts.Padding(widget.NewInsetsSimple(space)),
			widget.RowLayoutOpts.Spacing(space),
		),
	)
}

func gridLayout(hStretch []bool, vStretch []bool, space int, pad int) widget.ContainerOpt {
	return widget.ContainerOpts.Layout(
		widget.NewGridLayout(
			widget.GridLayoutOpts.Columns(len(hStretch)),
			widget.GridLayoutOpts.Stretch(hStretch, vStretch),
			widget.GridLayoutOpts.Padding(widget.NewInsetsSimple(pad)),
			widget.GridLayoutOpts.Spacing(space, space),
		),
	)
}

func (ui *UI) button(text string, handler func(args *widget.ButtonClickedEventArgs)) *widget.Button {
	return widget.NewButton(
		widget.ButtonOpts.Image(ui.defaultButtonImage()),
		widget.ButtonOpts.Text(text, ui.fonts.Default, &widget.ButtonTextColor{
			Idle: ui.sprites.TextColor,
		}),
		widget.ButtonOpts.TextPadding(widget.NewInsetsSimple(5)),
		widget.ButtonOpts.ClickedHandler(handler),
	)
}
