package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	testVals := []struct {
		nums     []int
		target   int
		expected bool
	}{
		{[]int{}, 5, false},
		{[]int{2, 5, 6, 0, 0, 1, 2}, 0, true},
		{[]int{2, 5, 6, 0, 0, 1, 2}, 3, false},
		{[]int{2, 5, 6, 0, 0, 1, 2}, 0, true},
	}

	for _, v := range testVals {
		assert.Equal(t, v.expected, search(v.nums, v.target))
	}
}
func TestSearch2(t *testing.T) {
	testVals := []struct {
		nums     []int
		target   int
		expected bool
	}{
		{[]int{}, 5, false},
		{[]int{2, 5, 6, 0, 0, 1, 2}, 0, true},
		{[]int{2, 5, 6, 0, 0, 1, 2}, 3, false},
		{[]int{2, 5, 6, 0, 0, 1, 2}, 0, true},
	}

	for _, v := range testVals {
		assert.Equal(t, v.expected, search2(v.nums, v.target))
	}
}
