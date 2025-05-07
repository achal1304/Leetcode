package main

import (
	"fmt"
	"math"
)

type MinStack struct {
	stack    []int
	minStack []int
}

// Constructor
func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{math.MaxInt64}, // initialize with max value
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)

	// Push the new minimum
	min := this.minStack[len(this.minStack)-1]
	if val < min {
		this.minStack = append(this.minStack, val)
	} else {
		this.minStack = append(this.minStack, min)
	}
}

func (this *MinStack) Pop() {
	if len(this.stack) > 0 {
		this.stack = this.stack[:len(this.stack)-1]
		this.minStack = this.minStack[:len(this.minStack)-1]
	}
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

// Test
func main() {
	ms := Constructor()
	ms.Push(-2)
	ms.Push(0)
	ms.Push(-3)
	fmt.Println(ms.GetMin()) // -3
	ms.Pop()
	fmt.Println(ms.Top())    // 0
	fmt.Println(ms.GetMin()) // -2
}
