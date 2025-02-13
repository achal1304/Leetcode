package main

import (
	"container/heap"
	"fmt"
)

type Worker struct {
	cost  int
	index int
}

type MinHeap []Worker

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool {
	if h[i].cost == h[j].cost {
		return h[i].index < h[j].index
	}
	return h[i].cost < h[j].cost
}
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Worker))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func totalCost(costs []int, k int, candidates int) int64 {
	n := len(costs)
	leftHeap, rightHeap := &MinHeap{}, &MinHeap{}
	heap.Init(leftHeap)
	heap.Init(rightHeap)

	left, right := 0, n-1
	for left < candidates && left <= right {
		heap.Push(leftHeap, Worker{costs[left], left})
		left++
	}
	for right >= n-candidates && right >= left {
		heap.Push(rightHeap, Worker{costs[right], right})
		right--
	}

	totalCost := 0

	for i := 0; i < k; i++ {
		var chosen Worker
		if leftHeap.Len() > 0 && rightHeap.Len() > 0 {
			if (*leftHeap)[0].cost <= (*rightHeap)[0].cost {
				chosen = heap.Pop(leftHeap).(Worker)
			} else {
				chosen = heap.Pop(rightHeap).(Worker)
			}
		} else if leftHeap.Len() > 0 {
			chosen = heap.Pop(leftHeap).(Worker)
		} else {
			chosen = heap.Pop(rightHeap).(Worker)
		}

		totalCost += chosen.cost

		if left <= right {
			if chosen.index < left {
				heap.Push(leftHeap, Worker{costs[left], left})
				left++
			} else {
				heap.Push(rightHeap, Worker{costs[right], right})
				right--
			}
		}
	}

	return int64(totalCost)
}

func main() {
	costs := []int{3, 2, 7, 7, 1, 2}
	k := 3
	candidates := 2

	fmt.Println("Total Hiring Cost:", totalCost(costs, k, candidates))
}
