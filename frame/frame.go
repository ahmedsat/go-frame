package frame

import (
	"image/color"

	Int "github.com/ahmedsat/go-frame/math/int"
)

type setPixel func(x, y int, color color.Color)
type getPixel func(x, y int) color.Color
type blendPixel func(x, y int, color color.Color)

type Frame struct {
	width      int
	height     int
	SetPixel   setPixel
	GetPixel   getPixel
	BlendPixel blendPixel
}

func NewFrame(width, height int, set setPixel, get getPixel) Frame {
	return Frame{
		width:    width,
		height:   height,
		SetPixel: set,
		GetPixel: get,
		BlendPixel: func(x, y int, c color.Color) {
			srcR, srcG, srcB, srcA := c.RGBA()
			dstR, dstG, dstB, _ := get(x, y).RGBA()
			r := Int.Lerp(int(dstR), int(srcR), int(srcA), 0xffff)
			g := Int.Lerp(int(dstG), int(srcG), int(srcA), 0xffff)
			b := Int.Lerp(int(dstB), int(srcB), int(srcA), 0xffff)
			c = color.RGBA{
				R: uint8(r),
				G: uint8(g),
				B: uint8(b),
				A: 0xff,
			}
			set(x, y, c)
		},
	}
}

func (f *Frame) Resize(newWidth, newHeight int) {
	f.width, f.height = newWidth, newHeight
}

func (f *Frame) UpdateSetPixelMethod(set setPixel) {
	f.SetPixel = set
}

func (f *Frame) UpdateGetPixelMethod(get getPixel) {
	f.GetPixel = get
}

func (f Frame) Fill(color color.Color) {
	f.FillRectangle(0, 0, f.width, f.height, color)
}

func (f Frame) DrawLine(x0, y0, x1, y1 int, color color.Color) {
	if x0 != x1 {
		f.drawLineByX(x0, y0, x1, y1, color)
	}
	if y0 != y1 {
		f.drawLineByY(x0, y0, x1, y1, color)
	}

}

func (f Frame) drawLineByX(x0, y0, x1, y1 int, color color.Color) {
	maxX := max(x0, x1)
	minX := min(x0, x1)

	for dx := minX; dx <= maxX; dx++ {

		dy := (dx * (y1 - y0) / (x1 - x0)) + (y0 - x0*(y1-y0)/(x1-x0))

		f.BlendPixel(dx, dy, color)

	}
}

func (f Frame) drawLineByY(x0, y0, x1, y1 int, color color.Color) {
	maxY := max(y0, y1)
	minY := min(y0, y1)
	for dy := minY; dy <= maxY; dy++ {

		var dx int
		if x0 == x1 {
			dx = x0
		} else {
			dx = (x1 - x0) * (dy - (y0 - x0*(y1-y0)/(x1-x0))) / (y1 - y0)
		}
		f.BlendPixel(dx, dy, color)
	}

}

func (f Frame) DrawRectangle(x0, y0, x1, y1 int, color color.Color) {

	f.DrawLine(x0, y0, x1, y0, color)
	f.DrawLine(x1, y0, x1, y1, color)
	f.DrawLine(x1, y1, x0, y1, color)
	f.DrawLine(x0, y1, x0, y0, color)

}

func (f Frame) FillRectangle(x0, y0, x1, y1 int, color color.Color) {
	maxX := max(x0, x1)
	minX := min(x0, x1)
	maxY := max(y0, y1)
	minY := min(y0, y1)
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			f.BlendPixel(x, y, color)
		}
	}
}

func (f Frame) FillCircle(x0, y0, r int, color color.Color) {
	maxX := x0 + r
	minX := x0 - r
	maxY := y0 + r
	minY := y0 - r
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			dx := x - x0
			dy := y - y0
			dSquare := dx*dx + dy*dy
			if dSquare <= r*r {
				f.BlendPixel(x, y, color)
			}
		}
	}
}

func (f Frame) DrawCircle(x0, y0, r int, color color.Color) {
	maxX := x0 + r
	minX := x0 - r
	maxY := y0 + r
	minY := y0 - r

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			dx := x - x0
			dy := y - y0
			dSquare := dx*dx + dy*dy
			d := (dSquare - r*r)
			if d <= r && d >= -r {
				f.BlendPixel(x, y, color)
			}
		}
	}
}

func (f Frame) DrawTriangle(x1, y1, x2, y2, x3, y3 int, color color.Color) {
	f.DrawLine(x1, y1, x2, y2, color)
	f.DrawLine(x1, y1, x3, y3, color)
	f.DrawLine(x3, y3, x2, y2, color)
}

func (f Frame) FillTriangle(ax, ay, bx, by, cx, cy int, color color.Color) {
	minX, maxX := min(ax, bx, cx), max(ax, bx, cx)
	minY, maxY := min(ay, by, cy), max(ay, by, cy)

	v1, v2, v3 := Int.Vec2{X: ax, Y: ay}, Int.Vec2{X: bx, Y: by}, Int.Vec2{X: cx, Y: cy}

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			p := Int.Vec2{X: x, Y: y}
			if p.RightTo(v1, v2) && p.RightTo(v2, v3) && p.RightTo(v3, v1) {
				f.BlendPixel(x, y, color)
			}
		}
	}
}
