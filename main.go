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
	f := frame.NewFrame(WIDTH, HEIGHT, func(x, y int, c color.Color) {
		img.Set(x, y, c)
	})

	{
		f.Fill(frame.WHITE)

		f.FillCircle(WIDTH/2, HEIGHT/2, 200, frame.GREEN)
		f.DrawCircle(WIDTH/2, HEIGHT/2, 250, frame.BLUE)
		f.SetPixel(WIDTH/2, HEIGHT/2, frame.RED)

	}
	// save image
	file, err := os.Create("img/out.png")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	png.Encode(file, img)

}
