package main

import "fmt"

func isAnagram(s string, t string) bool {
	// If the lengths are not equal, they cannot be anagrams
	if len(s) != len(t) {
		return false
	}

	// Create frequency maps for both strings
	countS := make(map[rune]int)
	countT := make(map[rune]int)

	// Count the characters in both strings
	for _, char := range s {
		countS[char]++
	}
	for _, char := range t {
		countT[char]++
	}

	// Compare the frequency maps
	for char, count := range countS {
		if countT[char] != count {
			return false
		}
	}

	return true
}

func main() {
	// Test case 1
	s1, t1 := "anagram", "nagaram"
	fmt.Println(isAnagram(s1, t1)) // Expected output: true

	// Test case 2
	s2, t2 := "rat", "car"
	fmt.Println(isAnagram(s2, t2)) // Expected output: false
}
