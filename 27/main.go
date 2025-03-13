package main

import "fmt"

func main() {
	fmt.Println(removeElement([]int{1, 2, 3, 4, 2, 2, 4, 5, 6}, 2))
}

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	point := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[point] = nums[i]
			point++
		}
	}
	fmt.Println(nums)
	return point
}
