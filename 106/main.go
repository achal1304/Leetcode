package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 || len(inorder) == 0 {
		return nil
	}

	ele := postorder[len(postorder)-1]
	root := &TreeNode{Val: ele}
	splitIndex := 0

	for i := 0; i < len(inorder); i++ {
		if inorder[i] == ele {
			splitIndex = i
		}
	}

	root.Left = buildTree(inorder[:splitIndex], postorder[:splitIndex])
	root.Right = buildTree(inorder[splitIndex+1:], postorder[splitIndex:len(postorder)-1])
	return root
}
