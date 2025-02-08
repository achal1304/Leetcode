package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	n := len(arr)

	root := InsertLevelOrder(arr, 0, n)
	PrintLevelOrder(root)
	fmt.Println(maxDepth(root))
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

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	return 1 + max(leftDepth, rightDepth)
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
