package canvas

import "github.com/fogleman/gg"

type RectangleFrame struct {
	RGBA
	width  float64
	height float64
	radius float64
}

func NewRectangleFrame(w float64, h float64) *RectangleFrame {
	return &RectangleFrame{
		width:  w,
		height: h,
	}
}

func (t *RectangleFrame) SetRadius(r float64) *RectangleFrame {
	t.radius = r
	return t
}

func (t *RectangleFrame) SetRGBA(r, g, b, a float64) *RectangleFrame {
	t.RGBA.SetRGBA(r, g, b, a)
	return t
}

func (t *RectangleFrame) Draw(dc *gg.Context, x, y float64) error {
	dc.Push()
	defer dc.Pop()
	t.applyRGBA(dc)
	if t.radius > 0 {
		dc.DrawRoundedRectangle(x, y, t.width, t.height, t.radius)
	} else {
		dc.DrawRectangle(x, y, t.width, t.height)
	}

	dc.Fill()
	return nil
}
