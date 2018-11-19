package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidNumber(t *testing.T) {
	assert.False(t, isValidNumber("0"))
	assert.False(t, isValidNumber("00"))
	assert.False(t, isValidNumber("0.0"))
	assert.False(t, isValidNumber("00.1"))
	assert.False(t, isValidNumber(".1"))
	assert.False(t, isValidNumber("1."))
	assert.False(t, isValidNumber("1.0"))
	assert.False(t, isValidNumber("1.00"))

	assert.True(t, isValidNumber("0.1"))
	assert.True(t, isValidNumber("0.01"))
	assert.True(t, isValidNumber("0.001"))
	assert.True(t, isValidNumber("1"))
	assert.True(t, isValidNumber("1.1"))
	assert.True(t, isValidNumber("1.01"))
	assert.True(t, isValidNumber("10.1"))

}
func TestPermutations(t *testing.T) {
	results := permutations("12")
	assert.Equal(t, 2, len(results))
	assert.Equal(t, "12", results[0])
	assert.Equal(t, "1.2", results[1])

	results = permutations("123")
	assert.Equal(t, 3, len(results))
	assert.Equal(t, "123", results[0])
	assert.Equal(t, "1.23", results[1])
	assert.Equal(t, "12.3", results[2])
}

func TestAmbiguousCoordinates_testCaseOne(t *testing.T) {
	results := ambiguousCoordinates("(123)")

	assert.Equal(t, 4, len(results))
	assert.Equal(t, "(1, 23)", results[0])
	assert.Equal(t, "(1, 2.3)", results[1])
	assert.Equal(t, "(12, 3)", results[2])
	assert.Equal(t, "(1.2, 3)", results[3])
}

func TestAmbiguousCoordinates_testCaseTwo(t *testing.T) {
	results := ambiguousCoordinates("(00011)")

	assert.Equal(t, 2, len(results))
	assert.Equal(t, "(0.001, 1)", results[0])
	assert.Equal(t, "(0, 0.011)", results[1])
}

func TestAmbiguousCoordinates_testCaseThree(t *testing.T) {
	results := ambiguousCoordinates("(0123)")

	assert.Equal(t, 6, len(results))
	assert.Equal(t, "(0, 123)", results[0])
	assert.Equal(t, "(0, 12.3)", results[1])
	assert.Equal(t, "(0, 1.23)", results[2])
	assert.Equal(t, "(0.1, 23)", results[3])
	assert.Equal(t, "(0.1, 2.3)", results[4])
	assert.Equal(t, "(0.12, 3)", results[5])
}

func TestAmbiguousCoordinates_testCaseFour(t *testing.T) {
	results := ambiguousCoordinates("(1234)")

	assert.Equal(t, 10, len(results))
	assert.Equal(t, "(1, 234)", results[0])
	assert.Equal(t, "(1, 2.34)", results[1])
	assert.Equal(t, "(1, 23.4)", results[2])
	assert.Equal(t, "(12, 34)", results[3])
	assert.Equal(t, "(12, 3.4)", results[4])
	assert.Equal(t, "(1.2, 34)", results[5])
	assert.Equal(t, "(1.2, 3.4)", results[6])
	assert.Equal(t, "(123, 4)", results[7])
	assert.Equal(t, "(1.23, 4)", results[8])
	assert.Equal(t, "(12.3, 4)", results[9])

}
