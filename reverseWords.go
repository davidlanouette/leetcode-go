package leetcode

import (
	"bytes"
	"strings"
)

// This DOES work for multi-byte characters.
func reverseWords(s string) string {

	var buf bytes.Buffer

	words := strings.Split(s, " ")

	numSpaces := strings.Count(s, " ")
	for n, word := range words {
		runes := []rune(word)
		for i := len(runes) - 1; i >= 0; i-- {
			buf.WriteString(string(runes[i]))
		}
		if n < numSpaces {
			buf.WriteString(" ")
		}
	}

	return buf.String()
}
