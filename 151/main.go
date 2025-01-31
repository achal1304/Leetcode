package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseWords("the sky is blue")) // Output: "blue is sky the"
	fmt.Println(reverseWords("  hello world  ")) // Output: "world hello"
}

func reverseWords(s string) string {
	s = strings.TrimSpace(s)

	words := strings.Fields(s)

	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	return strings.Join(words, " ")
}
