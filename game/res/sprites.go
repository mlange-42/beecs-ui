package res

import (
	"io/fs"
	"log"
	"path"

	"github.com/ebitenui/ebitenui/image"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	background = "ui_panel/ui_panel.png"
)

// Sprites holds all tileset data.
type Sprites struct {
	Background          *ebiten.Image
	BackgroundNineSlice *image.NineSlice
}

// NewSprites creates a new Sprites resource from the given tileset folder.
func NewSprites(fSys fs.FS, dir string) Sprites {
	background, _, err := ebitenutil.NewImageFromFileSystem(fSys, path.Join(dir, background))
	if err != nil {
		log.Fatal("error reading image: ", err)
	}
	w := background.Bounds().Dx()

	return Sprites{
		Background:          background,
		BackgroundNineSlice: image.NewNineSliceSimple(background, w/4, w/2),
	}
}
