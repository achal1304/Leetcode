package main

// REVISIT

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func main() {

}
func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	current := root
	for current != nil {
		dummyNode := &Node{}
		prev := dummyNode

		for node := current; node != nil; node = node.Next {
			if node.Left != nil {
				prev.Next = node.Left
				prev = prev.Next
			}
			if node.Right != nil {
				prev.Next = node.Right
				prev = prev.Next
			}
		}

		current = dummyNode.Next
	}

	return root
}
