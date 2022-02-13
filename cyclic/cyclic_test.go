package cyclic

import (
	"testing"
)

// Checks that permutations are correctly detected
func TestPermutationChecking(t *testing.T) {
	var r bool

	// False permutations
	r = checkPermutation(0, 10)
	if r == true {
		t.Fatalf("CheckPermutation of 0 and 10 shall return false")
	}

	r = checkPermutation(10, 0)
	if r == true {
		t.Fatalf("CheckPermutation of 10 and 0 shall return false")
	}

	r = checkPermutation(0, 1)
	if r == true {
		t.Fatalf("CheckPermutation of 0 and 1 shall return false")
	}

	r = checkPermutation(1, 0)
	if r == true {
		t.Fatalf("CheckPermutation of 1 and 0 shall return false")
	}

	r = checkPermutation(12, 13)
	if r == true {
		t.Fatalf("CheckPermutation of 12 and 13 shall return false")
	}

	// True permutations
	r = checkPermutation(0, 0)
	if r == false {
		t.Fatalf("CheckPermutation of 0 and 0 shall return true")
	}

	r = checkPermutation(12, 21)
	if r == false {
		t.Fatalf("CheckPermutation of 12 and 13 shall return true")
	}

	r = checkPermutation(1234, 4231)
	if r == false {
		t.Fatalf("CheckPermutation of 1234 and 4231 shall return true")
	}

}

func TestCyclicNumbers(t *testing.T) {
	// Zero is a special case, but still meeting the requirements
	if r := IsCyclic(0, 2); !r {
		t.Fatalf("Zero is cyclic")
	}

	// Here, a more aggresive test could be done, testing a lot of numbers
	// and breaking if any of them but the correct one reports a true
	if r := IsCyclic(142857, 2); !r {
		t.Fatalf("142857 is cyclic")
	}

	if r := IsCyclic(2245, 2); r {
		t.Fatalf("2245 is NOT cyclic")
	}
}

func TestResourceUsage(t *testing.T) {
	matches := 0
	for i := 0; matches < 3; i++ { // Ugly loop, but feeling nasty
		if IsCyclic(i, 2) {
			matches++
		}
	}
}
