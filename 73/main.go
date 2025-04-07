package main

import "fmt"

func setZeroes(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])

	// Flags to track if the first row and first column should be zeroed
	zeroFirstRow := false
	zeroFirstCol := false

	// Step 1: Check if the first row needs to be zeroed
	for j := 0; j < n; j++ {
		if matrix[0][j] == 0 {
			zeroFirstRow = true
			break
		}
	}

	// Step 2: Check if the first column needs to be zeroed
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			zeroFirstCol = true
			break
		}
	}

	// Step 3: Use the first row and column to mark which rows and columns need to be zeroed
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0 // Mark the first column
				matrix[0][j] = 0 // Mark the first row
			}
		}
	}

	// Step 4: Zero out the cells based on the markers in the first row and first column
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	// Step 5: Zero out the first row if needed
	if zeroFirstRow {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}

	// Step 6: Zero out the first column if needed
	if zeroFirstCol {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}

func main() {
	matrix := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}

	fmt.Println("Original matrix:")
	for _, row := range matrix {
		fmt.Println(row)
	}

	setZeroes(matrix)

	fmt.Println("\nMatrix after setZeroes:")
	for _, row := range matrix {
		fmt.Println(row)
	}
}
