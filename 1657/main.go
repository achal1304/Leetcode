package main

import (
	"fmt"
	"sort"
)

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	freq1 := make(map[rune]int)
	freq2 := make(map[rune]int)

	for _, ch := range word1 {
		freq1[ch]++
	}
	for _, ch := range word2 {
		freq2[ch]++
	}

	uniqueChars1 := make(map[rune]struct{})
	uniqueChars2 := make(map[rune]struct{})

	for ch := range freq1 {
		uniqueChars1[ch] = struct{}{}
	}
	for ch := range freq2 {
		uniqueChars2[ch] = struct{}{}
	}

	if len(uniqueChars1) != len(uniqueChars2) {
		return false
	}

	for ch := range uniqueChars1 {
		if _, exists := uniqueChars2[ch]; !exists {
			return false
		}
	}

	counts1 := make([]int, 0, len(freq1))
	counts2 := make([]int, 0, len(freq2))

	for _, count := range freq1 {
		counts1 = append(counts1, count)
	}
	for _, count := range freq2 {
		counts2 = append(counts2, count)
	}

	sort.Ints(counts1)
	sort.Ints(counts2)

	for i := range counts1 {
		if counts1[i] != counts2[i] {
			return false
		}
	}

	return true
}

func main() {
	word1 := "cabbba"
	word2 := "abbccc"
	fmt.Println(closeStrings(word1, word2))
}
