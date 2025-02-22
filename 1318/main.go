package main

import (
	"fmt"
)

func main() {
	fmt.Println(minFlips(8, 3, 5))
}

func minFlips(a int, b int, c int) int {
	flips := 0

	for i := 0; i < 32; i++ {
		bitA := (a >> i) & 1
		bitB := (b >> i) & 1
		bitC := (c >> i) & 1

		if bitC == 1 {
			if bitA == 0 && bitB == 0 {
				flips++
			}
		} else {
			flips += bitA + bitB
		}
	}

	return flips
}
