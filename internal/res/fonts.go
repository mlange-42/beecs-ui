package res

import (
	"io/fs"
	"log"
	"path"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/font/sfnt"
)

const fontFile = "mplus-1m-regular.ttf"

const fontSize = 18

// Fonts resource for access to UI fonts.
type Fonts struct {
	Default text.Face
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

func makeSize(tt *sfnt.Font, size int) (text.Face, error) {
	const dpi = 72
	fontFace, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    float64(size),
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		return nil, err
	}
	goFont := text.NewGoXFace(fontFace)
	//goFont = text.FaceWithLineHeight(goFont, float64(size))
	return goFont, nil
}
