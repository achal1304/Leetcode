package main

import "fmt"

func main() {
	fmt.Println(letterCombinations("234"))
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	phone := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	var res []string

	var backtrack func(index int, currString string)
	backtrack = func(index int, currString string) {
		if index >= len(digits) {
			res = append(res, currString)
			return
		}

		keys := phone[digits[index]]
		for _, ele := range keys {
			backtrack(index+1, currString+string(ele))
		}
	}

	backtrack(0, "")
	return res
}
