package main

import (
	"fmt"
)

func main() {
	fmt.Println(combat(1.0, 100.0))
}

func combat(health, damage float64) float64 {
	life := health - damage
	if life <= 0 {
		return 0
	}
	return life

}
