package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxOperations([]int{1, 2, 3, 4}, 5))
}

func maxOperations(nums []int, k int) int {
	left, right := 0, len(nums)-1
	count := 0
	sort.Ints(nums)
	for left < right {
		currSum := nums[left] + nums[right]
		if currSum == k {
			count++
			left++
			right--
			continue
		}

		if currSum > k {
			right--
		} else {
			left++
		}
	}
	return count
}
