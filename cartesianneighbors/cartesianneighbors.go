package main

import (
	"fmt"
)

func main() {
	fmt.Println("\n", CartesianNeighbor(2, 2))
}

//CartesianNeighbor returns neighbors    func CartesianNeighbor(x,y int) [][]int
func CartesianNeighbor(x, y int) [][]int {
	var r [][]int
	var n []int
	xoptions := [3]int{x - 1, x, x + 1}
	yoptions := [3]int{y - 1, y, y + 1}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if xoptions[i] == x && yoptions[j] == y {
				continue
			}
			n = append(n, xoptions[i], yoptions[j])
			r = append(r, n)
			n = nil
		}
	}
	return r
}
