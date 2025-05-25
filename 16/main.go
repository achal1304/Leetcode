package main

import (
	"fmt"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	if len(nums) < 3 {
		return -1
	}

	currMax := 0
	currMinDiff := 9999999999
	for i := 0; i < len(nums); i++ {
		left := i + 1
		right := len(nums) - 1
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for left < right {
			fmt.Println(i, left, right)
			sum := nums[i] + nums[left] + nums[right]
			diff := abs(sum - target)
			if diff < currMinDiff {
				currMax = sum
				currMinDiff = diff
			}

			if sum == target {
				return sum
			} else if sum < target {
				left++
			} else {
				right--
			}
		}

	}
	return currMax
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func main() {
	fmt.Println(threeSumClosest([]int{1, 1, 1, 1}, 0))
}
