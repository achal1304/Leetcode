package main

import "fmt"

func main() {
	n := 22
	result := guessNumber(n)
	fmt.Printf("The picked number is: %d\n", result)
}
func guess(num int) int {
	pick := 10
	if num == pick {
		return 0
	} else if num > pick {
		return -1
	} else {
		return 1
	}
}

func guessNumber(n int) int {
	low, high := 1, n

	for low <= high {
		mid := low + (high-low)/2
		result := guess(mid)

		if result == 0 {
			return mid
		} else if result == 1 {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}
