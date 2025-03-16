package main

import "fmt"

func main() {
	fmt.Println(canJump([]int{1, 1, 3, 4, 1, 0, 0, 1}))
}

func canJump(nums []int) bool {
	farthest := 0

	for i, jump := range nums {
		if i > farthest {
			return false
		}
		farthest = max(farthest, i+jump)
		if farthest >= len(nums)-1 {
			return true
		}
	}

	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
