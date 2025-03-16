package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{1, 2, 4, 6, 2, 4, 3, 7, 8, 3, 8, 9}))
}

func maxProfit(prices []int) int {
	if len(prices) == 0 || len(prices) == 1 {
		return 0
	}

	maxProfit, minEle := 0, prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] > minEle {
			if prices[i]-minEle > maxProfit {
				maxProfit = prices[i] - minEle
			}
		} else {
			minEle = prices[i]
		}
	}
	return maxProfit
}
