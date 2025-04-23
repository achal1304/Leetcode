package main

import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {
	indexMap := make(map[int]int) // Map to store the last index of each number
	for i, num := range nums {
		if lastIndex, exists := indexMap[num]; exists {
			if i-lastIndex <= k {
				return true
			}
		}
		indexMap[num] = i
	}

	return false
}

func main() {
	// Test cases
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1}, 3))          // Expected output: true
	fmt.Println(containsNearbyDuplicate([]int{1, 0, 1, 1}, 1))          // Expected output: true
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3, 2}, 4)) // Expected output: false
}
