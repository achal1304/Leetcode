package main

import (
	"fmt"
	"math"
)

func main() {
	cost2 := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}

	fmt.Println("Minimum cost for cost2:", minCostClimbingStairs(cost2))
}

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return cost[0]
	}

	dp := make([]int, n)
	dp[0] = cost[0]
	dp[1] = cost[1]

	for i := 2; i < n; i++ {
		dp[i] = cost[i] + int(math.Min(float64(dp[i-1]), float64(dp[i-2])))
	}

	return int(math.Min(float64(dp[n-1]), float64(dp[n-2])))
}
