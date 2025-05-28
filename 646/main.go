package main

import "sort"

func main() {

}

func findLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] < pairs[j][1]
	})

	count := 1
	prevPair := pairs[0]
	for i := 1; i < len(pairs); i++ {
		if prevPair[1] < pairs[i][0] {
			count++
			prevPair = pairs[i]
		}
	}

	return count
}
