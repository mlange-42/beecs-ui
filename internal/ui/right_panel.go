package ui

import (
	"github.com/ebitenui/ebitenui/widget"
	"github.com/mlange-42/beecs-ui/internal/config"
)

func (ui *UI) createRightPanel(layout *config.Layout) *widget.Container {
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

func (ui *UI) createLayoutRow(row *config.LayoutRow) *widget.Container {
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
		root.AddChild(ui.imagePanel(&p, row.Height))
	}

	return root
}
