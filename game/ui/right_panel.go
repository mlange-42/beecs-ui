package ui

import (
	"github.com/ebitenui/ebitenui/widget"
)

func (ui *UI) createRightPanel(layout *Layout) *widget.Container {
	scroll, content := ui.scrollPanel(0)

	root := widget.NewContainer(
		//gridLayout([]bool{true}, []bool{false}, 4, 0),
		rowLayout(widget.DirectionVertical, 4, 0),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Stretch: true,
			}),
			widget.WidgetOpts.MinSize(200, 10),
		),
	)

	for _, row := range layout.Rows {
		root.AddChild(ui.createLayoutRow(&row))
	}

	content.AddChild(root)

	ui.imageGrid = scroll

	return scroll
}

func (ui *UI) createLayoutRow(row *LayoutRow) *widget.Container {
	cols := len(row.Panels)
	stretch := make([]bool, cols)
	for i := range stretch {
		stretch[i] = true
	}

	root := widget.NewContainer(
		gridLayout(stretch, []bool{true}, 4, 0),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Stretch: true,
			}),
			widget.WidgetOpts.MinSize(200, 10),
		),
	)

	for _, p := range row.Panels {
		_ = p
		root.AddChild(ui.imagePanel(p.Drawer, row.Height))
	}

	return root
}
