package leetcode

import (
	"bytes"
	"strings"
)

// This does NOT work for multi-byte characters.
func reverseWords(s string) string {

	var buf bytes.Buffer

	words := strings.Split(s, " ")

	numSpaces := strings.Count(s, " ")
	for n, word := range words {
		for i := len(word) - 1; i >= 0; i-- {
			buf.WriteString(string(word[i]))
		}
		if n < numSpaces {
			buf.WriteString(" ")
		}
	}

	return buf.String()
}
