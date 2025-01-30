package main

import (
	"fmt"
)

func main() {
	fmt.Println(reverseVowels("leetcode"))
}

var vowelsMap = map[byte]bool{
	'a': true,
	'e': true,
	'i': true,
	'o': true,
	'u': true,
	'A': true,
	'E': true,
	'I': true,
	'O': true,
	'U': true,
}

func reverseVowels(s string) string {
	l, r := 0, len(s)-1
	byteS := []byte(s)
	for r >= l {
		if vowelsMap[byteS[l]] {
			for !vowelsMap[byteS[r]] && r >= l {
				r -= 1
			}
			byteS[l], byteS[r] = byteS[r], byteS[l]
			r -= 1
		}
		l += 1
	}
	return string(byteS)
}
