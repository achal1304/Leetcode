package main

import "fmt"

// Function to check if the parentheses are valid
func isValid(s string) bool {
	// Stack to store opening parentheses
	stack := []rune{}

	// Map to match closing parentheses with opening parentheses
	matchingParentheses := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	// Iterate over each character in the string
	for _, char := range s {
		// If the character is an opening bracket, push it onto the stack
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
		} else if char == ')' || char == '}' || char == ']' {
			// If the character is a closing bracket
			if len(stack) == 0 {
				return false // No opening bracket to match with
			}
			top := stack[len(stack)-1]   // Peek the top element
			stack = stack[:len(stack)-1] // Pop the top element

			// Check if the top of the stack matches the expected opening bracket
			if top != matchingParentheses[char] {
				return false
			}
		}
	}

	// If the stack is empty, all opening brackets have been matched
	return len(stack) == 0
}

func main() {
	// Test cases
	fmt.Println(isValid("()"))     // Expected output: true
	fmt.Println(isValid("()[]{}")) // Expected output: true
	fmt.Println(isValid("(]"))     // Expected output: false
	fmt.Println(isValid("([)]"))   // Expected output: false
	fmt.Println(isValid("{[]}"))   // Expected output: true
	fmt.Println(isValid(""))       // Expected output: true (empty string is valid)
}
