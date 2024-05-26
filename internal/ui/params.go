package ui

import "github.com/ebitenui/ebitenui/widget"

type ParameterProperty interface {
	Name() string
	Value() any
}

type SliderPropertyInt struct {
	name   string
	slider *widget.Slider
}

func (p *SliderPropertyInt) Name() string {
	return p.name
}

func (p *SliderPropertyInt) Value() any {
	return p.slider.Current
}

type SliderPropertyFloat struct {
	name      string
	slider    *widget.Slider
	precision float64
}

func (p *SliderPropertyFloat) Name() string {
	return p.name
}

func (p *SliderPropertyFloat) Value() any {
	return float64(p.slider.Current) / p.precision
}

type ToggleProperty struct {
	name   string
	toggle *widget.Checkbox
}

func (p *ToggleProperty) Name() string {
	return p.name
}

func (p *ToggleProperty) Value() any {
	return p.toggle.State() == widget.WidgetChecked
}
