package main

import "fmt"

func main() {
	fmt.Println(intToRoman(1994))
}

func intToRoman(num int) string {
	romanDict := []struct {
		num   int
		value string
	}{
		{1, "I"},
		{4, "IV"},
		{5, "V"},
		{9, "IX"},
		{10, "X"},
		{40, "XL"},
		{50, "L"},
		{90, "XC"},
		{100, "C"},
		{400, "CD"},
		{500, "D"},
		{900, "CM"},
		{1000, "M"},
	}

	ans := ""
	for maxIndex := len(romanDict) - 1; maxIndex >= 0; maxIndex-- {
		if num == 0 {
			return ans
		}
		for num >= romanDict[maxIndex].num {
			num -= romanDict[maxIndex].num
			ans += romanDict[maxIndex].value
		}
	}

	return ans
}
