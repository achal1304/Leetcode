package main

// REVISIT
func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	if root.Left != nil {
		flatten(root.Left)

		tmpRight := root.Right

		root.Right = root.Left
		root.Left = nil

		current := root.Right
		for current.Right != nil {
			current = current.Right
		}

		current.Right = tmpRight
	}

	flatten(root.Right)
}
