package main

import "fmt"

func main() {
	fmt.Println(canVisitAllRooms([][]int{{1, 3}, {3, 0, 1}, {2}, {0}}))
}

func canVisitAllRooms(rooms [][]int) bool {
	if len(rooms) == 0 {
		return true
	}

	if len(rooms) == 1 {
		return true
	}

	if len(rooms) > 1 && len(rooms[0]) == 0 {
		return false
	}

	stack := []int{}
	for _, ele := range rooms[0] {
		stack = append(stack, ele)
	}
	visited := make(map[int]bool)
	visited[0] = true

	for len(stack) > 0 {
		ele := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if visited[ele] {
			continue
		}

		for _, key := range rooms[ele] {
			stack = append(stack, key)
		}

		visited[ele] = true
		if len(visited) == len(rooms) {
			return true
		}
	}

	return false
}
