package ui

import "github.com/ebitenui/ebitenui/widget"

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
