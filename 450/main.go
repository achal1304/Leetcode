package main

import (
	"fmt"
)

// REVISIT
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	arr := []int{5, 3, 6, 2, 4, 1, 7}
	n := len(arr)

	root := InsertLevelOrder(arr, 0, n)
	PrintLevelOrder(root)
	fmt.Println("answer is ")
	ans := deleteNode(root, 3)
	PrintLevelOrder(ans)
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

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else {
		if root.Left == nil && root.Right == nil {
			return nil
		}

		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}

		curr := root.Right
		for curr.Left != nil {
			curr = curr.Left
		}

		root.Val = curr.Val
		root.Right = deleteNode(root.Right, curr.Val)
	}
	return root
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
