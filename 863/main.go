package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findParents(root *TreeNode, parents map[*TreeNode]*TreeNode) {
	if root == nil {
		return
	}

	if root.Left != nil {
		parents[root.Left] = root
		findParents(root.Left, parents)
	}
	if root.Right != nil {
		parents[root.Right] = root
		findParents(root.Right, parents)
	}
}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	if root == nil {
		return []int{}
	}
	parents := make(map[*TreeNode]*TreeNode)
	findParents(root, parents)

	queue := []*TreeNode{target}
	distance := 0
	visited := make(map[*TreeNode]bool)
	visited[target] = true

	for len(queue) > 0 && distance < k {
		newQueue := []*TreeNode{}

		for _, node := range queue {
			if node.Left != nil && !visited[node.Left] {
				newQueue = append(newQueue, node.Left)
				visited[node.Left] = true
			}
			if node.Right != nil && !visited[node.Right] {
				newQueue = append(newQueue, node.Right)
				visited[node.Right] = true
			}
			if parents[node] != nil && !visited[parents[node]] {
				newQueue = append(newQueue, parents[node])
				visited[parents[node]] = true
			}
		}

		queue = newQueue
		distance++
	}
	result := []int{}

	for _, ele := range queue {
		result = append(result, ele.Val)
	}

	return result
}
