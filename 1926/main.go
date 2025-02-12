package main

import "fmt"

func main() {
	// maze := [][]byte{
	// 	{'+', '+', '.', '+'},
	// 	{'.', '.', '.', '+'},
	// 	{'+', '+', '+', '.'},
	// }

	maze := [][]byte{
		{'+', '+', '+'},
		{'.', '.', '.'},
		{'+', '+', '+'},
	}

	fmt.Println(nearestExit(maze, []int{1, 0}))
}

var Directions = [][]int{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}

type Position struct {
	row int
	col int
}
type PathValue struct {
	pos Position
	val int
}

func nearestExit(maze [][]byte, entrance []int) int {
	startCol := entrance[1]
	startRow := entrance[0]

	queue := []PathValue{{Position{startRow, startCol}, 0}}
	maze[startRow][startCol] = '+'
	iteration := 0
	for len(queue) > 0 {
		iteration++
		ele := queue[0]
		queue = queue[1:]
		for _, dir := range Directions {
			newRow := ele.pos.row + dir[0]
			newCol := ele.pos.col + dir[1]
			if newRow < len(maze) && newRow >= 0 && newCol < len(maze[0]) && newCol >= 0 && maze[newRow][newCol] == '.' {
				if newRow == 0 || newRow == len(maze)-1 || newCol == 0 || newCol == len(maze[0])-1 {
					return ele.val + 1
				}
				maze[newRow][newCol] = '+'
				queue = append(queue, PathValue{Position{newRow, newCol}, ele.val + 1})
			}
		}
	}

	return -1
}
