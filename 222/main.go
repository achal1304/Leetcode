package main

// REVISIT

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left, right := root, root
	lh, rh := 0, 0

	// Calculate left height
	for left != nil {
		lh++
		left = left.Left
	}

	// Calculate right height
	for right != nil {
		rh++
		right = right.Right
	}

	if lh == rh {
		// It's a perfect tree
		return (1 << lh) - 1
	}
	// Otherwise, recursively count
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}
