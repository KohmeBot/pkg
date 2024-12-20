package canvas

import "github.com/fogleman/gg"

type Fonts struct {
	path string
}

func NewFonts(path string) *Fonts {
	return &Fonts{path: path}
}

func (f *Fonts) Load(dc *gg.Context, points float64) error {
	return dc.LoadFontFace(f.path, points)
}
