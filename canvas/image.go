package canvas

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"image"
	"math"
	"net/http"
)

type ImageFactory struct {
	resource any
}
type (
	filePath string
	url      string
	b64      string
)

func (f *ImageFactory) Url(u string) {
	f.resource = url(u)
}
func (f *ImageFactory) File(path string) {
	f.resource = filePath(path)
}
func (f *ImageFactory) Base64(b64Str string) {
	f.resource = b64(b64Str)
}
func (f *ImageFactory) ByteData(data []byte) {
	f.resource = data
}
func (f *ImageFactory) Get() (image.Image, error) {
	switch r := f.resource.(type) {
	case filePath:
		return imaging.Open(string(r))
	case url:
		return FetchImage(string(r))
	case b64:
		return Base64ToImage(string(r))
	case []byte:
		return BytesToImage(r)
	default:
		return nil, fmt.Errorf("unknown resource type %T", f.resource)
	}
}

// FetchImage 从给定的URL下载图像，并返回 image.Image 对象
func FetchImage(url string) (image.Image, error) {
	// 发送 GET 请求
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch image: %v", err)
	}
	defer resp.Body.Close()

	// 检查 HTTP 状态码
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch image: status code %d", resp.StatusCode)
	}

	// 读取并解码图像数据
	img, _, err := image.Decode(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to decode image: %v", err)
	}
	return img, nil
}
func Base64ToImage(base64Str string) (image.Image, error) {
	// 解码 Base64 字符串为字节数据
	data, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		return nil, fmt.Errorf("failed to decode b64 string: %w", err)
	}
	return BytesToImage(data)
}
func BytesToImage(data []byte) (image.Image, error) {
	// 从字节数据解码为图像
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("failed to decode image: %w", err)
	}
	return img, nil
}

type ImageCircleFrame struct {
	img image.Image
	RGBA
	radius      float64 // 半径
	borderWidth float64 // 边框宽度
}

// NewImageCircleFrame 圆形图片
func NewImageCircleFrame(img image.Image, radius float64) *ImageCircleFrame {
	return &ImageCircleFrame{
		img:    img,
		radius: radius,
	}
}

// SetBorder 设置边框
func (f *ImageCircleFrame) SetBorder(r, g, b, a float64, width float64) *ImageCircleFrame {
	f.RGBA.SetRGBA(r, g, b, a)
	f.borderWidth = width
	return f
}

func (f *ImageCircleFrame) Draw(dc *gg.Context, x, y float64) error {
	dc.Push()
	defer dc.Pop()
	if f.borderWidth > 0 {
		f.applyRGBA(dc)
		dc.DrawCircle(x, y, f.radius+f.borderWidth)
		dc.Fill()
	}
	img := imaging.Resize(f.img, int(2*f.radius), int(2*f.radius), imaging.Lanczos)

	dc.DrawCircle(x, y, f.radius)
	dc.Clip()
	dc.DrawImageAnchored(img, int(x), int(y), 0.5, 0.5)
	dc.ResetClip()
	return nil
}

type ImageRectangleFrame struct {
	img image.Image
	RGBA
	width       float64 // 宽度
	height      float64 // 高度
	radius      float64 // 圆滑半径
	borderWidth float64 // 边框宽度
}

func NewImageRectangleFrame(img image.Image, width, height float64) *ImageRectangleFrame {
	return &ImageRectangleFrame{
		img:    img,
		width:  width,
		height: height,
	}
}
func (f *ImageRectangleFrame) SetBorder(r, g, b, a float64, width float64) *ImageRectangleFrame {
	f.RGBA.SetRGBA(r, g, b, a)
	f.borderWidth = width
	return f
}
func (f *ImageRectangleFrame) SetRadius(r float64) *ImageRectangleFrame {
	f.radius = r
	return f
}
func (f *ImageRectangleFrame) Draw(dc *gg.Context, x, y float64) error {
	dc.Push()
	defer dc.Pop()
	if f.borderWidth > 0 {
		f.applyRGBA(dc)
		dc.DrawRoundedRectangle(x-f.borderWidth, y-f.borderWidth, f.width+(2*f.borderWidth), f.height+(2*f.borderWidth), f.radius+f.borderWidth)
		dc.Fill()
	}
	img := imaging.Resize(f.img, int(f.width), int(f.height), imaging.Lanczos)
	dc.DrawRoundedRectangle(x, y, f.width, f.height, f.radius)
	dc.Clip()
	dc.DrawImageAnchored(img, int(x), int(y), 0, 0)
	dc.ResetClip()
	return nil
}

type ImageGridFrame struct {
	width, height float64 // 限定宽高
	spacing       float64 // 间距
	images        []image.Image
}

type imgParma struct {
	img  image.Image
	w, h float64 // 图片宽高
	x, y float64 // x y

}

// NewImageGrid 创建一个图片网格
func NewImageGrid(width, height float64, images ...image.Image) *ImageGridFrame {
	return &ImageGridFrame{
		width:  width,
		height: height,
		images: images,
	}
}

// SetSpacing 设置间距
func (f *ImageGridFrame) SetSpacing(spacing float64) *ImageGridFrame {
	f.spacing = spacing
	return f
}

func (f *ImageGridFrame) Draw(dc *gg.Context, x, y float64) error {
	dc.Push()
	defer dc.Pop()
	ps := f.measure()
	for _, p := range ps {
		r := NewImageRectangleFrame(p.img, p.w, p.h)
		if err := r.Draw(dc, x+p.x, y+p.y); err != nil {
			return err
		}
	}
	return nil
}

// 计算出每个图片的开始位置以及宽高
func (f *ImageGridFrame) measure() []imgParma {
	num := len(f.images) // 图片个数
	if num == 0 {
		return nil
	}

	// 假设网格是一个近似的正方形布局
	cols := int(math.Ceil(math.Sqrt(float64(num))))      // 列数
	rows := int(math.Ceil(float64(num) / float64(cols))) // 行数

	// 计算单元格的宽度和高度（包含间隔）
	cellWidth := (f.width - float64(cols-1)*f.spacing) / float64(cols)
	cellHeight := (f.height - float64(rows-1)*f.spacing) / float64(rows)

	var params []imgParma
	for i, img := range f.images {
		col := i % cols // 当前图片所在列
		row := i / cols // 当前图片所在行

		// 计算图片的坐标，加入间隔调整
		x := float64(col) * (cellWidth + f.spacing)
		y := float64(row) * (cellHeight + f.spacing)

		// 调整图片宽高，使其适应单元格
		imgWidth := cellWidth
		imgHeight := cellHeight
		imgBounds := img.Bounds()
		originalWidth := float64(imgBounds.Dx())
		originalHeight := float64(imgBounds.Dy())

		// 保持图片宽高比
		imgAspect := originalWidth / originalHeight
		cellAspect := cellWidth / cellHeight
		if imgAspect > cellAspect {
			// 图片更宽，按宽度缩放
			imgHeight = cellWidth / imgAspect
		} else {
			// 图片更高，按高度缩放
			imgWidth = cellHeight * imgAspect
		}

		params = append(params, imgParma{
			img: img,
			w:   imgWidth,
			h:   imgHeight,
			x:   x + (cellWidth-imgWidth)/2,   // 居中调整
			y:   y + (cellHeight-imgHeight)/2, // 居中调整
		})
	}

	return params
}
