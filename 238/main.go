package main

import (
	"fmt"
)

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
	fmt.Println(productExceptSelf([]int{-1, 1, 0, -3, 3}))
}

func productExceptSelf(nums []int) []int {
	n := len(nums)
	if n <= 1 {
		return nums
	}

	finalProduct := make([]int, n)

	finalProduct[0] = 1
	finalProduct[1] = nums[0]
	for i := 2; i < n; i++ {
		finalProduct[i] = nums[i-1] * finalProduct[i-1]
	}

	rightProduct := 1
	for i := n - 1; i >= 0; i-- {
		finalProduct[i] *= rightProduct
		rightProduct *= nums[i]
	}
	return finalProduct
}
