package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false // Strings must have the same length
	}

	// Create two maps to store character mappings
	mapS := make(map[byte]byte)
	mapT := make(map[byte]byte)

	// Traverse both strings
	for i := 0; i < len(s); i++ {
		// Check mapping for s -> t
		if mappedCharS, ok := mapS[s[i]]; ok {
			if mappedCharS != t[i] {
				return false // If there's a conflict, return false
			}
		} else {
			mapS[s[i]] = t[i] // Create mapping from s[i] to t[i]
		}

		// Check mapping for t -> s
		if mappedCharT, ok := mapT[t[i]]; ok {
			if mappedCharT != s[i] {
				return false // If there's a conflict, return false
			}
		} else {
			mapT[t[i]] = s[i] // Create mapping from t[i] to s[i]
		}
	}

	return true
}

func main() {
	// Test case 1
	s1, t1 := "egg", "add"
	fmt.Println(isIsomorphic(s1, t1)) // Expected output: true

	// Test case 2
	s2, t2 := "foo", "bar"
	fmt.Println(isIsomorphic(s2, t2)) // Expected output: false

	// Test case 3
	s3, t3 := "paper", "title"
	fmt.Println(isIsomorphic(s3, t3)) // Expected output: true
}
