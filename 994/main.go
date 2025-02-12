package main

import "fmt"

func main() {
	// maze := [][]byte{
	// 	{'+', '+', '.', '+'},
	// 	{'.', '.', '.', '+'},
	// 	{'+', '+', '+', '.'},
	// }

	maze := [][]int{
		{2, 1, 1},
		{1, 1, 0},
		{0, 1, 1},
	}

	fmt.Println(orangesRotting(maze))
}

var Directions = [][]int{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}

type Position struct {
	row int
	col int
}
type PathValue struct {
	pos Position
	val int
}

func orangesRotting(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	queue := []PathValue{}
	freshCount := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, PathValue{Position{i, j}, 0})
			} else if grid[i][j] == 1 {
				freshCount++
			}
		}
	}

	minutes := 0
	for len(queue) > 0 {
		r, c, time := queue[0].pos.row, queue[0].pos.col, queue[0].val
		queue = queue[1:]

		minutes = time

		for _, dir := range Directions {
			newR, newC := r+dir[0], c+dir[1]

			if newR >= 0 && newR < m && newC >= 0 && newC < n && grid[newR][newC] == 1 {
				grid[newR][newC] = 2
				queue = append(queue, PathValue{Position{newR, newC}, time + 1})
				freshCount--
			}
		}
	}

	if freshCount > 0 {
		return -1
	}
	return minutes
}
