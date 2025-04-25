package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
}

func longestConsecutive(nums []int) int {
	dict := make(map[int]bool)

	for _, ele := range nums {
		dict[ele] = true
	}

	maxLen := 0
	for num := range dict {
		if _, exisits := dict[num-1]; !exisits {
			currLen := 1

			for dict[num+1] {
				currLen += 1
				num += 1
			}

			maxLen = max(maxLen, currLen)
		}
	}
	return maxLen
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b

}
