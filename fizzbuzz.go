package leetcode

import "strconv"

/*
Write a program that outputs the string representation of numbers from 1 to n.

But for multiples of three it should output “Fizz” instead of the number and for the multiples of five output “Buzz”. For numbers which are multiples of both three and five output “FizzBuzz”.
*/

func fizzBuzz(n int) []string {
	out := make([]string, n)
	for i := 1; i <= n; i++ {
		isFizz := i%3 == 0
		isBuzz := i%5 == 0
		if isFizz && isBuzz {
			out[i-1] = "FizzBuzz"
		} else if isFizz {
			out[i-1] = "Fizz"
		} else if isBuzz {
			out[i-1] = "Buzz"
		} else {
			out[i-1] = strconv.Itoa(i)
		}
	}
	return out
}
