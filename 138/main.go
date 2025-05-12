package main

// Definition for a Node.
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	// Step 1: Copy all nodes and store mapping
	oldToNew := make(map[*Node]*Node)

	curr := head
	for curr != nil {
		oldToNew[curr] = &Node{Val: curr.Val}
		curr = curr.Next
	}

	// Step 2: Assign Next and Random pointers
	curr = head
	for curr != nil {
		copyNode := oldToNew[curr]
		copyNode.Next = oldToNew[curr.Next]
		copyNode.Random = oldToNew[curr.Random]
		curr = curr.Next
	}

	return oldToNew[head]
}
