package main

import (
	"fmt"
)

func main() {
	fmt.Println(gcdOfStrings("ABABABAB", "ABAB"))
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func gcdOfStrings(str1 string, str2 string) string {

	if str1+str2 != str2+str1 {
		return ""
	}

	gcdVal := gcd(len(str1), len(str2))

	return str1[:gcdVal]
}
