package canvas

import (
	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"image"
)

type Background struct {
	RGBA
	img    image.Image
	radius float64
}

// NewImgBackground 新建图像背景
func NewImgBackground(img image.Image) *Background {
	return NewImgBackgroundWithBlur(img, 0)
}

// NewImgBackgroundWithBlur 新建带模糊的图像背景
func NewImgBackgroundWithBlur(img image.Image, radius float64) *Background {
	return &Background{img: img, radius: radius}
}

// NewColorBackground 纯色背景
func NewColorBackground(r, g, b, a float64) *Background {
	return &Background{RGBA: RGBA{r, g, b, a}}
}

func (b *Background) Draw(dc *gg.Context, x, y float64) error {
	dc.Push()
	defer dc.Pop()
	if b.img == nil {
		b.applyRGBA(dc)
		dc.Clear()
		return nil
	}
	w, h := dc.Width(), dc.Height()

	img := imaging.Blur(b.img, b.radius)
	img = imaging.Resize(img, w, h, imaging.Lanczos)

	dc.DrawImage(img, 0, 0)
	return nil
}
