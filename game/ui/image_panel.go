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

type Layout struct {
	Rows []LayoutRow
}

type LayoutRow struct {
	Height int
	Panels []LayoutPanel
}

type LayoutPanel struct{}
