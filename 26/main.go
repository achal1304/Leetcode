package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 1, 2, 2, 2, 4, 4, 4, 5, 5, 5, 6, 7}))
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return 1
	}

	point := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[point] = nums[i]
			point++
		}
	}

	fmt.Println(nums)
	return point
}
