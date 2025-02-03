package main

import "fmt"

func main() {
	fmt.Println(findDifference([]int{1, 2, 3}, []int{2, 4, 6}))
}

func findDifference(nums1 []int, nums2 []int) [][]int {
	dict1 := make(map[int]bool, len(nums1))
	dict2 := make(map[int]bool, len(nums2))
	visited := make(map[int]bool)

	for i := 0; i < len(nums1); i++ {
		dict1[nums1[i]] = true
	}

	unique2 := []int{}
	for i := 0; i < len(nums2); i++ {
		dict2[nums2[i]] = true
		if !dict1[nums2[i]] && !visited[nums2[i]] {
			unique2 = append(unique2, nums2[i])
			visited[nums2[i]] = true
		}
	}

	unique1 := []int{}
	for i := 0; i < len(nums1); i++ {
		if !dict2[nums1[i]] && !visited[nums1[i]] {
			unique1 = append(unique1, nums1[i])
			visited[nums1[i]] = true
		}
	}

	return [][]int{unique1, unique2}
}
