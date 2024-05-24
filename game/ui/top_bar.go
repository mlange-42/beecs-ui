package ui

import "github.com/ebitenui/ebitenui/widget"

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
