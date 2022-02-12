package main

import (
	"fmt"
	"github.com/amcajal/go-cyclic-number/cyclic"
)

func main() {
	// TESTS
	fmt.Printf("Comparisson between %v and %v is %v\n", 0, 10, cyclic.CheckPermutation(0, 10))
	fmt.Printf("Comparisson between %v and %v is %v\n", 0, 0, cyclic.CheckPermutation(0, 0))
	fmt.Printf("Comparisson between %v and %v is %v\n", 0, 1, cyclic.CheckPermutation(0, 1))
	fmt.Printf("Comparisson between %v and %v is %v\n", 1, 0, cyclic.CheckPermutation(1, 0))

	fmt.Printf("Comparisson between %v and %v is %v\n", 12, 13, cyclic.CheckPermutation(12, 13))
	fmt.Printf("Comparisson between %v and %v is %v\n", 12, 21, cyclic.CheckPermutation(12, 21))
	fmt.Printf("Comparisson between %v and %v is %v\n", 1234, 4231, cyclic.CheckPermutation(1234, 4231))

	matches := 0
	for i := 0; matches < 3; i++ { // Ugly loop, but feeling nasty
		if cyclic.IsCyclic(i, 2) {
			fmt.Printf("Cyclic number found: %v\n", i)
			matches++
		}
	}

}
