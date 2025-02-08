package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// arr := []int{10, 5, -3, 3, 2, 0, 11, 3, -2, 0, 1}
	arr := []int{5, 4, 8, 11, 0, 13, 4, 7, 2, 0, 0, 5, 1}
	n := len(arr)

	root := InsertLevelOrder(arr, 0, n)
	PrintLevelOrder(root)
	fmt.Println("answer is ")
	fmt.Println(pathSum(root, 22))
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

func pathSum(root *TreeNode, targetSum int) int {
	prefixSumCount := make(map[int]int)
	prefixSumCount[0] = 1

	return countPaths(root, 0, targetSum, prefixSumCount)
}

func countPaths(node *TreeNode, currentSum, targetSum int, prefixSumCount map[int]int) int {
	if node == nil {
		return 0
	}

	currentSum += node.Val

	count := prefixSumCount[currentSum-targetSum]

	prefixSumCount[currentSum]++

	count += countPaths(node.Left, currentSum, targetSum, prefixSumCount)
	count += countPaths(node.Right, currentSum, targetSum, prefixSumCount)

	prefixSumCount[currentSum]--

	return count
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
