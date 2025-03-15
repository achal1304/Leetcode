package main

import "fmt"

func main() {
	fmt.Println(majorityElement([]int{3, 3, 2, 2, 3, 3, 4}))
}

func majorityElement(nums []int) int {
	dic := make(map[int]int)
	maxCount := 0
	maxEle := 0
	for _, ele := range nums {
		dic[ele]++
		if dic[ele] > maxCount {
			maxCount = dic[ele]
			maxEle = ele
		}
	}

	return maxEle
}
