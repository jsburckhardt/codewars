package main

import (
	"fmt"
)

func main() {
	fmt.Println(RemoveChar("eloquent"))
}

//RemoveChar multiplies by -1
func RemoveChar(word string) string {
	return word[1 : len(word)-1]
}
