package main

import "fmt"

func main() {
	fmt.Println(spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
}

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	top, down, left, right := 0, m-1, 0, n-1
	result := []int{}
	for top <= down && left <= right {
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}
		top++

		for i := top; i <= down; i++ {
			result = append(result, matrix[i][right])
		}
		right--

		if top <= down {
			for i := right; i >= left; i-- {
				result = append(result, matrix[down][i])
			}
		}
		down--

		if left <= right {
			for i := down; i >= top; i-- {
				result = append(result, matrix[i][left])
			}
		}
		left++
	}
	return result
}
