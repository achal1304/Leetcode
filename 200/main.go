package main

func numIslands(grid [][]byte) int {
	count := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				visitIslands(grid, []int{i, j})
				count++
			}
		}
	}
	return count
}

var directions = [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func visitIslands(grid [][]byte, pos []int) {
	grid[pos[0]][pos[1]] = '0'

	for _, dir := range directions {
		nextX := pos[0] + dir[0]
		nextY := pos[1] + dir[1]
		if nextX >= 0 && nextX < len(grid) && nextY >= 0 && nextY < len(grid[0]) && grid[nextX][nextY] == '1' {
			visitIslands(grid, []int{nextX, nextY})
		}
	}

	return
}
