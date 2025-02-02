package main

import "fmt"

func main() {
	// fmt.Println(longestOnes([]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3))
	fmt.Println(longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2))
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

func longestOnes(nums []int, k int) int {
	n := len(nums)
	l, r := 0, 0

	maxCount, currCount := 0, 0
	for r < n {
		if nums[r] == 0 && k > 0 {
			currCount++
			k--
		} else if nums[r] == 0 && k <= 0 {
			// This loop handles providing us one more flip by moving the window to the last zero used
			maxCount = max(maxCount, len(nums[l:r]))
			if nums[l] == 0 {
				l++ // Just move to next number if current l is 0. So we will have 1 more flip which can be used by current 0
			} else {
				for nums[l] != 0 {
					l++
				}
				l++ // Increment one more time for cases like 1,1,1,0. So in for loop l will go till index 2 and then we need to reach 0 so increment one more time
			}
		} else {
			currCount++
		}
		r++
	}
	return max(maxCount, len(nums[l:r]))
}
