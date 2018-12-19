package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Input: nums = [2,5,6,0,0,1,2], target = 0
// Output: true
func TestSearch_HappyPath(t *testing.T) {
	nums := []int{2, 5, 6, 0, 0, 1, 2}
	target := 0

	assert.True(t, search(nums, target))
}

func TestSearch_TargetIsntInArray(t *testing.T) {
	nums := []int{2, 5, 6, 0, 0, 1, 2}
	target := 3

	assert.False(t, search(nums, target))
}

func TestSearch_LeetcodeInput1(t *testing.T) {
	nums := []int{2, 5, 6, 0, 0, 1, 2}
	target := 0

	assert.True(t, search(nums, target))
}
