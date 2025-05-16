package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	node := sliceToBinaryTree([]int{4, 2, 7, 1, 2, 6, 9})
	printTree(node)
	ans := invertTree(node)
	printTree(ans)

}
func printTree(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.Val)
	printTree(root.Left)
	printTree(root.Right)
}
func sliceToBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	// Create the root node
	root := &TreeNode{Val: nums[0]}
	queue := []*TreeNode{root}

	// Process each level using a queue
	index := 1
	for index < len(nums) {
		// Get the front node from the queue
		current := queue[0]
		queue = queue[1:]

		// Add the left child if available
		if index < len(nums) {
			left := &TreeNode{Val: nums[index]}
			current.Left = left
			queue = append(queue, left)
			index++
		}

		// Add the right child if available
		if index < len(nums) {
			right := &TreeNode{Val: nums[index]}
			current.Right = right
			queue = append(queue, right)
			index++
		}
	}

	return root
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	tmpRight := invertTree(root.Left)
	tmpLeft := invertTree(root.Right)

	root.Right = tmpRight
	root.Left = tmpLeft

	return root
}
