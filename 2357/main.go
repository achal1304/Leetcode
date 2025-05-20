package main

import "fmt"

func minimumOperations(nums []int) int {
	unique := make(map[int]bool)

	// Add all non-zero positive elements to the set
	for _, num := range nums {
		if num > 0 {
			unique[num] = true
		}
	}

	// The number of distinct positive elements is the number of operations required
	return len(unique)
}

func main() {
	// Example 1
	nums1 := []int{1, 5, 0, 3, 5}
	fmt.Println(minimumOperations(nums1)) // Output: 3

	// Example 2
	nums2 := []int{0}
	fmt.Println(minimumOperations(nums2)) // Output: 0

	// Example 3
	nums3 := []int{0, 0, 0}
	fmt.Println(minimumOperations(nums3)) // Output: 0
}
