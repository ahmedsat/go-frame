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
	f := frame.NewFrame(WIDTH, HEIGHT, func(x, y int64, c color.Color) {
		img.Set(int(x), int(y), c)
	})

	{
		f.Fill(frame.WHITE)
		f.DrawLine(0, 0, WIDTH, HEIGHT, frame.RED)
		f.DrawLine(WIDTH, 0, 0, HEIGHT, frame.RED)

		f.DrawLine(0, 0, WIDTH/4, HEIGHT, frame.GREEN)
		f.DrawLine(0, 0, WIDTH, HEIGHT/4, frame.GREEN)

		f.DrawLine(WIDTH/2, 0, WIDTH/2, HEIGHT, frame.BLUE)
		f.DrawLine(0, HEIGHT/2, WIDTH, HEIGHT/2, frame.BLUE)

	}
	// save image
	file, err := os.Create("img/out.png")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	png.Encode(file, img)

}
