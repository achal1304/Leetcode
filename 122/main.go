package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{1, 2, 4, 9, 8, 11, 10, 1, 11, 12, 16, 2, 3, 1, 6}))
}

func maxProfit(prices []int) int {
	if len(prices) == 0 || len(prices) == 1 {
		return 0
	}

	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-prices[i-1] > 0 {
			maxProfit += prices[i] - prices[i-1]
		}
	}
	return maxProfit
}
