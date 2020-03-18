package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
Given an array of integers, return indices of the two numbers such that they add up to a specific target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

	Given nums = [2, 7, 11, 15], target = 9,

	Because nums[0] + nums[1] = 2 + 7 = 9,
	return [0, 1].

*/
func twoSum(nums []int, target int) []int {
	// fmt.Println("nums =", nums, ", target =", target)
	numMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		// fmt.Println("numMap =", numMap)
		j, ok := numMap[target-nums[i]]
		// fmt.Println("i =", i, ", nums[i] =", nums[i], ", target-nums[i] =", target-nums[i], ", j =", j, ", ok =", ok)
		if ok && i != j {
			return []int{j, i}
		}
		numMap[nums[i]] = i
	}
	// couldn't find a match.  return empty slice
	return []int{}
}

func TestTwoSum(t *testing.T) {

	testVals := []struct {
		input    []int
		target   int
		expected []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{2, 7, 11, 15}, 13, []int{0, 2}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
	}

	for _, v := range testVals {
		result := twoSum(v.input, v.target)
		assert.Equal(t, 2, len(v.expected))
		assert.Equal(t, v.expected[0], result[0])
		assert.Equal(t, v.expected[1], result[1])
	}
}

func BenchmarkTwoSum(b *testing.B) {
	testVals := []struct {
		input    []int
		target   int
		expected []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{2, 7, 11, 15}, 13, []int{0, 2}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
	}

	var result []int

	for _, v := range testVals {
		for i := 0; i < b.N; i++ {
			result = twoSum(v.input, v.target)
		}
	}
	fmt.Print(result)
}
