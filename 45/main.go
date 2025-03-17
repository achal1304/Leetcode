package main

import "fmt"

func main() {
	fmt.Println(jump([]int{1, 1, 3, 4, 1, 0, 0, 1}))
}

func jump(nums []int) int {
	n := len(nums)
	if n == 1 || n == 0 {
		return 0
	}

	currEnd := 0
	jumpCount := 0
	maxDistance := 0
	for i := 0; i < n-1; i++ {
		maxDistance = max(maxDistance, i+nums[i])

		fmt.Println("outside ", i, currEnd, maxDistance)
		if i == currEnd {
			fmt.Println("inside ", i, currEnd, maxDistance)
			currEnd = maxDistance
			jumpCount++
			if currEnd >= n-1 {
				return jumpCount
			}
		}
	}

	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
