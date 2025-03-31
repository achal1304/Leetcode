package main

import (
	"fmt"
)

func strStr(haystack string, needle string) int {
	hLen, nLen := len(haystack), len(needle)

	if nLen == 0 {
		return 0
	}

	for i := 0; i <= hLen-nLen; i++ {
		if haystack[i:i+nLen] == needle {
			return i
		}
	}

	return -1
}

func main() {
	// Example usage
	fmt.Println(strStr("sadbutsad", "sad"))  // Output: 0
	fmt.Println(strStr("leetcode", "leeto")) // Output: -1
}
