package main

import (
	"fmt"
)

func main() {
	staircase(5)
}

func staircase(x int) {
	var line []string
	step := "#"
	// Pad the string and store it in a new string.
	for y := x; y > 0; y-- {
		for i := 0; i < y; i++ {
			line = append(line, " ")
		}
		r := append(line, step)
		fmt.Println(r)
		line = nil
	}
}
