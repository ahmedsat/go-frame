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

func NewFrame(width, height int, set setPixel, get getPixel) (f Frame) {
	f.height = height
	f.width = width
	f.SetPixel = set
	f.GetPixel = get

	f.BlendPixel = func(x, y int, c1 color.Color) {

		c0_r, c0_g, c0_b, c0_a := f.GetPixel(x, y).RGBA()
		c1_r, c1_g, c1_b, c1_a := c1.RGBA()

		if c0_a == 0 || c1_a == 0xffff {
			f.SetPixel(x, y, c1)
			return
		}

		r := Int.Lerp(int(c0_r), int(c1_r), int(c1_a), 0xffff)
		g := Int.Lerp(int(c0_g), int(c1_g), int(c1_a), 0xffff)
		b := Int.Lerp(int(c0_b), int(c1_b), int(c1_a), 0xffff)

		c := color.RGBA{
			R: uint8(r),
			G: uint8(g),
			B: uint8(b),
			A: 0xff, //uint8(a),
		}

		f.SetPixel(x, y, c)
	}

	return f
}

func (f *Frame) Resize(width, height int) {
	f.height = height
	f.width = width
}

func (f *Frame) UpdateSetPixel(set setPixel) {
	f.SetPixel = set
}

func (f *Frame) UpdateGetPixel(get getPixel) {
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
		// y = mx + c
		dy := (dx * (y1 - y0) / (x1 - x0)) + (y0 - x0*(y1-y0)/(x1-x0))

		f.BlendPixel(dx, dy, color)

	}
}

func (f Frame) drawLineByY(x0, y0, x1, y1 int, color color.Color) {
	maxY := max(y0, y1)
	minY := min(y0, y1)
	for dy := minY; dy <= maxY; dy++ {
		// y 	= mx + c
		// mx	= y - c
		// x = (y - c)/m

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
			// f.SetPixel(x, y, color)
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
func (f Frame) FillTriangle(x1, y1, x2, y2, x3, y3 int, color color.Color) {
	maxX := max(x1, x2, x3)
	minX := min(x1, x2, x3)
	maxY := max(y1, y2, y3)
	minY := min(y1, y2, y3)

	v1 := Int.Vec2{x1, y1}
	v2 := Int.Vec2{x2, y2}
	v3 := Int.Vec2{x3, y3}

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			p := Int.Vec2{x, y}
			t1 := p.RightTo(v1, v2)
			t2 := p.RightTo(v2, v3)
			t3 := p.RightTo(v3, v1)
			if t1 && t2 && t3 {
				f.BlendPixel(x, y, color)
			}
		}
	}

}
