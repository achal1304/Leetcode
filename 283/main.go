package main

import (
	"fmt"
	"os"
)

func main() {
	// moveZeroes([]int{0, 1, 0, 3, 12})
	// moveZeroes([]int{0})
	// moveZeroes([]int{2, 1})
	moveZeroes([]int{1, 2, 0, 0, 2, 3, 40, 0, 0, 0, 0, 50, 0, 23, 0, 0, 21, 0, 0, 0, 0, 0, 2})
}

// SOLUTION 1 - Inefficient
// func moveZeroes(nums []int) {
// 	n := len(nums)
// 	for i := 0; i < n; i++ {
// 		if nums[i] == 0 {
// 			j := i + 1
// 			for j < n && nums[j] == 0 {
// 				j++
// 			}

// 			if j > n-1 {
// 				fmt.Fprint(os.Stdout, nums)
// 				return
// 			}

//				nums[i], nums[j] = nums[j], nums[i]
//			}
//		}
//		fmt.Fprint(os.Stdout, nums)
//		return
//	}

// SOLUTION 2 : Efficient
func moveZeroes(nums []int) {
	n := len(nums)
	lastSwap := 0
	for i := 0; i < n; i++ {
		if nums[i] != 0 && i != lastSwap {
			nums[lastSwap], nums[i] = nums[i], nums[lastSwap]
			lastSwap++
		} else if nums[i] != 0 && i == lastSwap {
			lastSwap++
		}
	}
	fmt.Fprint(os.Stdout, nums)
	return
}
