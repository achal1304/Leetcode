package main

import "fmt"

func main() {
	fmt.Println(isSubsequence("abc", "ahbgdc"))
}

func isSubsequence(s string, t string) bool {
	subSeqI := 0
	for i := 0; i < len(t); i++ {
		if subSeqI > len(s)-1 {
			return true
		}

		if t[i] == s[subSeqI] {
			subSeqI++
		}
	}
	return subSeqI > len(s)-1
}
