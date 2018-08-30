package main

import (
	"fmt"
)

func main() {
	fmt.Println(PositiveSum([]int{-1, -2}))
}

//PositiveSum adds just positives
func PositiveSum(numbers []int) int {
	var r int
	for _, value := range numbers {
		if value >= 0 {
			r = r + value
		}
	}
	return r
}
