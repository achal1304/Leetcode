package main

import "fmt"

func main() {
	fmt.Println(canCompleteCircuit([]int{1, 2, 5, 2, 4, 3}, []int{4, 2, 2, 6, 2, 1}))
}

func canCompleteCircuit(gas []int, cost []int) int {
	if len(gas) == 0  {
		return 0
	}

	total := 0
	travelCost := make([]int, len(gas))
	for i, gasAtStation := range gas {
		travelCost[i] = gasAtStation - cost[i]
		total += travelCost[i]
	}

	if total < 0 {
		return -1
	}

	maxCost := 0
	maxIndex := 0
	currCost := 0
	for i := len(travelCost) - 1; i >= 0; i-- {
		currCost += travelCost[i]
		if currCost > maxCost {
			maxCost = currCost
			maxIndex = i
		}
	}

	return maxIndex
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
