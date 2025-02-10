package main

import "fmt"

func main() {
	fmt.Println(findCircleNum([][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}))
}

func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	paths := make(map[int][]int)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}

			if isConnected[i][j] == 1 {
				paths[i] = append(paths[i], j)
			}
		}
	}

	count := 0
	visited := make(map[int]bool)
	for i := 0; i < n; i++ {
		if _, ok := paths[i]; !ok {
			count++
			continue
		}

		if !visited[i] {
			count++
			visitCity(visited, paths, i)
		}
	}
	return count
}

// Approach 1 : For loop
func visitCity(visited map[int]bool, paths map[int][]int, ele int) {
	stack := []int{ele}
	for len(stack) > 0 {
		ele := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if visited[ele] {
			continue
		}

		for _, key := range paths[ele] {
			stack = append(stack, key)
		}

		visited[ele] = true
	}
}

// // Approach 2 - recursive
// func visitCity(visited map[int]bool, paths map[int][]int, city int) {
// 	visited[city] = true
// 	for _, newCity := range paths[city] {
// 		if !visited[newCity] {
// 			visitCity(visited, paths, newCity)
// 		}
// 	}
// }
