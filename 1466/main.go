package main

import "fmt"

// REVISIT
func main() {
	fmt.Println(minReorder(6, [][]int{{0, 1}, {1, 3}, {2, 3}, {4, 0}, {4, 5}}))
}

func minReorder(n int, connections [][]int) int {
	paths := make(map[int][][]int)
	for _, ele := range connections {
		paths[ele[0]] = append(paths[ele[0]], []int{ele[1], 1})
		paths[ele[1]] = append(paths[ele[1]], []int{ele[0], 0})
	}

	var dfs func(city, parent int) int
	dfs = func(city, parent int) int {
		shifts := 0
		for _, neighbour := range paths[city] {
			nextCity, shiftCount := neighbour[0], neighbour[1]
			if nextCity == parent {
				continue
			}
			shifts += shiftCount
			shifts += dfs(nextCity, city)
		}
		return shifts
	}

	return dfs(0, -1)
}
