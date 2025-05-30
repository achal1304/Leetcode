package main

import "fmt"

var Directions = [][]int{
	{1, 0},
	{0, 1},
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	dp := make([][]int, m)
	for i, _ := range dp {
		dp[i] = make([]int, n)
	}

	if obstacleGrid[0][0] != 1 {
		dp[0][0] = 1
	} else {
		dp[0][0] = 0
	}
	for i := 1; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			if n == 1 {
				return 0
			}
			dp[i][0] = 0
		} else {
			dp[i][0] = dp[i-1][0]
		}
	}
	for j := 1; j < n; j++ {
		if obstacleGrid[0][j] == 1 {
			if m == 1 {
				dp[0][j] = 0
			}
		} else {
			dp[0][j] = dp[0][j-1]
		}
	}
	fmt.Println("dp out", dp)
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			fmt.Println(dp)
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0

			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
