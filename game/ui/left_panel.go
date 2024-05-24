package ui

import "github.com/ebitenui/ebitenui/widget"

func (ui *UI) createLeftPanel() *widget.Container {
	scroll, content := ui.scrollPanel(260)

	content.AddChild(ui.parameterSliderI(0, 50_000, "params.InitialPopulation.Count"))
	content.AddChild(ui.parameterSliderF(100, 2000, 1, "params.HandlingTime.NectarGathering"))
	content.AddChild(ui.parameterSliderF(100, 2000, 1, "params.HandlingTime.PollenGathering"))

	return scroll
}
