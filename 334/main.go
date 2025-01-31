package main

import (
	"fmt"
)

func main() {
	fmt.Println(increasingTriplet([]int{1, 4, 2, 3, 4}))
	fmt.Println(increasingTriplet([]int{20, 100, 10, 12, 5, 13}))
	fmt.Println(increasingTriplet([]int{2, 1, 5, 0, 4, 6}))
}

func increasingTriplet(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}
	rightMax := make([]int, len(nums))
	currMax := nums[n-1]
	for i := n - 1; i >= 0; i-- {
		if nums[i] > currMax {
			currMax = nums[i]
		}
		rightMax[i] = currMax
	}

	currMin := nums[0]
	for i := 0; i < n; i++ {
		if nums[i] < currMin {
			currMin = nums[i]
		}
		if currMin < nums[i] && nums[i] < rightMax[i] {
			return true
		}
	}
	return false
}
