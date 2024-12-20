package canvas

import (
	"github.com/fogleman/gg"
	"strings"
)

type TextFrame struct {
	Fonts
	RGBA
	frame       *RectangleFrame
	content     string  // 内容
	points      int     // 字号
	lineSpacing float64 // 行间隔
	width       float64 // 限定宽度

	ax float64 // 距左右边距
	ay float64 // 距上下边距

	align gg.Align // 对齐方式
}

func NewTextFrame(c string, fonts *Fonts, points int, lineSpacing float64, width float64, align gg.Align) *TextFrame {
	return &TextFrame{
		Fonts:       *fonts,
		content:     c,
		points:      points,
		lineSpacing: lineSpacing,
		width:       width,
		align:       align,
		ax:          5,
		ay:          5,
	}
}

func (t *TextFrame) SetRGBA(r, g, b, a float64) *TextFrame {
	t.r, t.g, t.b, t.a = r, g, b, a
	return t
}
func (t *TextFrame) SetAx(x float64) *TextFrame {
	t.ax = x
	return t
}
func (t *TextFrame) SetAy(y float64) *TextFrame {
	t.ay = y
	return t
}
func (t *TextFrame) SetFrame(r, g, b, a, radius float64) *TextFrame {
	t.frame = &RectangleFrame{
		RGBA: RGBA{
			r: r,
			g: g,
			b: b,
			a: a,
		},
		radius: radius,
	}
	return t
}

func (t *TextFrame) Draw(dc *gg.Context, x, y float64) error {
	dc.Push()
	defer dc.Pop()
	if err := t.Load(dc, float64(t.points)); err != nil {
		return err
	}

	// 获得调整之后的行
	lines := t.adjust(dc, t.width, dc.WordWrap(t.content, t.width))
	content := strings.Join(lines, "\n")
	t.applyRGBA(dc)
	// 获取渲染宽高
	_, h := dc.MeasureMultilineString(content, t.lineSpacing)

	if t.frame != nil {
		t.frame.height = h + 2*(t.ax)
		t.frame.width = t.width + 2*(t.ay)
		if err := t.frame.Draw(dc, x, y); err != nil {
			return err
		}
	}
	dc.DrawStringWrapped(content, x+t.ax, y+t.ay, 0, 0, t.width, t.lineSpacing, t.align)

	return nil
}

// 调整行，符合宽度
func (t *TextFrame) adjust(dc *gg.Context, w float64, lines []string) []string {
	var res []string
	for _, line := range lines {
		sw, _ := dc.MeasureString(line)
		// 小于等于限定宽度
		if sw <= w {
			res = append(res, line)
			continue
		}
		// 调整该行
		var newLine strings.Builder
		for _, r := range line {
			// 逐个增加字符
			newLine.WriteRune(r)
			sw, _ := dc.MeasureString(newLine.String())
			if sw >= w {
				// 当前宽度超出，换行
				res = append(res, newLine.String())
				newLine.Reset()
			}
		}
		if newLine.Len() > 0 {
			res = append(res, newLine.String())
		}

	}
	return res
}
