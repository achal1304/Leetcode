package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minSubArrayLen(4, []int{1, 4, 4}))
}

func minSubArrayLen(target int, nums []int) int {

	minLen := math.MaxInt32
	start := 0
	currSum := 0

	for end := 0; end < len(nums); end++ {
		currSum += nums[end]
		for currSum >= target {
			if end-start+1 < minLen {
				minLen = end - start + 1
			}
			currSum -= nums[start]
			start++
		}
	}

	if minLen == math.MaxInt32 {
		return 0
	}
	return minLen
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
