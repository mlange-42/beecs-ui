package res

import (
	"io/fs"
	"log"
	"path"

	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/font/sfnt"
)

const fontFile = "WorkSans-Regular.ttf"

const fontSize = 18

// Fonts resource for access to UI fonts.
type Fonts struct {
	Default font.Face
}

func NewFonts(fSys fs.FS, dir string) Fonts {
	content, err := fs.ReadFile(fSys, path.Join(dir, fontFile))
	if err != nil {
		log.Fatal("error loading font file: ", err)
	}
	tt, err := opentype.Parse(content)
	if err != nil {
		log.Fatal(err)
	}

	defaultFace, err := makeSize(tt, fontSize)
	if err != nil {
		log.Fatal(err)
	}

	return Fonts{
		Default: defaultFace,
	}
}

func makeSize(tt *sfnt.Font, size int) (font.Face, error) {
	const dpi = 72
	fontFace, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    float64(size),
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		return nil, err
	}
	fontFace = text.FaceWithLineHeight(fontFace, float64(size))
	return fontFace, nil
}
