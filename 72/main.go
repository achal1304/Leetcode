package main

import "fmt"

// REVISIT
func main() {
	fmt.Println(minDistance("toseanss", "moseandtoss"))
}

func minDistance(word1 string, word2 string) int {
	// Create a 2D matrix to store the distances
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)

	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// Initialize base cases
	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}

	// Fill the dp table
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			cost := 0
			if word1[i-1] != word2[j-1] {
				cost = 1
			}
			dp[i][j] = min(dp[i-1][j]+1, min(dp[i][j-1]+1, dp[i-1][j-1]+cost))
		}
	}

	return dp[m][n]
}
