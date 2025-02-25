package main

import (
	"fmt"
	"sort"
)

// REVISIT
func main() {
	fmt.Println(eraseOverlapIntervals([][]int{{1, 3}, {2, 8}, {3, 7}, {4, 6}, {3, 4}, {6, 8}, {4, 8}}))
}

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 1 || len(intervals) == 0 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	removals := 0
	prevEnd := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < prevEnd {
			removals++
		} else {
			prevEnd = intervals[i][1]
		}
	}

	return removals
}
