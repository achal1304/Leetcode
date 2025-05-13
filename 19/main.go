package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	fast := dummy
	slow := dummy

	// Move fast n+1 steps ahead to maintain a gap of n between slow and fast
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}

	// Move both pointers until fast reaches end
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	// Remove nth node from end
	slow.Next = slow.Next.Next

	return dummy.Next
}
