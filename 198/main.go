package main

import "fmt"

func main() {
	fmt.Println(rob([]int{1, 2, 4, 1, 3, 9, 8, 4}))
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		if nums[0] < nums[1] {
			return nums[1]
		} else {
			return nums[0]
		}
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = nums[1]
	dp[2] = nums[0] + nums[2]
	for i := 3; i < len(nums); i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-3]+nums[i])
	}

	return max(dp[len(nums)-1], dp[len(nums)-2])
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
