package leetcode

import "math"

/*
23. Merge k Sorted Lists
Hard

1822

118

Favorite

Share
Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.

Example:

Input:
[
  1->4->5,
  1->3->4,
  2->6
]
Output: 1->1->2->3->4->4->5->6
*/

func mergeKLists(lists []*ListNode) *ListNode {

	var head, lastNode *ListNode

	for {
		lastIdx := -1
		lastVal := math.MaxInt32

		// find the minimum value in the lists
		for i, list := range lists {
			if list != nil {
				if list.Val <= lastVal {
					lastVal = list.Val
					lastIdx = i
				}
			}
		}

		// if we never changed lastIdx, then all the lists are nil
		if lastIdx == -1 {
			break
		}

		// once min is found, move the head of that list to the next item
		lists[lastIdx] = lists[lastIdx].Next

		// now, create the list, and add it to outbound list
		node := &ListNode{lastVal, nil}
		if lastNode == nil {
			lastNode = node
		} else {
			lastNode.Next = node
			lastNode = lastNode.Next
		}

		if head == nil {
			head = lastNode
		}
	}

	return head
}
