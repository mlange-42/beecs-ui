package ui

import "github.com/ebitenui/ebitenui/widget"

func (ui *UI) createLeftPanel() *widget.Container {
	scroll, content := ui.scrollPanel(260)

	initialBeesSlider := ui.parameterSliderI(0, 50_000, 0, "params.InitialPopulation.Count")

	content.AddChild(initialBeesSlider)

	return scroll
}
