package ui

import (
	"image/color"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mlange-42/beecs-ui/internal/config"
	"github.com/mlange-42/beecs/model"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/vgimg"
)

func (ui *UI) defaultButtonImage() *widget.ButtonImage {
	return &widget.ButtonImage{
		Idle:    ui.sprites.Background,
		Hover:   ui.sprites.BackgroundHover,
		Pressed: ui.sprites.BackgroundPressed,
	}
}

func rowLayout(d widget.Direction, space int, pad int) widget.ContainerOpt {
	return widget.ContainerOpts.Layout(
		widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(d),
			widget.RowLayoutOpts.Padding(widget.NewInsetsSimple(pad)),
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

func (ui *UI) label(text string, hPos widget.TextPosition) *widget.Text {
	return widget.NewText(
		widget.TextOpts.Text(text, ui.fonts.Default, ui.sprites.TextColor),
		widget.TextOpts.Position(hPos, widget.TextPositionCenter),
		widget.TextOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
			}),
		),
	)
}

func (ui *UI) slider(min, max, value int, width int, stretchRow bool, handler func(args *widget.SliderChangedEventArgs)) *widget.Slider {
	slider := widget.NewSlider(
		widget.SliderOpts.Direction(widget.DirectionHorizontal),
		widget.SliderOpts.MinMax(min, max),
		widget.SliderOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
				Stretch:  stretchRow,
			}),
			widget.WidgetOpts.MinSize(width, 6),
		),
		widget.SliderOpts.Images(
			&widget.SliderTrackImage{
				Idle:  ui.sprites.BackgroundPressed,
				Hover: ui.sprites.BackgroundPressed,
			},
			&widget.ButtonImage{
				Idle:    ui.sprites.BackgroundHover,
				Hover:   ui.sprites.BackgroundHover,
				Pressed: ui.sprites.BackgroundHover,
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

func (ui *UI) parameterSliderF(parameter string, units string, conf *config.SliderFloat) *widget.Container {
	root := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(ui.sprites.Background),
		rowLayout(widget.DirectionVertical, 4, 4),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Stretch: true,
			}),
			widget.WidgetOpts.MinSize(200, 10),
		),
	)
	labels := widget.NewContainer(
		gridLayout([]bool{true, false}, []bool{true}, 4, 0),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Stretch: true,
			}),
			widget.WidgetOpts.MinSize(200, 10),
		),
	)
	sliders := widget.NewContainer(
		rowLayout(widget.DirectionVertical, 0, 4),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Stretch: true,
			}),
			widget.WidgetOpts.MinSize(200, 12),
		),
	)

	label := ui.label(strings.Replace(parameter, "params.", "", 1), widget.TextPositionStart)
	valueLabel := ui.label("100", widget.TextPositionEnd)

	v, err := model.GetParameter(ui.world, parameter)
	if err != nil {
		log.Fatal(err)
	}
	vv, ok := v.(float64)
	if !ok {
		log.Fatal("error converting parameter value to float")
	}

	valueLabel.Label = strconv.FormatFloat(vv, 'f', -1, 64) + units
	value := int(vv * conf.Precision)
	slider := ui.slider(int(conf.Min*conf.Precision), int(conf.Max*conf.Precision), value, 0, true, func(args *widget.SliderChangedEventArgs) {
		v := float64(args.Current) / conf.Precision
		err := model.SetParameter(ui.world, parameter, v)
		valueLabel.Label = strconv.FormatFloat(v, 'f', -1, 64) + units
		if err != nil {
			log.Fatal(err)
		}
	})

	labels.AddChild(label)
	labels.AddChild(valueLabel)

	sliders.AddChild(slider)

	root.AddChild(labels)
	root.AddChild(sliders)

	ui.properties = append(ui.properties, &SliderPropertyFloat{
		name:      parameter,
		slider:    slider,
		precision: conf.Precision,
	})

	return root
}

func (ui *UI) parameterSliderI(parameter string, units string, conf *config.SliderInt) *widget.Container {
	root := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(ui.sprites.Background),
		rowLayout(widget.DirectionVertical, 4, 4),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Stretch: true,
			}),
			widget.WidgetOpts.MinSize(200, 10),
		),
	)
	labels := widget.NewContainer(
		gridLayout([]bool{true, false}, []bool{true}, 4, 0),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Stretch: true,
			}),
			widget.WidgetOpts.MinSize(200, 10),
		),
	)
	sliders := widget.NewContainer(
		rowLayout(widget.DirectionVertical, 0, 0),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Stretch: true,
			}),
			widget.WidgetOpts.MinSize(200, 12),
		),
	)

	label := ui.label(strings.Replace(parameter, "params.", "", 1), widget.TextPositionStart)
	valueLabel := ui.label("100", widget.TextPositionEnd)

	v, err := model.GetParameter(ui.world, parameter)
	if err != nil {
		log.Fatal(err)
	}
	value, ok := v.(int64)
	if !ok {
		log.Fatal("error converting parameter value to int")
	}

	valueLabel.Label = strconv.Itoa(int(value)) + units
	slider := ui.slider(conf.Min, conf.Max, int(value), 0, true, func(args *widget.SliderChangedEventArgs) {
		err := model.SetParameter(ui.world, parameter, args.Current)
		valueLabel.Label = strconv.Itoa(args.Current) + units
		if err != nil {
			log.Fatal(err)
		}
	})

	labels.AddChild(label)
	labels.AddChild(valueLabel)

	sliders.AddChild(slider)

	root.AddChild(labels)
	root.AddChild(sliders)

	ui.properties = append(ui.properties, &SliderPropertyInt{
		name:   parameter,
		slider: slider,
	})

	return root
}

