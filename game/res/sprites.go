package res

import (
	"image/color"
	"io/fs"
	"log"
	"path"

	"github.com/ebitenui/ebitenui/image"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	background        = "ui_panel/ui_panel.png"
	backgroundHover   = "ui_panel/ui_panel_hover.png"
	backgroundPressed = "ui_panel/ui_panel_pressed.png"
)

// Sprites holds all tileset data.
type Sprites struct {
	Background        *image.NineSlice
	BackgroundHover   *image.NineSlice
	BackgroundPressed *image.NineSlice

	TextColor color.RGBA
}

// NewSprites creates a new Sprites resource from the given tileset folder.
func NewSprites(fSys fs.FS, dir string) Sprites {
	bg, _, err := ebitenutil.NewImageFromFileSystem(fSys, path.Join(dir, background))
	if err != nil {
		log.Fatal("error reading image: ", err)
	}

	bgHover, _, err := ebitenutil.NewImageFromFileSystem(fSys, path.Join(dir, backgroundHover))
	if err != nil {
		log.Fatal("error reading image: ", err)
	}

	bgPressed, _, err := ebitenutil.NewImageFromFileSystem(fSys, path.Join(dir, backgroundPressed))
	if err != nil {
		log.Fatal("error reading image: ", err)
	}

	return Sprites{
		Background:        image.NewNineSliceSimple(bg, bg.Bounds().Dx()/4, bg.Bounds().Dx()/2),
		BackgroundHover:   image.NewNineSliceSimple(bgHover, bgHover.Bounds().Dx()/4, bgHover.Bounds().Dx()/2),
		BackgroundPressed: image.NewNineSliceSimple(bgPressed, bgPressed.Bounds().Dx()/4, bgPressed.Bounds().Dx()/2),
		TextColor:         color.RGBA{0, 0, 0, 255},
	}
}
