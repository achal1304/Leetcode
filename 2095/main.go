package main

import "fmt"

func main() {
	// head := CreateLinkedList([]int{1, 3, 4, 7, 1, 2, 6})
	head := CreateLinkedList([]int{1, 2, 3, 4})
	deleteMiddle(head)
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

func deleteMiddle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slow, fast := head, head
	prev := &ListNode{}

	for fast.Next != nil && fast.Next.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	if fast.Next == nil {
		prev.Next = slow.Next
	} else if fast.Next.Next == nil {
		slow.Next = slow.Next.Next
	}
	return head
}
