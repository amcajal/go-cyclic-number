package main

import (
	"fmt"
	"github.com/amcajal/go-cyclic-number/cyclic"
)

func main() {
	matches := 0
	for i := 0; matches < 3; i++ { // Ugly loop, but feeling nasty
		if cyclic.IsCyclic(i, 2) {
			fmt.Printf("Cyclic number found: %v\n", i)
			matches++
		}
	}

}
