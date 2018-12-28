package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructMaximumBinaryTree(t *testing.T) {

	data := []struct {
		input    []int
		expected string
	}{
		{
			[]int{3, 2, 1, 6, 0, 5},
			"[6,3,5,null,2,0,null,null,1]",
		},
	}

	for _, d := range data {

		n := constructMaximumBinaryTree(d.input)
		fmt.Println("n", n, "expected", d.expected)
		assert.Equal(t, d.expected, n.String())
	}
}
