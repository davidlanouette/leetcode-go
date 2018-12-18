package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckPossibility(t *testing.T) {
	assert.True(t, checkPossibility([]int{1, 2, 3}))
	assert.True(t, checkPossibility([]int{4, 2, 3}))
	assert.True(t, checkPossibility([]int{-1, 4, 2, 3}))

	assert.False(t, checkPossibility([]int{4, 2, 1}))
	assert.False(t, checkPossibility([]int{4, 2, 3, 4, 3}))
	assert.False(t, checkPossibility([]int{3, 4, 2, 3}))
}

func TestIsSorted(t *testing.T) {
	assert.True(t, isSorted([]int{1, 2, 3}))
	assert.True(t, isSorted([]int{4, 12, 13}))
	assert.True(t, isSorted([]int{-1, 0, 2, 3}))

	assert.False(t, isSorted([]int{4, 2, 1}))
	assert.False(t, isSorted([]int{4, 2, 3, 4, 3}))
	assert.False(t, isSorted([]int{2, 2, 3, 4, 3}))
	assert.False(t, isSorted([]int{3, 4, 2, 3}))
}
