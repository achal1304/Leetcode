package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findMinArrowShots([][]int{{1, 2},
		{3, 4}, {5, 6}, {7, 12}}))
	// fmt.Println(findMinArrowShots([][]int{{1, 4},
	// 	{2, 4}, {5, 7}, {7, 9}, {9, 11},
	// 	{12, 18}, {17, 18}, {4, 11}, {11, 13}, {7, 12}}))
}

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	arrows := 1
	prevEnd := points[0][1]
	for i := 1; i < len(points); i++ {
		if points[i][0] <= prevEnd {
			continue
		} else {
			prevEnd = points[i][1]
			arrows++
		}
	}

	return arrows
}
