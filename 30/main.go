package main

import "fmt"

// REVISIT - HARD

func main() {
	fmt.Println(findSubstring("barfoothefoobarman", []string{"bar", "foo"}))
}

func findSubstring(s string, words []string) []int {
	wordsLen := len(words[0])
	fixDIct := make(map[string]int)

	for i, _ := range words {
		fixDIct[words[i]]++
	}
	ans := []int{}

	totalSubLen := len(words) * wordsLen

	for i := 0; i < wordsLen; i++ { // only loop till wordlen as after that the pattern will repeat
		start := i
		wordDict := make(map[string]int)
		for end := i; end+wordsLen <= len(s); end += wordsLen {
			word := s[end : end+wordsLen]
			if _, ok := fixDIct[word]; ok {
				wordDict[word]++

				// Loop till the latest word frequency is maintined. decrease startword counter and increase its count
				// as we previously increased it in line 28. so as we move our window we need to restorce the previous count
				for wordDict[word] > fixDIct[word] && start+wordsLen <= len(s) {
					startWord := s[start : start+wordsLen]
					wordDict[startWord]--
					start += wordsLen
				}

				if end-start+wordsLen == totalSubLen {
					ans = append(ans, start)
				}
			} else {
				wordDict = make(map[string]int)
				start = end + wordsLen
			}
		}
	}
	return ans
}
