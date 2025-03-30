package ui

import (
	"fmt"
	"image"
	"io/fs"
	"log"
	"os"
	"strings"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mlange-42/ark/ecs"
	"github.com/mlange-42/beecs-ui/internal/config"
	"github.com/mlange-42/beecs-ui/internal/res"
)

// UI resource.Represents the complete game UI.
type UI struct {
	ui      *ebitenui.UI
	world   *ecs.World
	time    *res.GameTick
	fonts   *res.Fonts
	sprites *res.Sprites
	speed   *res.GameSpeed

	InfoLabel   *widget.Text
	SpeedLabel  *widget.Text
	PauseButton *widget.Button

	properties    []ParameterProperty
	images        []ImagePanel
	imageGrid     *widget.Container
	gridSize      image.Point
	layoutUpdated bool

	resetFn func(parameters map[string]any, speed uint8)
}

func New(world *ecs.World, data fs.FS, layout string, resetFn func(parameters map[string]any, speed uint8)) UI {
	ui := UI{
		world:   world,
		time:    ecs.GetResource[res.GameTick](world),
		fonts:   ecs.GetResource[res.Fonts](world),
		sprites: ecs.GetResource[res.Sprites](world),
		speed:   ecs.GetResource[res.GameSpeed](world),
		resetFn: resetFn,
	}

	rootContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)

	var lay *config.Layout
	var err error
	if strings.HasSuffix(layout, ".json") {
		lay, err = config.FromFile(os.DirFS("."), layout,
			data, "layouts/templates/params.json", "layouts/templates/panels.json")
	} else {
		lay, err = config.FromFile(data, fmt.Sprintf("layouts/%s.json", layout),
			data, "layouts/templates/params.json", "layouts/templates/panels.json")
	}
	if err != nil {
		log.Fatal(err)
	}

	rootContainer.AddChild(ui.createUI(lay))

	eui := ebitenui.UI{
		Container: rootContainer,
	}
	ui.ui = &eui

	for i := range ui.images {
		if err := ui.images[i].Initialize(world); err != nil {
			log.Fatal(err)
		}
	}

	return ui
}

func (ui *UI) UI() *ebitenui.UI {
	return ui.ui
}

func (ui *UI) Update() {
	if !ui.speed.Pause {
		for i := range ui.images {
			ui.images[i].Update(ui.world)
		}
	}

	ui.UI().Update()
}

func (ui *UI) Draw(screen *ebiten.Image) {
	sx, sy := ui.imageGrid.GetWidget().Rect.Dx(), ui.imageGrid.GetWidget().Rect.Dy()
	resize := ui.layoutUpdated || ui.gridSize.X != sx || ui.gridSize.Y != sy

	if resize {
		for i := range ui.images {
			ui.images[i].Resize()
		}
		ui.gridSize.X = sx
		ui.gridSize.Y = sy

		ui.layoutUpdated = !ui.layoutUpdated
	}

	for i := range ui.images {
		ui.images[i].Draw(ui.world)
	}

	ui.UI().Draw(screen)
}

func (ui *UI) createUI(layout *config.Layout) *widget.Container {
	root := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(ui.sprites.Background),
		gridLayout([]bool{true}, []bool{false, true}, 4, 0),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionEnd,
				VerticalPosition:   widget.AnchorLayoutPositionStart,
				StretchHorizontal:  true,
				StretchVertical:    true,
			}),
			widget.WidgetOpts.MinSize(40, 10),
		),
	)

	root.AddChild(ui.createTopBar())
	root.AddChild(ui.createMainPanel(layout))

	return root
}

func (ui *UI) createMainPanel(layout *config.Layout) *widget.Container {
	root := widget.NewContainer(
		//widget.ContainerOpts.BackgroundImage(ui.sprites.Background),
		gridLayout([]bool{false, true}, []bool{true}, 4, 0),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.GridLayoutData{}),
			widget.WidgetOpts.MinSize(40, 10),
		),
	)

	root.AddChild(ui.createLeftPanel(layout))
	root.AddChild(ui.createRightPanel(layout))

	return root
}
