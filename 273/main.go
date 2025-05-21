package main

import "fmt"

func main() {
	fmt.Println(numberToWords(9812))
}

var lessThan20 = []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
var tens = []string{"", "", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
var thousands = []string{"", "Thousand", "Million", "Billion"}

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}
	return helperWords(num)
}

func helperWords(num int) string {
	if num == 0 {
		return ""
	}

	result := ""
	for i := 0; num > 0; i++ {

		if num%1000 != 0 {
			if i == 0 {
				result = helperThree(num % 1000)
			} else {
				if result == "" {
					result = helperThree(num%1000) + " " + thousands[i]
				} else {
					result = helperThree(num%1000) + " " + thousands[i] + " " + result
				}
			}
		}

		fmt.Println("result ", result)

		num = num / 1000
	}

	return result
}

func helperThree(num int) string {
	if num == 0 {
		return ""
	}

	if num < 20 {
		return lessThan20[num]
	} else if num < 100 {
		if num%10 == 0 {
			return tens[num/10]
		} else {

			return tens[num/10] + " " + lessThan20[num%10]
		}
	} else {
		res := helperThree(num % 100)
		if res == "" {
			return lessThan20[num/100] + " Hundred"
		} else {
			return lessThan20[num/100] + " Hundred " + helperThree(num%100)
		}
	}
}
