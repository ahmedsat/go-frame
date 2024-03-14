package main

import (
	"fmt"
	"os"

	"github.com/ahmedsat/go-frame/frame"
)

func triangle() {

	const (
		WIDTH  = 800
		HEIGHT = 600
	)

	img, f := CreateImage(WIDTH, HEIGHT)

	f.Fill(frame.WHITE)

	f.DrawTriangle(WIDTH/2, HEIGHT/5, WIDTH*4/5, HEIGHT*4/5, WIDTH/5, HEIGHT*4/5, frame.GREEN)
	f.FillTriangle(WIDTH/2, HEIGHT/4, WIDTH*3/4, HEIGHT*3/4, WIDTH/4, HEIGHT*3/4, frame.BLUE)
	f.SetPixel(WIDTH/2, HEIGHT/2, frame.RED)

	if err := SavePNG("img/triangle.png", img); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
