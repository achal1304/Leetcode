package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}

	var result []float64
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSum := 0.0
		levelCount := len(queue)

		for i := 0; i < levelCount; i++ {
			node := queue[0]
			queue = queue[1:]

			levelSum += float64(node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		avg := levelSum / float64(levelCount)
		result = append(result, avg)
	}

	return result
}
