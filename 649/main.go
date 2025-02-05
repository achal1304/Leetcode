package main

import (
	"fmt"
)

func main() {
	fmt.Println(predictPartyVictory("RDDDRRRRDRDDDDDRRRD"))
}

func predictPartyVictory(senate string) string {
	rQueue := []int{}
	dQueue := []int{}
	n := len(senate)

	for i, ch := range senate {
		if ch == 'R' {
			rQueue = append(rQueue, i)
		} else {
			dQueue = append(dQueue, i)
		}
	}

	for len(rQueue) > 0 && len(dQueue) > 0 {
		rIndex := rQueue[0]
		dIndex := dQueue[0]
		rQueue = rQueue[1:]
		dQueue = dQueue[1:]

		if rIndex < dIndex {
			rQueue = append(rQueue, rIndex+n)
		} else {
			dQueue = append(dQueue, dIndex+n)
		}
	}

	if len(rQueue) > 0 {
		return "Radiant"
	}
	return "Dire"
}
