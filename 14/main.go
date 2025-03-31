package main

import (
	"fmt"
	"sort"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	sort.Strings(strs)
	minStr := strs[0]

	for _, str := range strs[1:] {
		for !strings.HasPrefix(str, minStr) {
			minStr = minStr[:len(minStr)-1]
			if minStr == "" {
				return ""
			}
		}
	}

	return minStr
}

func main() {
	words := []string{"flower", "flow", "flowht"}
	fmt.Println("Longest Common Prefix:", longestCommonPrefix(words))
}
