package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	targetSum -= root.Val

	if targetSum == 0 && root.Left == nil && root.Right == nil {
		return true
	}
	if root.Left != nil {
		if hasPathSum(root.Left, targetSum) {
			return true
		}
	}

	if root.Right != nil {
		if hasPathSum(root.Right, targetSum) {
			return true
		}
	}

	return false
}
