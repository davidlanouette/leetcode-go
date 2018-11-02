package leetcode

import (
	"fmt"
)

func isPalindrome(x int) bool {
	fmt.Println("x=", x)
	i := x

	reverse := 0
	for i > 0 {
		// get the lowest digit
		digit := i - (i / 10 * 10)

		// build the reverse number
		reverse = (reverse * 10) + digit

		// shift i over one place - ie, 123 -> 12
		i = i / 10

		fmt.Println("digit=", digit, " reverse = ", reverse, " i=", i)
	}
	fmt.Println("-> x=", x, " reverse=", reverse)
	return reverse == x
}
