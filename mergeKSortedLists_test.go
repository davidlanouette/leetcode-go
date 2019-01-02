package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeKLists(t *testing.T) {
	assert.True(t, true)

	testData := []struct {
		input    [][]int
		expected []int
	}{
		{
			[][]int{
				[]int{2, 3, 4, 5},
				[]int{1, 5, 6, 14},
			},
			[]int{1, 2, 3, 4, 5, 5, 6, 14},
		},
		{
			[][]int{
				[]int{9, 9, 9, 9},
				[]int{9, 9, 9, 9},
			},
			[]int{9, 9, 9, 9, 9, 9, 9, 9},
		},
		{
			[][]int{
				[]int{1, 2, 3},
				[]int{9, 9, 9, 9},
				[]int{1, 2, 3, 9, 9, 9, 9},
			},
			[]int{1, 1, 2, 2, 3, 3, 9, 9, 9, 9, 9, 9, 9, 9},
		},
		{
			[][]int{
				[]int{1},
				[]int{2, 2, 3},
			},
			[]int{1, 2, 2, 3},
		},
	}

	for _, tt := range testData {
		expected := sliceToList(tt.expected)
		input := make([]*ListNode, len(tt.input))
		for i, s := range tt.input {
			n := sliceToList(s)
			input[i] = n
		}
		actual := mergeKLists(input)
		assert.Equal(t, expected.Print(), actual.Print())
	}
}
