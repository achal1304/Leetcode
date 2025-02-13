package main

import (
	"container/heap"
	"fmt"
)

func main() {
	obj := Constructor()

	fmt.Println(obj.PopSmallest())
	fmt.Println(obj.PopSmallest())
	obj.AddBack(1)
	fmt.Println(obj.PopSmallest())
	fmt.Println(obj.PopSmallest())
}

type SmallestInfiniteSet struct {
	counter int
	minHeap *MinHeap
	present map[int]bool
}

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func Constructor() SmallestInfiniteSet {
	return SmallestInfiniteSet{
		counter: 1,
		minHeap: &MinHeap{},
		present: make(map[int]bool),
	}
}

func (s *SmallestInfiniteSet) PopSmallest() int {
	if s.minHeap.Len() > 0 {
		smallest := heap.Pop(s.minHeap).(int)
		delete(s.present, smallest)
		return smallest
	} else {
		smallest := s.counter
		s.counter++
		return smallest
	}
}

func (s *SmallestInfiniteSet) AddBack(num int) {
	if num < s.counter && !s.present[num] {
		heap.Push(s.minHeap, num)
		s.present[num] = true
	}
}
