package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
