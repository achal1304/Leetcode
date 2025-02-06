package main

import "fmt"

func main() {
	// head := CreateLinkedList([]int{1, 2, 3, 4, 5, 6, 7})
	head := CreateLinkedList([]int{1, 2, 3, 4, 5})
	PrintLinkedList(head)
	oddEvenList(head)
	PrintLinkedList(head)

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

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil || head.Next.Next == nil {
		return head
	}

	count := 1
	curr := head
	for curr.Next != nil {
		curr = curr.Next
		count++
	}
	last := curr
	curr = head

	i := 0
	for i < (count / 2) {
		temp := curr.Next
		curr.Next = curr.Next.Next
		last.Next = temp
		last.Next.Next = nil
		curr = curr.Next
		last = last.Next
		i++
	}

	return head

}
