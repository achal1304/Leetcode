package main

import (
	"fmt"
)

func main() {
	fmt.Println(combinationSum3(4, 24))
}

func combinationSum3(k int, n int) [][]int {
	var backtrack func(num, count int, currSelected []int, currSum int)
	result := [][]int{}
	backtrack = func(iValue, count int, currSelected []int, currSum int) {
		fmt.Println("ivalue ", iValue, "count ", count, "currseelcted", currSelected, "currSum ", currSum)
		if count == k && currSum == n {
			cpy := make([]int, len(currSelected))
			copy(cpy, currSelected)
			result = append(result, cpy)
			currSelected = currSelected[:len(currSelected)-1]
			fmt.Println("adding to list  ", currSelected, result)
			return
		}

		if currSum > n {
			currSelected = currSelected[:len(currSelected)-1]
			return
		}

		for i := iValue; i <= 9; i++ {
			backtrack(i+1, count+1, append(currSelected, i), currSum+i)
		}
	}

	backtrack(1, 0, []int{}, 0)
	return result
}
