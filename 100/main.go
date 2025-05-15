package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func isSameTree(p *TreeNode, q *TreeNode) bool {
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

	isLeftSame := isSameTree(p.Left, q.Left)
	isRIghtSame := isSameTree(p.Right, q.Right)

	return isLeftSame && isRIghtSame
}
