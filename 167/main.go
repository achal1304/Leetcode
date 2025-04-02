package main

import "fmt"

func main() {

	fmt.Println(twoSum([]int{2, 3, 4}, 6))
}

func twoSum(numbers []int, target int) []int {
	start := 0
	end := len(numbers) - 1
	for start < end {
		if numbers[start]+numbers[end] == target {
			return []int{start + 1, end + 1}
		} else if numbers[start]+numbers[end] < target {
			start++
		} else if numbers[start]+numbers[end] > target {
			end--
		}
	}

	return []int{start + 1, end + 1}
}