func (ui *UI) parameterToggle(parameter string) *widget.Container {
	root := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(ui.sprites.Background),
		gridLayout([]bool{true, false}, []bool{true}, 4, 4),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Stretch: true,
			}),
			widget.WidgetOpts.MinSize(200, 10),
		),
	)

	label := ui.label(strings.Replace(parameter, "params.", "", 1), widget.TextPositionStart)

	v, err := model.GetParameter(ui.world, parameter)
	if err != nil {
		log.Fatal(err)
	}
	value, ok := v.(bool)
	if !ok {
		log.Fatal("error converting parameter value to bool")
	}

	uncheckedImage := ebiten.NewImage(12, 12)
	uncheckedImage.Fill(color.Transparent)

	checkedImage := ebiten.NewImage(12, 12)
	checkedImage.Fill(ui.sprites.TextColor)

	toggle := widget.NewCheckbox(
		widget.CheckboxOpts.ButtonOpts(
			widget.ButtonOpts.WidgetOpts(
				widget.WidgetOpts.LayoutData(widget.RowLayoutData{}),
				widget.WidgetOpts.MinSize(20, 20),
			),
			widget.ButtonOpts.Image(ui.defaultButtonImage()),
		),
		widget.CheckboxOpts.Image(&widget.CheckboxGraphicImage{
			Unchecked: &widget.ButtonImageImage{
				Idle: uncheckedImage,
			},
			Checked: &widget.ButtonImageImage{
				Idle: checkedImage,
			},
		}),
		widget.CheckboxOpts.StateChangedHandler(func(args *widget.CheckboxChangedEventArgs) {
			value := args.State == widget.WidgetChecked
			model.SetParameter(ui.world, parameter, value)
		}),
	)

	if value {
		toggle.SetState(widget.WidgetChecked)
	}

	root.AddChild(label)
	root.AddChild(toggle)

	ui.properties = append(ui.properties, &ToggleProperty{
		name:   parameter,
		toggle: toggle,
	})

	return root
}

func (ui *UI) scrollPanel(width int) (*widget.Container, *widget.Container) {
	root := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(ui.sprites.Background),
		// the container will use an grid layout to layout its ScrollableContainer and Slider
		widget.ContainerOpts.Layout(widget.NewGridLayout(
			widget.GridLayoutOpts.Columns(2),
			widget.GridLayoutOpts.Spacing(2, 0),
			widget.GridLayoutOpts.Stretch([]bool{true, false}, []bool{true}),
		)),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.GridLayoutData{}),
			widget.WidgetOpts.MinSize(width, 10),
		),
	)

	content := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Spacing(4),
		)),
	)

	scrollContainer := widget.NewScrollContainer(
		widget.ScrollContainerOpts.Content(content),
		widget.ScrollContainerOpts.StretchContentWidth(),
		widget.ScrollContainerOpts.Padding(widget.NewInsetsSimple(2)),
		widget.ScrollContainerOpts.Image(&widget.ScrollContainerImage{
			Idle: ui.sprites.Background,
			Mask: ui.sprites.Background,
		}),
	)
	root.AddChild(scrollContainer)

	//Create a function to return the page size used by the slider
	pageSizeFunc := func() int {
		ps := int(math.Round(float64(scrollContainer.ViewRect().Dy()) / float64(content.GetWidget().Rect.Dy()) * 1000))
		return ps
	}

	vSlider := widget.NewSlider(
		widget.SliderOpts.Direction(widget.DirectionVertical),
		widget.SliderOpts.MinMax(0, 1000),
		widget.SliderOpts.PageSizeFunc(pageSizeFunc),
		widget.SliderOpts.ChangedHandler(func(args *widget.SliderChangedEventArgs) {
			scrollContainer.ScrollTop = float64(args.Slider.Current) / 1000
		}),
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
	)
	//Set the slider's position if the scrollContainer is scrolled by other means than the slider
	scrollContainer.GetWidget().ScrolledEvent.AddHandler(func(args interface{}) {
		a := args.(*widget.WidgetScrolledEventArgs)
		p := pageSizeFunc() / 3
		if p < 1 {
			p = 1
		}
		dy := math.Copysign(1.0, a.Y)
		vSlider.Current -= int(math.Round(dy * float64(p)))
	})

	root.AddChild(vSlider)

	return root, content
}

func (ui *UI) imagePanel(panel *config.LayoutPanel, height int) *widget.Container {
	img := ebiten.NewImage(1, 1)
	img.Fill(color.RGBA{180, 180, 180, 255})

	root := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(ui.sprites.Background),
		rowLayout(widget.DirectionVertical, imageMargin, imageMargin),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.GridLayoutData{}),
			widget.WidgetOpts.MinSize(10, height),
		),
	)

	graphic := widget.NewGraphic(
		widget.GraphicOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Stretch: true,
			}),
			widget.WidgetOpts.MinSize(10, 10),
		),
		widget.GraphicOpts.Image(img),
	)

	root.AddChild(graphic)

	ui.images = append(ui.images, ImagePanel{
		Container: root,
		Graphic:   graphic,
		Drawer:    panel.Drawer,
		Observer:  panel.Observer,
		canvas: vgimg.New(
			vg.Points(float64(img.Bounds().Dx())),
			vg.Points(float64(img.Bounds().Dy()))),
	})

	return root
}
