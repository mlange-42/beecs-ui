package ui

import "github.com/ebitenui/ebitenui/widget"

func (ui *UI) createLeftPanel() *widget.Container {
	scroll, content := ui.scrollPanel(260)

	pars := Parameters{
		Sections: []ParameterSection{
			{
				Text: "Initialization",
				Parameters: []Parameter{
					{
						Parameter: "params.InitialPopulation.Count",
						SliderInt: &SliderInt{
							Min: 1_000,
							Max: 50_000,
						},
					},
				},
			},
			{
				Text: "HandlingTime",
				Parameters: []Parameter{
					{
						Parameter: "params.HandlingTime.ConstantHandlingTime",
						Toggle:    &Toggle{},
					},
					{
						Parameter: "params.HandlingTime.NectarGathering",
						SliderFloat: &SliderFloat{
							Min:       60,
							Max:       3600,
							Precision: 1,
						},
					},
					{
						Parameter: "params.HandlingTime.PollenGathering",
						SliderFloat: &SliderFloat{
							Min:       60,
							Max:       3600,
							Precision: 1,
						},
					},
				},
			},
		},
	}

	content.AddChild(ui.crateParameters(pars))

	return scroll
}

func (ui *UI) crateParameters(p Parameters) *widget.Container {
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

	for _, sec := range p.Sections {
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
