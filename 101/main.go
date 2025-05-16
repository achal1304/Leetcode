package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	node := sliceToBinaryTree([]int{1, 2, 2, 3, 4, 4, 3})
	printTree(node)
	ans := isSymmetric(node)
	fmt.Println(ans)

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

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return symCheck(root.Left, root.Right)
}

func symCheck(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil && q != nil {
		return false
	} else if p != nil && q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}
	return symCheck(p.Left, q.Right) && symCheck(p.Right, q.Left)
}
