package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanJump(t *testing.T) {

	assert.True(t, canJump([]int{0}))
	assert.True(t, canJump([]int{2, 3, 1, 1, 4}))
	assert.False(t, canJump([]int{3, 2, 1, 0, 4}))

	// This fails!  Fix it.  Secret is that each move is a MAX number of places to move, not absolute number.
	assert.False(t, canJump([]int{2, 5, 0, 0}))
}
