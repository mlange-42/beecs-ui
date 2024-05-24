package ui

import (
	"image/color"

	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type ImagePanel struct {
	Container *widget.Container
	Graphic   *widget.Graphic
}

func (p *ImagePanel) Update() {
	if len(p.Container.Children()) > 0 {
		p.Container.RemoveChild(p.Graphic)
		return
	}
	ssx, ssy := p.Container.GetWidget().Rect.Dx(), p.Container.GetWidget().Rect.Dy()
	img := ebiten.NewImage(ssx-8, ssy-8)
	img.Fill(color.RGBA{180, 180, 180, 255})
	vector.StrokeCircle(img,
		float32(ssx/2), float32(ssy/2),
		50, 2, color.RGBA{255, 0, 0, 255}, false)
	p.Graphic.Image = img

	p.Container.RemoveChild(p.Graphic)
	p.Container.AddChild(p.Graphic)
}

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

	ui.imageGrid = scroll

	return scroll
}
