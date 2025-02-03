package main

import (
	"fmt"
)

func equalPairs(grid [][]int) int {
	n := len(grid)
	rowMap := make(map[string]int)

	for _, row := range grid {
		key := fmt.Sprint(row)
		rowMap[key]++
	}

	count := 0
	for j := 0; j < n; j++ {
		var col []int
		for i := 0; i < n; i++ {
			col = append(col, grid[i][j])
		}
		key := fmt.Sprint(col)
		count += rowMap[key]
	}

	return count
}

func main() {
	fmt.Println(equalPairs([][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}}))
}
