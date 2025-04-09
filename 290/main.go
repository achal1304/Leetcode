package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, s string) bool {
	// Split the string s into words by spaces
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false // The lengths must match
	}

	// Two maps: one for letter to word mapping, and one for word to letter mapping
	letterToWord := make(map[byte]string)
	wordToLetter := make(map[string]byte)

	// Traverse both the pattern and the words simultaneously
	for i := 0; i < len(pattern); i++ {
		letter := pattern[i]
		word := words[i]

		// Check if the current letter is already mapped to a different word
		if mappedWord, exists := letterToWord[letter]; exists {
			if mappedWord != word {
				return false // Inconsistent mapping
			}
		} else {
			letterToWord[letter] = word
		}

		// Check if the current word is already mapped to a different letter
		if mappedLetter, exists := wordToLetter[word]; exists {
			if mappedLetter != letter {
				return false // Inconsistent mapping
			}
		} else {
			wordToLetter[word] = letter
		}
	}

	// If all checks passed, the mapping is consistent
	return true
}

func main() {
	// Test case 1
	fmt.Println(wordPattern("abba", "dog cat cat dog")) // Expected output: true

	// Test case 2
	fmt.Println(wordPattern("abba", "dog cat cat fish")) // Expected output: false

	// Test case 3
	fmt.Println(wordPattern("aaaa", "dog dog dog dog")) // Expected output: true

	// Test case 4
	fmt.Println(wordPattern("abba", "dog dog dog dog")) // Expected output: false
}
