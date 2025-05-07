package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
}

func evalRPN(tokens []string) int {
	stack := []string{}
	dict := map[string]bool{
		"+": true,
		"-": true,
		"*": true,
		"/": true,
	}

	for _, ele := range tokens {
		if dict[ele] {
			num1, _ := strconv.Atoi(stack[len(stack)-2])
			num2, _ := strconv.Atoi(stack[len(stack)-1])
			stack = stack[:len(stack)-2]
			ansNum := 0

			switch ele {
			case "+":
				ansNum = num1 + num2
			case "-":
				ansNum = num1 - num2
			case "*":
				ansNum = num1 * num2
			case "/":
				ansNum = num1 / num2
			}

			stack = append(stack, fmt.Sprint(ansNum))
		} else {
			stack = append(stack, ele)
		}
	}

	ans, _ := strconv.Atoi(stack[0])
	return ans
}
