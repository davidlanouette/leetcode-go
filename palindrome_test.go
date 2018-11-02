package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalendrome(t *testing.T) {
	assert.True(t, isPalindrome(121))
	assert.True(t, isPalindrome(1211221121))
	assert.True(t, isPalindrome(1234567890987654321))

	assert.False(t, isPalindrome(-121))
	assert.False(t, isPalindrome(123))
	assert.False(t, isPalindrome(1233210))

}
