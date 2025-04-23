package main

import "fmt"

// Function to calculate the sum of squares of digits of a number
func sumOfSquares(n int) int {
	sum := 0
	for n > 0 {
		digit := n % 10
		sum += digit * digit
		n /= 10
	}
	return sum
}

// Function to determine if a number is a happy number
func isHappy(n int) bool {
	seen := make(map[int]bool) // To track numbers we've already seen

	// Keep replacing the number with the sum of squares of its digits
	for n != 1 {
		n = sumOfSquares(n)

		// If we've seen this number before, it's in a cycle
		if seen[n] {
			return false
		}

		// Mark this number as seen
		seen[n] = true
	}

	// If the loop exits with n == 1, it's a happy number
	return true
}

func main() {
	// Test cases
	fmt.Println(isHappy(19)) // Expected output: true
	fmt.Println(isHappy(2))  // Expected output: false
}
