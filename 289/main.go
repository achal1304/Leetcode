package main

// REVISIT - use 2 for live but converted to dead and 3 for dead but converted to live
func main() {
	gameOfLife([][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}})
}

var directions = []struct{ x, y int }{{1, 1}, {1, 0}, {0, 1}, {-1, -1}, {-1, 0}, {0, -1}, {1, -1}, {-1, 1}}

func gameOfLife(board [][]int) {
	m := len(board)
	n := len(board[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			count := countLive(i, j, board, m, n)
			if board[i][j] == 0 && count == 3 {
				board[i][j] = 3
			} else if board[i][j] == 1 {
				if count < 2 || count > 3 {
					board[i][j] = 2
				}
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 3 {
				board[i][j] = 1
			} else if board[i][j] == 2 {
				board[i][j] = 0
			}
		}
	}
}

func countLive(x, y int, board [][]int, m, n int) int {
	liveCount := 0
	for _, dir := range directions {
		nextEleX := x + dir.x
		nextEleY := y + dir.y

		if nextEleX >= 0 && nextEleX < m && nextEleY >= 0 && nextEleY < n {
			if board[nextEleX][nextEleY] == 1 || board[nextEleX][nextEleY] == 2 {
				liveCount++
			}
		}
	}
	return liveCount
}
