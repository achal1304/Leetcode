package main

import (
	"fmt"
)

// REVISIT
func main() {
	fmt.Println(maxProfit([]int{1, 3, 7, 5, 10, 3}, 3))
}

func maxProfit(prices []int, fee int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

	cash := 0
	hold := -prices[0]

	for i := 1; i < n; i++ {
		cash = max(cash, hold+prices[i]-fee)
		hold = max(hold, cash-prices[i])
	}

	return cash
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
