package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(mergeAlternately("abc", "pqrs"))
}

func mergeAlternately(word1, word2 string) string {
	minLen := len(word1)
	if len(word2) < minLen {
		minLen = len(word2)
	}

	var builder strings.Builder
	builder.Grow(len(word1) + len(word2)) // Preallocate memory

	for i := 0; i < minLen; i++ {
		builder.WriteByte(word1[i])
		builder.WriteByte(word2[i])
	}

	// Append the remaining part of the longer word
	builder.WriteString(word1[minLen:])
	builder.WriteString(word2[minLen:])

	return builder.String()
}
