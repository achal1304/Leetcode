package main

import "fmt"

// REVISIT
func main() {
	fmt.Println(numberOfWays("11100"))
}

func numberOfWays(s string) int64 {
	var countOne, countZero int64 = 0, 0

	for _, ele := range s {
		if ele == '1' {
			countOne++
		} else {
			countZero++
		}
	}

	var ans int64 = 0
	var numZero, numOne int64 = 0, 0
	for _, ele := range s {
		if ele == '0' {
			numZero++
			ans += (numOne * (countOne - numOne))
		} else {
			numOne++
			ans += (numZero * (countZero - numZero))
		}
	}

	return ans

}
