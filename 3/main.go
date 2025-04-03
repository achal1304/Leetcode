package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("tmmzuxt"))
}

func lengthOfLongestSubstring(s string) int {
	visited := make(map[byte]int)
	start := 0
	maxLen := 0

	for end := 0; end < len(s); end++ {
		if id, ok := visited[s[end]]; ok {
			if id >= start {
				start = id + 1
			}
			visited[s[end]] = end
		} else {
			visited[s[end]] = end
		}
		if end-start+1 > maxLen {
			fmt.Println("currmax ", s[start:end+1])
			maxLen = end - start + 1
		}
	}
	return maxLen
}
