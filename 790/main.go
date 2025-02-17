package main

import "fmt"

// REVISIT

func main() {
	fmt.Println(numTilings(4))
}

const mod = 1_000_000_000_7

func numTilings(n int) int {
	if n == 0 {
		return 1
	} else if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}

	dp := make([]int, n+1)
	dp[0], dp[1], dp[2] = 1, 1, 2
	for i := 3; i <= n; i++ {
		dp[i] = (2*dp[i-1] + dp[i-3]) % mod
	}
	return dp[n]
}
