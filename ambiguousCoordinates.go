package leetcode

import (
	"fmt"
	"strings"
)

/*
We had some 2-dimensional coordinates, like "(1, 3)" or "(2, 0.5)".  Then, we removed all commas, decimal points, and spaces, and ended up with the string S.  Return a list of strings representing all possibilities for what our original coordinates could have been.

Our original representation never had extraneous zeroes, so we never started with numbers like "00", "0.0", "0.00", "1.0", "001", "00.01", or any other number that can be represented with less digits.  Also, a decimal point within a number never occurs without at least one digit occuring before it, so we never started with numbers like ".1".

The final answer list can be returned in any order.  Also note that all coordinates in the final answer have exactly one space between them (occurring after the comma.)
*/

func ambiguousCoordinates(S string) []string {

	fmt.Println("input string: ", S)
	results := []string{}

	cleanString := S[1 : len(S)-1]

	for i := 1; i < len(cleanString); i++ {
		x1 := cleanString[:i]
		y1 := cleanString[i:]
		allX := permutations(x1)
		allY := permutations(y1)

		for _, x := range allX {
			for _, y := range allY {
				results = append(results, fmt.Sprintf("(%s, %s)", x, y))
			}
		}
	}
	return results
}

// permutations calculates all permutations of a given number
func permutations(num string) []string {
	results := []string{}

	// add the whole number
	if isValidNumber(num) {
		results = append(results, num)
	}

	for i := 0; i < len(num); i++ {
		n1 := num[:i]
		n2 := num[i:]
		y := fmt.Sprintf("%s.%s", n1, n2)
		if isValidNumber(y) {
			results = append(results, y)
		}
	}
	return results
}

// isValidNumber checks if both sides of the decimal is not zero
func isValidNumber(num string) bool {
	parts := strings.Split(num, ".")

	// does the number start with "0"?
	if len(parts) == 0 {
		return false
	}
	// is it too long?
	if len(parts) > 2 {
		return false
	}

	if len(parts) >= 1 {
		if len(parts[0]) == 0 {
			return false
		}
		// is the first part only "0"?
		if len(parts[0]) == 1 && parts[0] == "0" {
			return true
		}
		// does the first part end with "0"?
		if strings.HasSuffix(parts[0], "0") {
			return false
		}
	}

	// does the second part end with "0"?
	if len(parts) == 2 {
		if len(parts[1]) == 0 {
			return false
		} else if strings.HasSuffix(parts[1], "0") {
			return false
		}
	}
	return true
}
