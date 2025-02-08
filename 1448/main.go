package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	arr := []int{3, 1, 4, 3, 0, 1, 5}
	n := len(arr)

	root := InsertLevelOrder(arr, 0, n)
	PrintLevelOrder(root)
	fmt.Println("answer is ")
	fmt.Println(goodNodes(root))
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

func goodNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return checkCount(root, root.Val)
}

func checkCount(root *TreeNode, maxVal int) int {
	if root == nil {
		return 0
	}

	currCount := 0
	if root.Val >= maxVal {
		currCount++
		maxVal = root.Val
	}

	currCount += checkCount(root.Left, maxVal)
	currCount += checkCount(root.Right, maxVal)

	return currCount
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
