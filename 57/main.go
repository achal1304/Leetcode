package main

import "fmt"

// REVISIT

func main() {
	fmt.Println(insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}))
}

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}
	result := [][]int{}
	i := 0

	for i < len(intervals) && intervals[i][1] <= newInterval[0] {
		result = append(result, intervals[i])
		i++
	}

	// Step 2: Merge overlapping intervals
	for i < len(intervals) && intervals[i][0] <= newInterval[1] {
		newInterval[0] = min(newInterval[0], intervals[i][0])
		newInterval[1] = max(newInterval[1], intervals[i][1])
		i++
	}

	// Add the merged new interval
	result = append(result, newInterval)

	// Step 3: Add the remaining intervals
	for i < len(intervals) {
		result = append(result, intervals[i])
		i++
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
