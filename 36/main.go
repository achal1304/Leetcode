package main

// REVISIT
func main() {

}

func isValidSudoku(board [][]byte) bool {
	rowSet := make([]map[byte]bool, 9)
	colSet := make([]map[byte]bool, 9)
	qaudSet := make([]map[byte]bool, 9)

	for i := 0; i < 9; i++ {
		rowSet[i] = make(map[byte]bool)
		qaudSet[i] = make(map[byte]bool)
		colSet[i] = make(map[byte]bool)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			num := board[i][j]
			if board[i][j] == '.' {
				continue
			}

			if rowSet[i][num] {
				return false
			}
			rowSet[i][num] = true

			if colSet[j][num] {
				return false
			}
			colSet[j][num] = true

			quad := (i/3)*3 + (j / 3)
			if qaudSet[quad][num] {
				return false
			}
			qaudSet[quad][num] = true
		}
	}

	return true
}
