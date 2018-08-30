package main

import (
	"fmt"
)

func main() {
	fmt.Println(Opposite(14))
}

//Opposite multiplies by -1
func Opposite(value int) int {
	return value * -1
}
