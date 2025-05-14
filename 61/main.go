package main

import "fmt"

func main() {
	head := &ListNode{Val: 1}
	curr := head
	curr.Next = &ListNode{Val: 2}
	curr = curr.Next
	curr.Next = &ListNode{Val: 3}
	curr = curr.Next
	curr.Next = &ListNode{Val: 4}
	curr = curr.Next
	curr.Next = &ListNode{Val: 5}
	ans := rotateRight(head, 7)
	for ans != nil {
		fmt.Println(ans.Val)
		ans = ans.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}
	dummy := &ListNode{Next: head}
	prev := dummy

	count := 0
	curr := head
	for curr != nil {
		count++
		curr = curr.Next
	}

	if k > count {
		k = k % count
		if k == 0 {
			return head
		}
	}

	fast := head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}

	if fast == nil {
		return head
	}

	slow := head
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	tmp := slow.Next
	slow.Next = nil
	prev.Next = tmp
	fast.Next = head

	return dummy.Next

}
