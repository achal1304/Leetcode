package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(convert("PAYPALISHIRING", 4))
}

func convert(s string, numRows int) string {
	if len(s) <= 1 || len(s) < numRows || numRows == 1 {
		return s
	}

	res := make([]string, numRows)
	isdown := false
	currRow := 0

	for _, char := range s {
		res[currRow] += string(char)
		if currRow == 0 || currRow == numRows-1 {
			isdown = !isdown
		}
		if isdown {
			currRow++
		} else {
			currRow--
		}
	}

	return strings.Join(res, "")
}
