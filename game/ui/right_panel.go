package ui

import "github.com/ebitenui/ebitenui/widget"

func (ui *UI) createRightPanel() *widget.Container {
	scroll, content := ui.scrollPanel(0)

	root := widget.NewContainer(
		//widget.ContainerOpts.BackgroundImage(ui.sprites.Background),
		gridLayout([]bool{true, true}, []bool{true}, 4, 0),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Stretch: true,
			}),
			widget.WidgetOpts.MinSize(200, 10),
		),
	)

	root.AddChild(ui.imagePanel())
	root.AddChild(ui.imagePanel())
	root.AddChild(ui.imagePanel())
	root.AddChild(ui.imagePanel())
	content.AddChild(root)

	return scroll
}
