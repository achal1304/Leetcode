package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(numDecodings("11106"))
}

func numDecodings(s string) int {
	dp := make(map[string]int)
	return numDecodingsWithDp(s, dp)
}

func numDecodingsWithDp(s string, dp map[string]int) int {
	if len(s) == 0 {
		return 1
	}

	if s[0] == '0' {
		return 0
	}

	if val, ok := dp[s]; ok {
		return val
	}

	currCount := 0

	for i := 1; i <= len(s); i++ {
		currString := s[:i]
		intVal, _ := strconv.Atoi(currString)
		if intVal >= 1 && intVal <= 26 {
			count := numDecodingsWithDp(s[i:], dp)
			currCount += count
		} else {
			dp[s] = currCount
			return currCount
		}
	}
	dp[s] = currCount
	return currCount
}
