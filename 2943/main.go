package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println(maximizeSquareHoleArea(3, 2, []int{3, 2, 4}, []int{2, 3}))
}

func maximizeSquareHoleArea(n int, m int, hBars []int, vBars []int) int {
	if len(hBars) == 0 || len(vBars) == 0 {
		return 1
	}
	sort.Ints(hBars)
	sort.Ints(vBars)
	hCount := make(map[int]bool)
	count := 2
	hCount[2] = true
	for i := 1; i < len(hBars); i++ {
		if hBars[i] == hBars[i-1]+1 {
			count++
			hCount[count] = true
		} else {
			count = 2
		}
	}

	vCount := make(map[int]bool)
	count = 2
	vCount[2] = true
	for i := 1; i < len(vBars); i++ {
		if vBars[i] == vBars[i-1]+1 {
			count++
			vCount[count] = true
		} else {
			count = 2
		}
	}
	maxCount := 2
	for k, _ := range hCount {
		if vCount[k] {
			if k > maxCount {
				maxCount = k
			}
		}
	}
	return maxCount * maxCount
}
