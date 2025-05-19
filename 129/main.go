package main

import "fmt"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return generateNum(root, 0, 0)
}

func generateNum(root *TreeNode, num int, sum int) int {
	if root == nil {
		return 0
	}

	num = num*10 + root.Val
	if root.Left == nil && root.Right == nil {
		sum += num
		fmt.Println("adding sum", sum, num)
		return sum
	}

	if root.Left != nil {
		sum = generateNum(root.Left, num, sum)
	}

	if root.Right != nil {
		sum = generateNum(root.Right, num, sum)
	}

	return sum
}
