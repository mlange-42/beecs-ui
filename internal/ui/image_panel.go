package ui

import (
	"image"
	"log"

	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mlange-42/ark/ecs"
	"github.com/mlange-42/beecs-ui/internal/plot"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/vgimg"
)

const imageMargin = 12

var scale = calcScaleCorrection()

type ImagePanel struct {
	Container *widget.Container
	Graphic   *widget.Graphic
	Drawer    plot.Drawer
	Observer  any
	canvas    *vgimg.Canvas
}

func (p *ImagePanel) Initialize(world *ecs.World) error {
	return p.Drawer.Initialize(world, p.Observer)
}

func (p *ImagePanel) Update(world *ecs.World) {
	p.Drawer.Update(world)
}

func (p *ImagePanel) Resize() {
	if len(p.Container.Children()) > 0 {
		p.Container.RemoveChild(p.Graphic)
		return
	}
	ssx, ssy := p.Container.GetWidget().Rect.Dx(), p.Container.GetWidget().Rect.Dy()
	w, h := float64(ssx-2*imageMargin), float64(ssy-2*imageMargin)

	p.canvas = vgimg.New(
		vg.Points(w*scale),
		vg.Points(h*scale))
	tempImg := p.canvas.Image()

	img := ebiten.NewImage(tempImg.Bounds().Dx(), tempImg.Bounds().Dy())
	p.Graphic.Image = img

	p.Container.AddChild(p.Graphic)
	p.Drawer.SetChanged()
}

func (p *ImagePanel) Draw(world *ecs.World) {
	if p.Drawer.Draw(world, p.canvas) {
		img := p.canvas.Image()

		rgb, ok := img.(*image.RGBA)
		if !ok {
			log.Fatal("not an RGBA image")
		}
		p.Graphic.Image.WritePixels(rgb.Pix)
	}
}

// Calculate scale correction for scaled monitors.
func calcScaleCorrection() float64 {
	width := 100.0
	c := vgimg.New(vg.Points(width), vg.Points(width))
	img := c.Image()
	return width / float64(img.Bounds().Dx())
}
