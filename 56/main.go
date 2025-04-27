package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}}))
}

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := [][]int{}
	res = append(res, intervals[0])
	for i := 1; i < len(intervals); i++ {
		lastele := res[len(res)-1]
		if intervals[i][0] <= lastele[1] {
			res[len(res)-1][1] = max(lastele[1], intervals[i][1])
		} else {
			res = append(res, intervals[i])
		}
	}
	return res
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
