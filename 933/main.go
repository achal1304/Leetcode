package main

import "fmt"

func main() {
	obj := Constructor()
	fmt.Println(obj.Ping(1))
	fmt.Println(obj.Ping(100))
	fmt.Println(obj.Ping(3001))
	fmt.Println(obj.Ping(3002))
}

type RecentCounter struct {
	visitedArray []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		visitedArray: []int{},
	}
}

func (this *RecentCounter) Ping(t int) int {
	this.visitedArray = append(this.visitedArray, t)
	if len(this.visitedArray) == 0 {
		return 1
	}

	rangeS := t - 3000

	i := 0
	for i < len(this.visitedArray) {
		if this.visitedArray[i] >= rangeS {
			break
		} else {
			i++
		}
	}

	this.visitedArray = this.visitedArray[i:len(this.visitedArray)]
	return len(this.visitedArray)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
