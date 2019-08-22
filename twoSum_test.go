package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	//Given nums = [2, 7, 11, 15], target = 9,
	//	Because nums[0] + nums[1] = 2 + 7 = 9,
	// return [0, 1].
	input := []int{2, 7, 11, 15}
	result := twoSum(input, 9)
	assert.Equal(t, 2, len(result))

}
