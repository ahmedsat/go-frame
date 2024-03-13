package main

import (
	"fmt"
	"os"

	"github.com/ahmedsat/go-frame/frame"
)

func main() {

	const (
		WIDTH  = 800
		HEIGHT = 600
	)

	img, f := CreateImage(WIDTH, HEIGHT)

	f.Fill(frame.WHITE)
	f.DrawLine(0, 0, WIDTH, HEIGHT, frame.RED)
	f.DrawLine(WIDTH, 0, 0, HEIGHT, frame.RED)

	f.DrawLine(0, 0, WIDTH/4, HEIGHT, frame.GREEN)
	f.DrawLine(0, 0, WIDTH, HEIGHT/4, frame.GREEN)

	f.DrawLine(WIDTH/2, 0, WIDTH/2, HEIGHT, frame.BLUE)
	f.DrawLine(0, HEIGHT/2, WIDTH, HEIGHT/2, frame.BLUE)

	if err := SavePNG("img/line.png", img); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
