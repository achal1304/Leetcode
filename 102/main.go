package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	res := [][]int{}

	for len(queue) > 0 {
		qLen := len(queue)
		tempArr := []int{}

		for i := 0; i < qLen; i++ {
			node := queue[0]
			queue = queue[1:]
			tempArr = append(tempArr, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		res = append(res, tempArr)

	}

	return res

}
