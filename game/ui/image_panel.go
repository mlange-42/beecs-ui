package ui

import (
	"image"
	"image/color"
	"log"

	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/beecs-ui/game/plot"
	"golang.org/x/image/colornames"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/vgimg"
)

var scale = calcScaleCorrection()

var defaultColors = []color.Color{
	colornames.Blue,
	colornames.Orange,
	colornames.Green,
	colornames.Purple,
	colornames.Red,
	colornames.Turquoise,
}

type ImagePanel struct {
	Container *widget.Container
	Graphic   *widget.Graphic
	Drawer    plot.Drawer
	canvas    *vgimg.Canvas
}

func (p *ImagePanel) Initialize(world *ecs.World) {
	p.Drawer.Initialize(world)
}

func (p *ImagePanel) Update(world *ecs.World) {
	p.Drawer.Update(world)
}

func (p *ImagePanel) Draw(world *ecs.World, resize bool) {
	if resize {
		if len(p.Container.Children()) > 0 {
			p.Container.RemoveChild(p.Graphic)
			return
		}
		ssx, ssy := p.Container.GetWidget().Rect.Dx(), p.Container.GetWidget().Rect.Dy()
		w, h := float64(ssx-8), float64(ssy-8)

		p.canvas = vgimg.New(
			vg.Points(w*scale),
			vg.Points(h*scale))
		tempImg := p.canvas.Image()

		img := ebiten.NewImage(tempImg.Bounds().Dx(), tempImg.Bounds().Dy())
		p.Graphic.Image = img

		p.Container.RemoveChild(p.Graphic)
		p.Container.AddChild(p.Graphic)
	}

	p.Drawer.Draw(world, p.canvas)

	img := p.canvas.Image()

	rgb, ok := img.(*image.RGBA)
	if !ok {
		log.Fatal("not an RGBA image")
	}
	p.Graphic.Image.WritePixels(rgb.Pix)
}

type Layout struct {
	Rows []LayoutRow
}

type LayoutRow struct {
	Height int
	Panels []LayoutPanel
}

type LayoutPanel struct {
	Drawer plot.Drawer
}

// Calculate scale correction for scaled monitors.
func calcScaleCorrection() float64 {
	width := 100.0
	c := vgimg.New(vg.Points(width), vg.Points(width))
	img := c.Image()
	return width / float64(img.Bounds().Dx())
}
