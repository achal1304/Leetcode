package main

import (
	"fmt"
	"unicode"
)

// REVISIT - HARD
func main() {
	fmt.Println(calculate("(1+(4+5+2)-3)+(6+8)"))
}

func calculate(s string) int {
	stack := []int{}
	num := 0
	sign := 1
	result := 0

	for i := 0; i < len(s); i++ {
		char := s[i]

		if unicode.IsDigit(rune(char)) {
			num = 0
			for i < len(s) && unicode.IsDigit(rune(s[i])) {
				num = (num * 10) + int(s[i]-'0')
				i++
			}
			i--
			result += num * sign
		} else if char == '+' {
			sign = 1
		} else if char == '-' {
			sign = -1
		} else if char == '(' {
			stack = append(stack, result)
			stack = append(stack, sign)
			result = 0
			sign = 1
		} else if char == ')' {
			prevSign := stack[len(stack)-1]
			prevRes := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			result = prevRes + result*prevSign
		}
	}

	return result
}
