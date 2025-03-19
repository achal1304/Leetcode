package main

import "fmt"

// REVISIT
func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}

func trap(height []int) int {
	if len(height) <= 2 {
		return 0
	}

	maxLeft := height[0]
	maxHeight := make([]int, len(height))
	maxHeight[0] = height[0]
	for i := 0; i < len(height); i++ {
		if height[i] > maxLeft {
			maxLeft = height[i]
		}
		maxHeight[i] = maxLeft
	}

	maxRight := height[len(height)-1]
	for i := len(height) - 1; i >= 0; i-- {
		if height[i] > maxRight {
			maxRight = height[i]
		}
		maxHeight[i] = min(maxHeight[i], maxRight)
	}

	totalWater := 0

	for i := 0; i < len(maxHeight); i++ {
		totalWater += diff(height[i], maxHeight[i])
	}
	return totalWater
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func diff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
