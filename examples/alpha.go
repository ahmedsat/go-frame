package main

import (
	"fmt"
	"image/color"
	"os"

	"github.com/ahmedsat/go-frame/frame"
)

func alpha() {

	const (
		WIDTH  = 800
		HEIGHT = 600
	)

	img, f := CreateImage(WIDTH, HEIGHT)

	f.Fill(frame.WHITE)

	f.FillRectangle(0, 0, WIDTH*3/4, HEIGHT*3/4, frame.GREEN)
	f.FillRectangle(WIDTH/4, HEIGHT/4, WIDTH, HEIGHT, color.RGBA{
		R: 0x00,
		G: 0x00,
		B: 0xff,
		A: 0xbb,
	})

	f.FillTriangle(WIDTH/2, HEIGHT/5, WIDTH*4/5, HEIGHT*4/5, WIDTH/5, HEIGHT*4/5, color.RGBA{
		R: 0xff,
		G: 0x00,
		B: 0x00,
		A: 0xcc,
	})

	f.FillCircle(WIDTH/2-100, HEIGHT/2+100, 200, color.RGBA{
		R: 0x00,
		G: 0xff,
		B: 0x00,
		A: 0xaa,
	})

	if err := SavePNG("img/alpha.png", img); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
