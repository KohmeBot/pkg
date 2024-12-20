package canvas

import "github.com/fogleman/gg"

type RGBA struct {
	r, g, b, a float64
}

func (rgba *RGBA) SetRGBA(r, g, b, a float64) *RGBA {
	rgba.r, rgba.g, rgba.b, rgba.a = r, g, b, a
	return rgba
}

func (rgba *RGBA) applyRGBA(dc *gg.Context) {
	dc.SetRGBA(rgba.r, rgba.g, rgba.b, rgba.a)
}
