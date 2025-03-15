package main

import "fmt"

func main() {
	rotate([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 13)
}

func rotate(nums []int, k int) {
	if k == 0 {
		return
	}

	if k > len(nums) {
		k = k % len(nums)
	}

	reverse(nums, 0, len(nums)-1)
	fmt.Println(nums)

	reverse(nums, 0, k-1)
	fmt.Println(nums)

	reverse(nums, k, len(nums)-1)
	fmt.Println(nums)
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
