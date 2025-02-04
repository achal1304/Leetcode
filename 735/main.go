package main

import "fmt"

func main() {
	fmt.Println(asteroidCollision([]int{-2, -2, 1, -2}))
}

type Stack struct {
	items []int
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	lastIndex := len(s.items) - 1
	item := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return item, true
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func asteroidCollision(asteroids []int) []int {
	stack := &Stack{}

	for _, ele := range asteroids {
		collision := false

		for len(stack.items) > 0 && ele < 0 && stack.items[len(stack.items)-1] > 0 {
			top := stack.items[len(stack.items)-1]

			if abs(ele) > abs(top) {
				stack.Pop()
				continue
			} else if abs(ele) == abs(top) {
				stack.Pop()
				collision = true
				break
			} else {
				collision = true
				break
			}
		}

		if !collision {
			stack.Push(ele)
		}
	}

	return stack.items
}
