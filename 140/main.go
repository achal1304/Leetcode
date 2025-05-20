package main

import "fmt"

func main() {
	fmt.Println(wordBreak("catsanddog", []string{"cat", "cats", "and", "sand", "dog"}))
}

func wordBreak(s string, wordDict []string) []string {
	wordsDict := make(map[string]bool)

	for _, word := range wordDict {
		wordsDict[word] = true
	}

	memo := make(map[string][]string)
	return wordBreakHelper(s, wordsDict, memo)
}

func wordBreakHelper(s string, wordSet map[string]bool, memo map[string][]string) []string {
	if res, ok := memo[s]; ok {
		return res
	}
	fmt.Println(s)

	if len(s) == 0 {
		return []string{s}
	}

	var result []string

	for i := 1; i <= len(s); i++ {
		currWord := s[:i]
		if wordSet[currWord] {
			resultList := wordBreakHelper(s[i:], wordSet, memo)
			for _, ele := range resultList {
				if ele == "" {
					result = append(result, currWord)
				} else {

					result = append(result, currWord+" "+ele)
				}
			}
		}
	}

	memo[s] = result
	return result
}
