package main

import (
	"fmt"
	"sort"
)

// Function to find the longest string chain
func longestStrChain(words []string) int {
	// Sort words by length
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	// dp will store the longest chain for each word
	dp := make(map[string]int)
	longestChain := 1

	// Process each word
	for _, word := range words {
		// For each word, try to form predecessors by removing one character
		dp[word] = 1 // Minimum chain length for any word is 1 (the word itself)
		for i := 0; i < len(word); i++ {
			prev := word[:i] + word[i+1:] // Remove the i-th character to form a predecessor
			if _, exists := dp[prev]; exists {
				dp[word] = max(dp[word], dp[prev]+1)
			}
		}
		// Track the longest chain
		longestChain = max(longestChain, dp[word])
	}

	return longestChain
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	words := []string{"abc", "a", "b", "ba", "bca", "bda", "bdca"}
	result := longestStrChain(words)
	fmt.Println(result) // Output: 4
}
