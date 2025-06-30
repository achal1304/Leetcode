package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	minDiff := math.MaxInt32
	var prev *int

	var inOrder func(node *TreeNode)
	inOrder = func(node *TreeNode) {
		if node == nil {
			return
		}

		inOrder(node.Left)

		if prev != nil {
			diff := abs(node.Val - *prev)
			if diff < minDiff {
				minDiff = diff
			}
		}
		val := node.Val
		prev = &val

		inOrder(node.Right)
	}

	inOrder(root)
	return minDiff
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
