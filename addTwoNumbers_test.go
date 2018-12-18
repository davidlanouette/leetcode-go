package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumbers(t *testing.T) {
	assert.True(t, true)

	testData := []struct {
		num1     []int
		num2     []int
		expected []int
	}{
		{
			[]int{2, 4, 3},
			[]int{5, 6, 4},
			[]int{7, 0, 8},
		},
		{
			[]int{9, 9, 9, 9},
			[]int{9, 9, 9, 9},
			[]int{8, 9, 9, 9, 1},
		},
		{
			[]int{3, 2, 1},
			[]int{1},
			[]int{4, 2, 1},
		},
		{
			[]int{1},
			[]int{3, 2, 1},
			[]int{4, 2, 1},
		},
	}

	for _, tt := range testData {
		expected := sliceToList(tt.expected)
		num1 := sliceToList(tt.num1)
		num2 := sliceToList(tt.num2)
		actual := addTwoNumbers(num1, num2)
		fmt.Printf("num1: %v, num2: %v, expected: %v, actual: %v\n",
			num1.Print(), num2.Print(), expected.Print(), actual.Print())
		assert.Equal(t, expected.Print(), actual.Print())
	}
}

func sliceToList(nums []int) *ListNode {

	var head, prev *ListNode

	for _, i := range nums {
		node := &ListNode{i, nil}
		if prev == nil {
			head = node
		} else {
			prev.Next = node
		}
		prev = node
	}
	return head
}
