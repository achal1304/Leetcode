package main

import (
	"container/heap"
	"fmt"
)

// REVISIT : Medium

// Definition for a MaxHeap.
type MaxHeap []struct {
	Char  byte
	Count int
}

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].Count > h[j].Count } // Max heap based on Count
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(struct {
		Char  byte
		Count int
	}))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

// Function to rearrange the string
func reorganizeString(s string) string {
	// Step 1: Count the frequency of each character
	counts := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		counts[s[i]]++
	}

	// Step 2: Initialize the max heap with character frequencies
	h := &MaxHeap{}
	for char, count := range counts {
		heap.Push(h, struct {
			Char  byte
			Count int
		}{char, count})
	}

	// Step 3: Reorganize the string
	result := []byte{}
	var prev struct {
		Char  byte
		Count int
	}

	for h.Len() > 0 {
		// Get the character with the highest frequency
		current := heap.Pop(h).(struct {
			Char  byte
			Count int
		})

		// If previous character still has remaining count, push it back to heap
		if prev.Count > 0 {
			heap.Push(h, prev)
		}

		// Add the current character to the result
		result = append(result, current.Char)

		// Decrease the count of the current character
		current.Count--

		// Update the previous character
		prev = current
	}

	// If the length of result matches the input string length, return the result
	if len(result) == len(s) {
		return string(result)
	}
	// Otherwise, it's impossible to rearrange the string
	return ""
}

func main() {
	// Example 1
	s1 := "aab"
	fmt.Println(reorganizeString(s1)) // Output: "aba" or "baa"

	// Example 2
	s2 := "aaab"
	fmt.Println(reorganizeString(s2)) // Output: ""

	// Example 3
	s3 := "aabbcc"
	fmt.Println(reorganizeString(s3)) // Output: "abcabc" or any valid rearrangement
}
