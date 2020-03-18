package leetcode

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.
You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:

	Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
	Output: 7 -> 0 -> 8
	Explanation: 342 + 465 = 807.
*/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var head, prev, node *ListNode
	var num1, num2, leftOver int

	for (l1 != nil) || (l2 != nil) {
		if l1 != nil {
			num1 = l1.Val
		} else {
			num1 = 0
		}
		if l2 != nil {
			num2 = l2.Val
		} else {
			num2 = 0
		}

		sum := num1 + num2 + leftOver

		leftOver = sum / 10

		node = &ListNode{sum % 10, nil}

		if prev == nil {
			head = node
		} else {
			prev.Next = node
		}
		prev = node

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	if leftOver != 0 {
		node.Next = &ListNode{leftOver, nil}
	}

	return head
}

// ListNode is a node in a singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Print converts the linked list into a string
func (ln *ListNode) Print() string {
	out := ""

	for ln != nil {
		out = out + strconv.Itoa(ln.Val)
		ln = ln.Next
	}

	return out
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
