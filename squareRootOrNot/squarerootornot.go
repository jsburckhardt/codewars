package main

import (
	"fmt"
	"math"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("dcl:", a)

	r := SquareOrSquareRoot(a)
	fmt.Println("dcl:", r)
}

//SquareOrSquareRoot sucks
func SquareOrSquareRoot(a []int) []int {
	r := make([]int, 0)
	for _, value := range a {
		if whole, decimal := math.Modf(math.Sqrt(float64(value))); decimal == 0 {
			r = append(r, int(whole))
		} else {
			r = append(r, int(math.Pow(float64(value), 2)))
		}
	}
	return r
}
