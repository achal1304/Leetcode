package main

func main() {}

// REVISIT - in dp we are trying to check for each coin how much minimum coins of that amount is required
// so we are taking one coin as current then taking the ramining amount as previous dp value which is i-coin
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0

	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}

	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}
