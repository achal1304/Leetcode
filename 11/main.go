package main

import "fmt"

func main() {
	// fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(maxArea([]int{1, 1}))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxArea(height []int) int {
	n := len(height)
	left, right := 0, n-1

	maxArea := 0
	for left < right {
		currArea := min(height[right], height[left]) * (right - left)
		maxArea = max(maxArea, currArea)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}
