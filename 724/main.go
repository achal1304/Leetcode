package main

import "fmt"

func pivotIndex(nums []int) int {
	n := len(nums)
	leftSum := make([]int, n)
	for i := 1; i < n; i++ {
		leftSum[i] = leftSum[i-1] + nums[i-1]
	}

	rightSum := make([]int, n)
	for i := n - 2; i >= 0; i-- {
		rightSum[i] = rightSum[i+1] + nums[i+1]
	}

	for i := 0; i < n; i++ {
		if leftSum[i] == rightSum[i] {
			return i
		}
	}
	return -1
}

func main() {
	// Example 1
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))
}
