package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	res := [][]int{}
	isLeft := true

	for len(queue) > 0 {
		qLen := len(queue)
		tempArr := make([]int, qLen)

		for i := 0; i < qLen; i++ {
			node := queue[0]
			queue = queue[1:]

			if isLeft {
				tempArr[i] = node.Val
			} else {
				tempArr[qLen-i-1] = node.Val
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		isLeft = !isLeft
		res = append(res, tempArr)
	}

	return res
}
