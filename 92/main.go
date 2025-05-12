package main

import "fmt"

// REVISIT
func main() {
	head := &ListNode{Val: 3}
	curr := head
	curr.Next = &ListNode{Val: 5}
	// curr = curr.Next
	// curr.Next = &ListNode{Val: 3}
	// curr = curr.Next
	// curr.Next = &ListNode{Val: 4}
	// curr = curr.Next
	// curr.Next = &ListNode{Val: 5}
	// curr = curr.Next
	// curr.Next = &ListNode{Val: 6}
	// curr = curr.Next
	// curr.Next = &ListNode{Val: 7}
	head2 := (reverseBetween(head, 1, 2))

	for head2 != nil {
		fmt.Println(head2.Val)
		head2 = head2.Next
	}

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil {
		return nil
	}

	if left == right {
		return head
	}

	dummy := &ListNode{Next: head}
	prev := dummy

	for i := 1; i < left; i++ {
		prev = prev.Next
	}

	curr := prev.Next
	var prevNode *ListNode
	for i := 0; i < (right-left)+1; i++ {
		temp := curr.Next
		curr.Next = prevNode
		prevNode = curr
		curr = temp
	}

	tail := prev.Next
	prev.Next = prevNode
	tail.Next = curr

	return dummy.Next
}
