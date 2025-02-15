package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(successfulPairs([]int{1, 11, 5, 6, 9}, []int{3, 2, 5, 6, 1}, 12))
}
func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
	m := len(potions)

	binarySearch := func(target int64) int {
		low, high := 0, m-1
		for low <= high {
			mid := (high + low) / 2
			if int64(potions[mid]) >= target {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
		return low
	}

	result := make([]int, len(spells))
	for i, spell := range spells {
		target := (success + int64(spell) - 1) / int64(spell)
		if target < int64(potions[0]) {
			result[i] = len(potions)
			continue
		} else if target > int64(potions[len(potions)-1]) {
			result[i] = 0
			continue
		}
		index := binarySearch(target)
		result[i] = m - index
	}

	return result
}
