package main

// REVISIT - shift rows and columns first and then swap the row from first to last ele
func main() {
	rotate([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
}

func rotate(matrix [][]int) {
	n := len(matrix)

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-j-1] = matrix[i][n-j-1], matrix[i][j]
		}
	}
}
