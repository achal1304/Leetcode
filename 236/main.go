package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// arr := []int{10, 5, -3, 3, 2, 0, 11, 3, -2, 0, 1}
	arr := []int{3, 5, 1, 6, 2, 0, 8, 0, 0, 7, 4}
	n := len(arr)

	root := InsertLevelOrder(arr, 0, n)
	PrintLevelOrder(root)
	fmt.Println("answer is ")
	fmt.Println(lowestCommonAncestor(root, &TreeNode{Val: 5}, &TreeNode{Val: 1}))
}

func InsertLevelOrder(arr []int, i, n int) *TreeNode {
	if i >= n {
		return nil
	}

	root := &TreeNode{Val: arr[i]}
	root.Left = InsertLevelOrder(arr, 2*i+1, n)
	root.Right = InsertLevelOrder(arr, 2*i+2, n)
	return root
}

func PrintLevelOrder(root *TreeNode) {
	if root == nil {
		return
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		fmt.Print(node.Val, " ")

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
}

// REVISIT

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	}
	return right
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
