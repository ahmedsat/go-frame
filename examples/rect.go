package main

import (
	"fmt"
	"os"

	"github.com/ahmedsat/go-frame/frame"
)

func rect() {

	const (
		WIDTH  = 800
		HEIGHT = 600
	)

	img, f := CreateImage(WIDTH, HEIGHT)

	f.Fill(frame.WHITE)

	f.FillRectangle(0, 0, WIDTH*3/4, HEIGHT*3/4, frame.GREEN)
	f.FillRectangle(WIDTH, HEIGHT, WIDTH/4, HEIGHT/4, frame.BLUE)
	f.DrawRectangle(WIDTH/3, HEIGHT/3, WIDTH*2/3, HEIGHT*2/3, frame.RED)

	if err := SavePNG("img/rect.png", img); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
