package main

import "fmt"

// Function to find the length of the longest increasing subsequence
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// dp array where dp[i] is the length of the longest increasing subsequence ending at i
	dp := make([]int, len(nums))
	// Initially, every element has a subsequence length of 1 (the element itself)
	for i := range dp {
		dp[i] = 1
	}

	// Fill the dp array
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	// The longest increasing subsequence is the maximum value in the dp array
	longest := 0
	for _, length := range dp {
		longest = max(longest, length)
	}
	return longest
}

// Helper function to return the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	result := lengthOfLIS(nums)
	fmt.Println(result) // Output: 4 (The longest increasing subsequence is [2, 3, 7, 101])
}
