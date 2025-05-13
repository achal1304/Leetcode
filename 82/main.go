package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy

	for head != nil {
		if head.Next != nil && head.Val == head.Next.Val {
			dupVal := head.Val
			for head != nil && head.Val == dupVal {
				head = head.Next
			}
			prev.Next = head
		} else {
			prev = prev.Next
			head = head.Next
		}
	}
	return dummy.Next
}
