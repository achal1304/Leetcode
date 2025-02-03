package main

import (
	"fmt"
)

func uniqueOccurrences(arr []int) bool {
	occurrenceMap := make(map[int]int)

	for _, num := range arr {
		occurrenceMap[num]++
	}

	occurrenceSet := make(map[int]struct{})

	for _, count := range occurrenceMap {
		if _, exists := occurrenceSet[count]; exists {
			return false
		}
		occurrenceSet[count] = struct{}{}
	}

	return true
}

func main() {
	arr3 := []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}
	fmt.Println(uniqueOccurrences(arr3))
}
