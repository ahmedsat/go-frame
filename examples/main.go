package main

import (
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/ahmedsat/go-frame/frame"
)

func CreateImage(width, height int) (*image.RGBA, frame.Frame) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	f := frame.NewFrame(width, height, func(x, y int, c color.Color) {
		img.Set(x, y, c)
	})

	return img, f
}

func SavePNG(path string, img image.Image) (err error) {
	file, err := os.Create(path)
	if err != nil {
		return
	}
	png.Encode(file, img)
	return
}

func main() {
	line()
	rect()
	circle()
}
