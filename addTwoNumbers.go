package leetcode

import (
	"strconv"
)

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
