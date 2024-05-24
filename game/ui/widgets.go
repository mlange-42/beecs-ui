package ui

import (
	"github.com/ebitenui/ebitenui/widget"
)

func (ui *UI) defaultButtonImage() *widget.ButtonImage {
	return &widget.ButtonImage{
		Idle:    ui.sprites.Background,
		Hover:   ui.sprites.BackgroundHover,
		Pressed: ui.sprites.BackgroundPressed,
	}
}

func rowLayout(d widget.Direction, space int) widget.ContainerOpt {
	return widget.ContainerOpts.Layout(
		widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(d),
			widget.RowLayoutOpts.Padding(widget.NewInsetsSimple(space)),
			widget.RowLayoutOpts.Spacing(space),
		),
	)
}

func gridLayout(hStretch []bool, vStretch []bool, space int, pad int) widget.ContainerOpt {
	return widget.ContainerOpts.Layout(
		widget.NewGridLayout(
			widget.GridLayoutOpts.Columns(len(hStretch)),
			widget.GridLayoutOpts.Stretch(hStretch, vStretch),
			widget.GridLayoutOpts.Padding(widget.NewInsetsSimple(pad)),
			widget.GridLayoutOpts.Spacing(space, space),
		),
	)
}

func (ui *UI) button(text string, handler func(args *widget.ButtonClickedEventArgs)) *widget.Button {
	return widget.NewButton(
		widget.ButtonOpts.Image(ui.defaultButtonImage()),
		widget.ButtonOpts.Text(text, ui.fonts.Default, &widget.ButtonTextColor{
			Idle: ui.sprites.TextColor,
		}),
		widget.ButtonOpts.TextPadding(widget.NewInsetsSimple(5)),
		widget.ButtonOpts.ClickedHandler(handler),
	)
}

func (ui *UI) label(text string) *widget.Text {
	return widget.NewText(
		widget.TextOpts.Text(text, ui.fonts.Default, ui.sprites.TextColor),
		widget.TextOpts.Position(widget.TextPositionStart, widget.TextPositionCenter),
		widget.TextOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
			}),
		),
	)
}

func (ui *UI) slider(min, max, value int, width int, handler func(args *widget.SliderChangedEventArgs)) *widget.Slider {
	slider := widget.NewSlider(
		widget.SliderOpts.Direction(widget.DirectionHorizontal),
		widget.SliderOpts.MinMax(min, max),
		widget.SliderOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
				Stretch:  false,
			}),
			widget.WidgetOpts.MinSize(width, 6),
		),
		widget.SliderOpts.Images(
			&widget.SliderTrackImage{
				Idle:  ui.sprites.BackgroundHover,
				Hover: ui.sprites.BackgroundHover,
			},
			&widget.ButtonImage{
				Idle:    ui.sprites.BackgroundPressed,
				Hover:   ui.sprites.BackgroundPressed,
				Pressed: ui.sprites.BackgroundPressed,
			},
		),
		widget.SliderOpts.FixedHandleSize(6),
		widget.SliderOpts.TrackPadding(widget.Insets{Top: -4, Bottom: -4, Left: 4, Right: 12}),
		widget.SliderOpts.PageSizeFunc(func() int {
			return 1
		}),
		widget.SliderOpts.ChangedHandler(handler),
	)
	slider.Current = value

	return slider
}
