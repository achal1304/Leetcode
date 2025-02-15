package main

import (
	"fmt"
)

func findPeakElement(nums []int) int {
	low, high := 0, len(nums)-1

	for low < high {
		mid := low + (high-low)/2

		if nums[mid] < nums[mid+1] {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return low
}

func main() {
	arr1 := []int{1, 2, 3, 1}
	arr2 := []int{1, 2, 1, 3, 5, 6, 4}
	arr3 := []int{1, 2, 3, 4, 5}

	fmt.Println("Peak Index (arr1):", findPeakElement(arr1))
	fmt.Println("Peak Index (arr2):", findPeakElement(arr2))
	fmt.Println("Peak Index (arr3):", findPeakElement(arr3))
}
