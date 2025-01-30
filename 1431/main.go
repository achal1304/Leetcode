package main

import (
	"fmt"
)

func main() {
	fmt.Println(kidsWithCandies([]int{2, 3, 5, 1, 3}, 3))
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	candyBool := make([]bool, len(candies))
	maxCandy := candies[0]
	for _, val := range candies {
		if val > maxCandy {
			maxCandy = val
		}
	}

	for i, val := range candies {
		if val+extraCandies >= maxCandy {
			candyBool[i] = true
		}
	}
	return candyBool
}
