package main

import "fmt"

func main() {
	fmt.Println(longestCommonSubsequence("abcba", "abcbcba"))
}

func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text1)+1)
	for i, _ := range dp {
		dp[i] = make([]int, len(text2)+1)
	}

	for i := 1; i < len(text1)+1; i++ {
		for j := 1; j < len(text2)+1; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], (dp[i][j-1]))
			}
		}
	}
	for _, ele := range dp {
		fmt.Println(ele)
	}
	return dp[len(text1)][len(text2)]
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
