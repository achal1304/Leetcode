package main

import "fmt"

func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
	// fmt.Println(dailyTemperatures([]int{30, 40, 50, 60}))
}

func dailyTemperatures(temperatures []int) []int {
	if len(temperatures) == 0 {
		return []int{}
	}

	ans := make([]int, len(temperatures))
	stack := []int{}
	stack = append(stack, 0)

	for i := 1; i < len(temperatures); i++ {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			ele := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[ele] = i - ele
		}
		stack = append(stack, i)
	}

	return ans
}
