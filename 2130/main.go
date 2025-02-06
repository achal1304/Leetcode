package main

import "fmt"

func main() {
	// head := CreateLinkedList([]int{1, 3, 4, 7, 1, 2, 6})
	head := CreateLinkedList([]int{1, 3, 4, 2, 3, 1, 4, 2})
	PrintLinkedList(head)
	fmt.Println(pairSum(head))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateLinkedList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{Val: nums[0]}
	current := head

	for _, num := range nums[1:] {
		current.Next = &ListNode{Val: num}
		current = current.Next
	}

	return head
}

func PrintLinkedList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Print(current.Val)
		if current.Next != nil {
			fmt.Print(" -> ")
		}
		current = current.Next
	}
	fmt.Println()
}

func pairSum(head *ListNode) int {
	if head == nil {
		return 0
	}
	count := 1
	curr := head
	for curr.Next != nil {
		curr = curr.Next
		count++
	}
	m := count / 2
	i := 0
	currMax := 0
	curr = head
	twinArray := make([]int, m)
	for i < count {
		if i < m {
			twinArray[i] = curr.Val
		} else {
			sum := twinArray[count-i-1] + curr.Val
			if currMax < sum {
				currMax = sum
			}
		}
		i++
		curr = curr.Next
	}
	return currMax
}
