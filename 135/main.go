package main

import "fmt"

// REVISIT
func main() {
	fmt.Println(candy([]int{1, 2, 3, 5, 4, 3, 2, 1, 8, 4, 3, 2, 5, 6, 7}))
}

func candy(ratings []int) int {
	if len(ratings) == 1 {
		return 1
	}
	if len(ratings) == 0 {
		return 0
	}

	candyAlloc := make([]int, len(ratings))
	candyAlloc[0] = 1
	for i := 1; i < len(candyAlloc); i++ {
		if ratings[i] > ratings[i-1] {
			candyAlloc[i] = candyAlloc[i-1] + 1
		} else {
			candyAlloc[i] = 1
		}
	}

	total := candyAlloc[len(candyAlloc)-1]
	for i := len(candyAlloc) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			if candyAlloc[i] <= candyAlloc[i+1] {
				candyAlloc[i] = candyAlloc[i+1] + 1
			}
		}
		total += candyAlloc[i]
	}

	fmt.Println(candyAlloc)

	return total
}
