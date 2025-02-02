package main

import "fmt"

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

func maxVowels(s string, k int) int {
	n := len(s)
	if k > n {
		return -1
	}

	maxCount := 0
	for i := 0; i < k; i++ {
		if vowelsMap[s[i]] {
			maxCount++
		}
	}

	fmt.Println("maxCount ", maxCount)
	checkCount := maxCount
	for i := k; i < n; i++ {
		if vowelsMap[s[i-k]] {
			checkCount--
		}
		if vowelsMap[s[i]] {
			checkCount++
			maxCount = max(maxCount, checkCount)
		}

		fmt.Println("current masx and check count ", checkCount, maxCount, i)
	}

	return maxCount
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// fmt.Println(maxVowels("abciiidef", 3))
	fmt.Println(maxVowels("bcdarettitotuu", 2))
}
