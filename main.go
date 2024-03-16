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
	WIDTH  = 800
	HEIGHT = 600
)

func main() {

	img := image.NewRGBA(image.Rect(0, 0, WIDTH, HEIGHT))
	f := frame.NewFrame(
		WIDTH, HEIGHT,
		func(x, y int, c color.Color) {
			img.Set(x, y, c)
		},
		func(x, y int) color.Color {
			return img.At(x, y)
		},
	)

	{

		f.Fill(frame.WHITE)

		f.SetPixel(WIDTH/2, HEIGHT/2, frame.RED)
	}

	// fmt.Println(img)

	//save image
	file, err := os.Create("img/out.png")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	png.Encode(file, img)

}

// todo: Linear interpolation
