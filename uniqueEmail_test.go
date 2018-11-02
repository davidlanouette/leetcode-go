package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumUniqueEmails(t *testing.T) {
	in := []string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}
	assert.Equal(t, 2, numUniqueEmails(in))

	in = []string{"a@leetcode.com", "b@leetcode.com", "c@leetcode.com"}
	assert.Equal(t, 3, numUniqueEmails(in))
}
