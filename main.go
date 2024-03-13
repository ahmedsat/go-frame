package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/ahmedsat/go-frame/frame"
)

const (
	WIDTH  = 80
	HEIGHT = 60
)

func main() {

	img := image.NewRGBA(image.Rect(0, 0, WIDTH, HEIGHT))
	f := frame.NewFrame(WIDTH, HEIGHT, func(x, y int, c color.Color) {
		img.Set(x, y, c)
	})

	{
		f.Fill(frame.WHITE)

		f.SetPixel(10, 10, frame.RED)
		f.SetPixel(10, 50, frame.RED)
		f.SetPixel(50, 10, frame.RED)
		f.SetPixel(50, 50, frame.RED)
		// f.FillRectangle(10, 10, 50, 50, frame.GREEN)
		f.FillRectangle(50, 50, 10, 10, frame.GREEN)
		// f.DrawRectangle(10, 10, 50, 50, frame.BLUE)
	}
	// save image
	file, err := os.Create("img/out.png")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	png.Encode(file, img)

}

// todo: drawCircle
// todo: fillCircle
