package main

import (
	"fmt"
	"os"

	"github.com/ahmedsat/go-frame/frame"
)

func circle() {
	const (
		WIDTH  = 800
		HEIGHT = 600
	)

	img, f := CreateImage(WIDTH, HEIGHT)

	f.Fill(frame.WHITE)

	f.FillCircle(WIDTH/2, HEIGHT/2, 200, frame.GREEN)
	f.DrawCircle(WIDTH/2, HEIGHT/2, 250, frame.BLUE)
	f.SetPixel(WIDTH/2, HEIGHT/2, frame.RED)

	if err := SavePNG("img/circle.png", img); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
