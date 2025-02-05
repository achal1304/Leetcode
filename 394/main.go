package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(decodeString("2[abc]3[cd]ef"))
}

func decodeString(s string) string {
	currStack := []string{}
	numStack := []int{}

	currString := ""
	currNum := 0

	for _, ele := range s {
		if unicode.IsDigit(ele) {
			currNum = currNum*10 + int(ele-'0')
		} else if ele == '[' {
			numStack = append(numStack, currNum)
			currStack = append(currStack, currString)
			currString = ""
			currNum = 0
		} else if ele == ']' {
			num := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]
			lastEle := ""
			if len(currStack) > 0 {
				lastEle = currStack[len(currStack)-1]
				currStack = currStack[:len(currStack)-1]
			}
			decoded := ""
			for i := 0; i < num; i++ {
				decoded += currString
			}
			currString = lastEle + decoded
		} else {
			currString += string(ele)
		}
	}
	return currString
}
