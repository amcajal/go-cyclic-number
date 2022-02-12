package cyclic

import (
	"strconv"
	"strings"
)

// Check if two numbers have the same digits,
// no matter their position (i.e: 123 and 312 are "equal")
func CheckPermutation(leftN, rightN int) (result bool) {
	if leftN == rightN {
		return true
	}
	textLeft := strconv.FormatInt(int64(leftN), 10)
	textRight := strconv.FormatInt(int64(rightN), 10)

	if len(textLeft) != len(textRight) {
		return false
	}

	for {
		var lsd string
		// Less significant digit, NOT LDS as a drug (unless...)
		if leftN == 0 {
			lsd = "0"
		} else {
			lsd = strconv.FormatInt((int64(leftN) % 10), 10)
		}

		// Improvement_1: if lsd is in the list of already processed digits, ignore it

		leftCount := strings.Count(textLeft, lsd)
		rightCount := strings.Count(textRight, lsd)

		if rightCount != leftCount {
			return false
		}

		leftN = leftN / 10

		if leftN == 0 {
			break
		}
		// Improvement_1: lsd could be added to a list here
	}
	return true
}

// Recursive function: check permutations for multipliers from 2 to 6
func IsCyclic(number, factor int) bool {
	var result bool
	multiplication := number * factor
	if CheckPermutation(number, multiplication) {
		if factor < 6 {
			result = IsCyclic(number, factor+1)
		} else {
			result = true
		}
	} else {
		result = false
	}

	return result
}
