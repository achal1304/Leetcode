package main

func solve(board [][]byte) {
	for i := 0; i < len(board); i++ {
		if board[i][0] == 'O' {
			board[i][0] = 'T'
			visitBoard(board, []int{i, 0})
		}
		yLen := len(board[0]) - 1
		if board[i][yLen] == 'O' {
			board[i][yLen] = 'T'
			visitBoard(board, []int{i, yLen})
		}
	}

	for i := 0; i < len(board[0]); i++ {
		if board[0][i] == 'O' {
			board[0][i] = 'T'
			visitBoard(board, []int{0, i})
		}
		xLen := len(board) - 1
		if board[xLen][i] == 'O' {
			board[xLen][i] = 'T'
			visitBoard(board, []int{xLen, i})
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == 'T' {
				board[i][j] = 'O'
			}
		}
	}
}

var directions = [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func visitBoard(grid [][]byte, pos []int) {
	for _, dir := range directions {
		nextX := pos[0] + dir[0]
		nextY := pos[1] + dir[1]
		if nextX >= 0 && nextX < len(grid) && nextY >= 0 && nextY < len(grid[0]) && grid[nextX][nextY] == 'O' {
			grid[nextX][nextY] = 'T'
			visitBoard(grid, []int{nextX, nextY})
		}
	}
}
