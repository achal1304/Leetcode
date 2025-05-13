package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	prevGroup := dummy

	for {
		kth := getKthNode(prevGroup, k)
		if kth == nil {
			break
		}

		prev := kth.Next
		nextGroup := kth.Next
		curr := prevGroup.Next

		for curr != nextGroup {
			tmp := curr.Next
			curr.Next = prev
			prev = curr
			curr = tmp
		}

		tmp := prevGroup.Next
		prevGroup.Next = kth
		prevGroup = tmp

	}
	return dummy.Next
}

func getKthNode(curr *ListNode, k int) *ListNode {
	for curr != nil && k > 0 {
		curr = curr.Next
		k--
	}
	return curr
}
