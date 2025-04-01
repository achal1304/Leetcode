package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	start := 0
	end := len(s) - 1
	for start < end {
		for !isAlphanumeric(rune(s[start])) && start < end {
			start++
		}

		for !isAlphanumeric(rune(s[end])) && end > start {
			end--
		}

		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}

func isAlphanumeric(s rune) bool {
	return unicode.IsDigit(s) || unicode.IsLetter(s)
}
