package main

import (
	"fmt"
)

func main() {
	fmt.Println(InAscOrder([]int{1}))
}

//InAscOrder checking order of values
func InAscOrder(list []int) bool {
	for i := 1; i < len(list); i++ {
		if list[i-1] > list[i] {
			return false
		}
	}
	return true
}
