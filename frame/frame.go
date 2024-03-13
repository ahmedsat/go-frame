package frame

import (
	"image/color"
)

type drawPixel func(x, y int, color color.Color)

type Frame struct {
	width    int
	height   int
	SetPixel drawPixel
}

func NewFrame(width, height int, draw drawPixel) (f Frame) {
	f.height = height
	f.width = width
	f.SetPixel = draw
	return f
}

func (f *Frame) Resize(width, height int) {
	f.height = height
	f.width = width
}

func (f *Frame) UpdateDrawer(draw drawPixel) {
	f.SetPixel = draw
}

func (f Frame) Fill(color color.Color) {
	for y := int(0); y < f.height; y++ {
		for x := int(0); x < f.width; x++ {
			f.SetPixel(x, y, color)
		}
	}
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

		f.SetPixel(dx, dy, color)

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
		f.SetPixel(dx, dy, color)
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
			f.SetPixel(x, y, color)
		}
	}
}
