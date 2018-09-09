package main

import (
	"fmt"
	"image/color"
)

func main() {
	fmt.Println(DrawLines([][]int{{1, 2}, {1, 2}}))
}

func DrawLines(points [][]int) string {
	canvas := NewCanvas(100, 100)
	canvas.Clear(color.RGBA{255, 255, 255, 255})
	canvas.PenColor = color.RGBA{255, 0, 0, 255}

	return canvas.EncodeString()
}
