package leetcode

import (
	"fmt"
	"math"
	"strings"
)

// TreeNode is the definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (n *TreeNode) String() string {
	var buf strings.Builder
	q := make([]*TreeNode, 0)
	var i *TreeNode
	q = append(q, n)

	buf.WriteString("[")
	for len(q) != 0 {
		if buf.Len() > 1 {
			buf.WriteString(",")
		}
		i = q[0]
		q = q[1:]
		if i != nil {
			buf.WriteString(fmt.Sprintf("%v", i.Val))
			if i.Left != nil || i.Right != nil {
				q = append(q, i.Left)
				q = append(q, i.Right)
			}
		} else {
			buf.WriteString("null")
		}
	}
	buf.WriteString("]")
	return buf.String()
}

func constructMaximumBinaryTree(nums []int) *TreeNode {

	fmt.Println("nums", nums)
	if len(nums) == 0 {
		return nil
	}

	maxIdx, maxVal := findMax(nums)
	fmt.Println("maxIdx", maxIdx, "maxVal", maxVal)

	root := &TreeNode{
		Val:   maxVal,
		Left:  constructMaximumBinaryTree(nums[0:maxIdx]),
		Right: constructMaximumBinaryTree(nums[maxIdx+1:]),
	}

	return root
}

func findMax(nums []int) (int, int) {
	maxIdx := 0
	maxVal := math.MinInt32

	for i, v := range nums {
		if v > maxVal {
			maxIdx = i
			maxVal = v
		}
	}

	return maxIdx, maxVal
}
