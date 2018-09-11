package main

import (
	"fmt"
)

func main() {
	fmt.Println("unique? ", HasUniqueChar("+-"))
}

//HasUniqueChar identifies if a string has repeated chars
func HasUniqueChar(str string) bool {
	var base, compare string
	for i, v := range str {
		for _, sv := range str[i+1 : len(str)] {
			base = string(v)
			compare = string(sv)
			if base == compare {
				return false
			}
		}
	}
	return true
}
