package main

import "fmt"

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	romanDict := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	currNum := 0
	for i := 0; i < len(s)-1; i++ {
		if romanDict[s[i]] >= romanDict[s[i+1]] {
			currNum += romanDict[s[i]]
		} else {
			currNum -= romanDict[s[i]]
		}
	}
	currNum += romanDict[s[len(s)-1]]
	return currNum
}
