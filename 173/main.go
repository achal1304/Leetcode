package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	stack []*TreeNode
}

// Helper: Push all left children to the stack
func (it *BSTIterator) pushLeft(node *TreeNode) {
	for node != nil {
		it.stack = append(it.stack, node)
		node = node.Left
	}
}

// Constructor
func Constructor(root *TreeNode) BSTIterator {
	it := BSTIterator{stack: []*TreeNode{}}
	it.pushLeft(root)
	return it
}

// Returns true if there's a next smallest element
func (it *BSTIterator) HasNext() bool {
	return len(it.stack) > 0
}

// Returns the next smallest element
func (it *BSTIterator) Next() int {
	n := len(it.stack)
	node := it.stack[n-1]
	it.stack = it.stack[:n-1] // pop
	if node.Right != nil {
		it.pushLeft(node.Right)
	}
	return node.Val
}
