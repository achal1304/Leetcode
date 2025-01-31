package main

import (
	"fmt"
)

func main() {
	// fmt.Println(compress([]byte{'a', 'a', 'a', 'b', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c'}))
	// fmt.Println(compress([]byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}))
	fmt.Println(compress([]byte{'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'a'}))
	// fmt.Println(compress([]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}))
	// fmt.Println(compress([]byte{'a'}))
	// fmt.Println(compress([]byte{'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'b', 'c', 'c',
	// 	'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c'}))
}

func compress(chars []byte) int {
	write := 0
	read := 0
	n := len(chars)

	for read < n {
		char := chars[read]
		count := 0

		for read < n && chars[read] == char {
			read++
			count++
		}

		chars[write] = char
		write++

		if count > 1 {
			for _, digit := range fmt.Sprintf("%d", count) {
				chars[write] = byte(digit)
				write++
			}
		}
	}

	chars = chars[:write]
	fmt.Println(string(chars))
	return len(chars)
}
