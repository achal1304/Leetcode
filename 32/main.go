package main

func main() {

}

// REVISIT
func longestValidParentheses(s string) int {
	stack := []int{-1}
	if len(s) <= 1 {
		return 0
	}

	maxCount := 0

	for i, ele := range s {
		if ele == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				currLen := i - stack[len(stack)-1]
				if currLen > maxCount {
					maxCount = currLen
				}
			}
		}
	}

	return maxCount
}
