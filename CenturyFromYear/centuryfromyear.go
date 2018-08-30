package main

import (
	"fmt"
	"math"
)

func main() {
	a := century(2000)
	fmt.Println(a)
}

func century(year int) int {
	var c int
	c = year / 100
	if math.Mod(float64(year), 100) != 0 {
		c = c + 1
	}
	return c
}
