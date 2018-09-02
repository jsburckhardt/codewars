package main

import (
	"fmt"
)

func main() {
	fmt.Println(MakeNegative(12))
}

//MakeNegative makes a number negative
func MakeNegative(x int) int {
	if x > 0 {
		return x * -1
	}
	return x
}
