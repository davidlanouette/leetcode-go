package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {

	testVals := []struct {
		input    int
		expected int
	}{
		{123, 321},
		{-123, -321},
		{120, 21},
		{10000, 1},
		{1534236469, 0},
		{-2147483648, 0},
	}

	for _, v := range testVals {
		assert.Equal(t, v.expected, reverse(v.input))
	}
}
