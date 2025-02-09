package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// arr := []int{10, 5, -3, 3, 2, 0, 11, 3, -2, 0, 1}
	arr := []int{1, 2, 3, 0, 5, 0}
	n := len(arr)

	root := InsertLevelOrder(arr, 0, n)
	PrintLevelOrder(root)
	fmt.Println("answer is ")
	fmt.Println(rightSideView(root))
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

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	arr := []int{}
	rightViewTraversal(root, &arr, 0)
	return arr
}

func rightViewTraversal(root *TreeNode, arr *[]int, index int) {
	if root == nil {
		return
	}

	if len(*arr)-1 < index {
		*arr = append(*arr, root.Val)
	}

	rightViewTraversal(root.Right, arr, index+1)
	rightViewTraversal(root.Left, arr, index+1)
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
