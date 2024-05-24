package ui

import "github.com/ebitenui/ebitenui/widget"

func (ui *UI) createLeftPanel(layout *Layout) *widget.Container {
	scroll, content := ui.scrollPanel(260)

	content.AddChild(ui.crateParameters(layout.Parameters))

	return scroll
}

func (ui *UI) crateParameters(p []ParameterSection) *widget.Container {
	root := widget.NewContainer(
		//widget.ContainerOpts.BackgroundImage(ui.sprites.Background),
		rowLayout(widget.DirectionVertical, 4, 0),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Stretch: true,
			}),
			widget.WidgetOpts.MinSize(200, 10),
		),
	)

	for _, sec := range p {
		root.AddChild(ui.crateParametersSection(sec))
	}

	return root
}

func (ui *UI) crateParametersSection(p ParameterSection) *widget.Container {
	root := widget.NewContainer(
		//widget.ContainerOpts.BackgroundImage(ui.sprites.Background),
		rowLayout(widget.DirectionVertical, 4, 4),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Stretch: true,
			}),
			widget.WidgetOpts.MinSize(200, 10),
		),
	)

	root.AddChild(ui.label(p.Text, widget.TextPositionCenter))

	for _, par := range p.Parameters {
		if par.SliderFloat != nil {
			root.AddChild(ui.parameterSliderF(par.SliderFloat.Min, par.SliderFloat.Max, par.SliderFloat.Precision, par.Parameter))
		}
		if par.SliderInt != nil {
			root.AddChild(ui.parameterSliderI(par.SliderInt.Min, par.SliderInt.Max, par.Parameter))
		}
		if par.Toggle != nil {
			root.AddChild(ui.parameterToggle(par.Parameter))
		}
	}

	return root
}
