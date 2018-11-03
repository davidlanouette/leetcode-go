package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz(t *testing.T) {
	results := fizzBuzz(15)
	fmt.Println("results", results)

	assert.Equal(t, 15, len(results))

	assert.Equal(t, "1", results[0])
	assert.Equal(t, "2", results[1])
	assert.Equal(t, "Fizz", results[2])
	assert.Equal(t, "4", results[3])
	assert.Equal(t, "Buzz", results[4])

	assert.Equal(t, "Fizz", results[5])
	assert.Equal(t, "Buzz", results[9])
	assert.Equal(t, "FizzBuzz", results[14])
}
