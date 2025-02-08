package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// arr := []int{10, 5, -3, 3, 2, 0, 11, 3, -2, 0, 1}
	arr := []int{1, 0, 1, 1, 1, 0, 0, 1, 1, 0, 1, 0, 0, 0, 1}
	n := len(arr)

	root := InsertLevelOrder(arr, 0, n)
	PrintLevelOrder(root)
	fmt.Println("answer is ")
	fmt.Println(longestZigZag(root))
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

func longestZigZag(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxCount := 0

	zigZag(root.Left, 1, &maxCount, false)
	zigZag(root.Right, 1, &maxCount, true)

	return maxCount

}

func zigZag(root *TreeNode, currCount int, maxCount *int, isRight bool) {
	if root == nil {
		return
	}

	fmt.Println("currentcount ", currCount, "maxCount ", maxCount, "root", root, isRight)
	if currCount > *maxCount {
		*maxCount = currCount
	}

	if isRight {
		zigZag(root.Left, currCount+1, maxCount, false)
		zigZag(root.Right, 1, maxCount, true)
	} else {
		zigZag(root.Left, 1, maxCount, false)
		zigZag(root.Right, currCount+1, maxCount, true)
	}

	fmt.Println("returning maxCount ", maxCount)
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
