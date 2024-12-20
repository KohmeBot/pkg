package canvas

import (
	"bytes"
	"encoding/base64"
	"github.com/fogleman/gg"
	"image"
	"image/jpeg"
)

type Canvas struct {
	w  int // 宽
	h  int // 高
	dc *gg.Context
}

type Element interface {
	Draw(dc *gg.Context, x, y float64) error
}

func NewCanvas(w, h int) *Canvas {
	return &Canvas{
		w:  w,
		h:  h,
		dc: gg.NewContext(w, h),
	}
}

func (c *Canvas) SetBackground(element Element) error {
	return c.DrawWith(element, 0, 0)
}

func (c *Canvas) DrawWith(element Element, x, y float64) error {
	return element.Draw(c.dc, x, y)
}

func (c *Canvas) ToBytes() ([]byte, error) {
	var buf bytes.Buffer
	img := c.ToImage()
	if err := jpeg.Encode(&buf, img, nil); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (c *Canvas) Save(path string) error {
	return c.dc.SavePNG(path)
}
func (c *Canvas) ToImage() image.Image {
	return c.dc.Image()
}

func (c *Canvas) ToBase64() (string, error) {
	data, err := c.ToBytes()
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(data), nil
}
