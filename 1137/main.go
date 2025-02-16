package main

import "fmt"

func main() {
	fmt.Println(tribonacci(25))
}

func tribonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else if n == 2 {
		return 1
	}

	start := 0
	mid := 1
	end := 1

	for i := 3; i < n; i++ {
		newEnd := start + mid + end
		start = mid
		mid = end
		end = newEnd
	}
	return start + mid + end
}
