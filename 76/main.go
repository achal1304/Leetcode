package main

import (
	"fmt"
)

// REVISIT : HARD

func main() {
	fmt.Println(minWindow("abbaa", "baa"))
}
func minWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}

	// Frequency map for characters in t
	tFreq := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		tFreq[t[i]]++
	}

	// Window frequency map
	windowFreq := make(map[byte]int)
	left, right := 0, 0
	minLen := len(s) + 1
	minStart := 0
	required := len(tFreq)
	formed := 0

	// Sliding window
	for right < len(s) {
		// Add the current character to the window
		char := s[right]
		windowFreq[char]++

		// If the current character's frequency matches the frequency in t, increase the formed count
		if windowFreq[char] == tFreq[char] {
			formed++
		}

		// Try to shrink the window if it's valid
		for left <= right && formed == required {
			windowSize := right - left + 1
			if windowSize < minLen {
				minLen = windowSize
				minStart = left
			}

			// Remove the character at the left pointer from the window
			windowFreq[s[left]]--
			if windowFreq[s[left]] < tFreq[s[left]] {
				formed--
			}
			left++
		}

		// Expand the window by moving the right pointer
		right++
	}

	// If no valid window was found
	if minLen == len(s)+1 {
		return ""
	}
	return s[minStart : minStart+minLen]
}
