package frame

import (
	"image/color"
)

type drawPixel func(x, y int64, color color.Color)

type frame struct {
	width    int64
	height   int64
	SetPixel drawPixel
}

func NewFrame(width, height int64, draw drawPixel) (f frame) {
	f.height = height
	f.width = width
	f.SetPixel = draw
	return f
}

func (f *frame) Resize(width, height int64) {
	f.height = height
	f.width = width
}

func (f *frame) UpdateDrawer(draw drawPixel) {
	f.SetPixel = draw
}

func (f frame) Fill(color color.Color) {
	for y := int64(0); y < f.height; y++ {
		for x := int64(0); x < f.width; x++ {
			f.SetPixel(x, y, color)
		}
	}
}

func (f frame) DrawLine(x0, y0, x1, y1 int64, color color.Color) {

	f.drawLineByX(x0, y0, x1, y1, color)
	f.drawLineByY(x0, y0, x1, y1, color)

}

func (f frame) drawLineByX(x0, y0, x1, y1 int64, color color.Color) {
	maxX := max(x0, x1)
	minX := min(x0, x1)

	for dx := minX; dx < maxX-minX; dx++ {
		// y = mx + c
		dy := (dx * (y1 - y0) / (x1 - x0)) + (y0 - x0*(y1-y0)/(x1-x0))
		f.SetPixel(dx, dy, color)
	}

}
func (f frame) drawLineByY(x0, y0, x1, y1 int64, color color.Color) {
	maxY := max(y0, y1)
	minY := min(y0, y1)
	for dy := minY; dy < maxY-minY; dy++ {
		// y 	= mx + c
		// mx	= y - c
		// x = (y - c)/m

		var dx int64
		if x0 == x1 {
			dx = x0
		} else {
			dx = (x1 - x0) * (dy - (y0 - x0*(y1-y0)/(x1-x0))) / (y1 - y0)
		}
		f.SetPixel(dx, dy, color)
	}

}
