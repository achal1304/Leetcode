package main

import (
	"fmt"
	"math"
)

func main() {
	piles1 := []int{3, 6, 7, 11}
	h1 := 8
	fmt.Println("Minimum speed for test case 1:", minEatingSpeed(piles1, h1))

	piles2 := []int{30, 11, 23, 4, 20}
	h2 := 5
	fmt.Println("Minimum speed for test case 2:", minEatingSpeed(piles2, h2))

	piles3 := []int{30, 11, 23, 4, 20}
	h3 := 6
	fmt.Println("Minimum speed for test case 3:", minEatingSpeed(piles3, h3))
}

func minEatingSpeed(piles []int, h int) int {
	timeRequired := func(speed int) int {
		totalHours := 0
		for _, bananas := range piles {
			totalHours += int(math.Ceil(float64(bananas) / float64(speed)))
		}
		return totalHours
	}

	low, high := 1, max(piles)
	for low < high {
		mid := low + (high-low)/2
		if timeRequired(mid) <= h {
			high = mid
		} else {
			low = mid + 1
		}
	}

	return low
}

func max(arr []int) int {
	maxVal := arr[0]
	for _, v := range arr {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}
