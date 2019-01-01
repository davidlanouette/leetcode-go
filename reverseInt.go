package leetcode

import "math"

/*
7. Reverse Integer
Easy
18032571FavoriteShare
Given a 32-bit signed integer, reverse digits of an integer.

Example 1:
	Input: 123
	Output: 321
Example 2:
	Input: -123
	Output: -321
Example 3:
	Input: 120
	Output: 21

Note:
Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1]. For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.
*/

func reverse(x int) int {
	var out int

	for x != 0 {
		out *= 10
		tmp := x % 10

		out += tmp
		x /= 10
	}

	if out > math.MaxInt32 || out < math.MinInt32 {
		return 0
	}

	return out
}
