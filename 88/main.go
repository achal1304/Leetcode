package main

import "fmt"

func main() {
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 4, 6}, 3)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	currIndex := len(nums1) - 1
	nums1I := m - 1
	nums2I := n - 1

	for currIndex >= 0 && nums1I >= 0 && nums2I >= 0 {
		fmt.Println("nums1", nums1, "nums2I", nums2I, nums1I)
		if nums2[nums2I] > nums1[nums1I] {
			nums1[currIndex] = nums2[nums2I]
			nums2I--
			currIndex--
		} else {
			nums1[currIndex] = nums1[nums1I]
			nums1I--
			currIndex--
		}
	}
	fmt.Println("outside nums1", nums1, currIndex, "nums2I", nums2I, nums1I)

	if nums1I < 0 && nums2I >= 0 {
		for nums2I >= 0 {
			nums1[currIndex] = nums2[nums2I]
			nums2I--
			currIndex--
		}
	} else if nums2I < 0 && nums1I >= 0 {
		for nums1I >= 0 {
			nums1[currIndex] = nums1[nums1I]
			nums1I--
			currIndex--
		}
	}

	fmt.Println(nums1)
}
