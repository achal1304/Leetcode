package main

import "fmt"

type Value struct {
	Den string
	Val float64
}

// REVISIT
func main() {
	equations := [][]string{
		{"a", "b"},
		{"b", "c"},
	}

	values := []float64{2.0, 3.0}

	queries := [][]string{
		{"a", "c"},
		{"b", "a"},
		{"a", "e"},
		{"a", "a"},
		{"x", "x"},
	}

	fmt.Println(calcEquation(equations, values, queries))
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	paths := make(map[string][]Value)
	solved := []float64{}
	for i, ele := range equations {
		paths[ele[0]] = append(paths[ele[0]], Value{ele[1], values[i]})
		paths[ele[1]] = append(paths[ele[1]], Value{ele[0], 1 / values[i]})
	}
	fmt.Println(paths)

	bfs := func(paths map[string][]Value, num, den string) float64 {
		queue := []Value{{num, 1.0}}
		visited := make(map[string]bool)

		for len(queue) > 0 {
			num := queue[0].Den
			currVal := queue[0].Val
			queue = queue[1:]
			for _, ele := range paths[num] {
				if ele.Den == den {
					return currVal * ele.Val
				} else {
					if !visited[ele.Den] {
						queue = append(queue, Value{ele.Den, currVal * ele.Val})
						visited[ele.Den] = true
					}
				}
			}
		}

		return -1.0
	}

	for _, ele := range queries {
		solved = append(solved, bfs(paths, ele[0], ele[1]))
	}

	return solved
}
