package main

import (
	"fmt"
)

func main() {
	base := []int{1, 2, 3, 4, 5}
	fmt.Println(Maps(base))
}

//Maps do something
func Maps(x []int) []int {
	var r []int
	for _, value := range x {
		r = append(r, value*2)
	}
	return r
}
