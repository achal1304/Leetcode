package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	// Frequency map to count characters in magazine
	magazineCount := make(map[rune]int)

	// Count characters in magazine
	for _, char := range magazine {
		magazineCount[char]++
	}

	// Check if each character in ransomNote can be formed using characters in magazine
	for _, char := range ransomNote {
		if magazineCount[char] > 0 {
			magazineCount[char]-- // Use one instance of the character from the magazine
		} else {
			return false // Not enough characters in magazine
		}
	}

	return true
}

func main() {
	// Test case 1
	ransomNote := "aabb"
	magazine := "ababbab"
	fmt.Println(canConstruct(ransomNote, magazine)) // Expected output: true

	// Test case 2
	ransomNote = "aa"
	magazine = "ab"
	fmt.Println(canConstruct(ransomNote, magazine)) // Expected output: false
}
