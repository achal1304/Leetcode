package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const nilVal = -99

func main() {
	// arr := []int{3, 5, 1, 6, 2, 9, 8, nilVal, nilVal, 7, 4}
	// arr2 := []int{3, 5, 1, 6, 7, 4, 2, nilVal, nilVal, nilVal, nilVal, nilVal, nilVal, 9, 8}
	arr := []int{1, 2, 3}
	arr2 := []int{1, 3, 2}

	n := len(arr)
	n2 := len(arr2)
	root := InsertLevelOrder(arr, 0, n)
	root2 := InsertLevelOrder(arr2, 0, n2)
	PrintLevelOrder(root)
	fmt.Println()
	PrintLevelOrder(root2)
	fmt.Println()
	fmt.Println(leafSimilar(root, root2))
}

func InsertLevelOrder(arr []int, i, n int) *TreeNode {
	if i >= n {
		return nil
	}

	root := &TreeNode{Val: arr[i]}
	root.Left = InsertLevelOrder(arr, 2*i+1, n)
	root.Right = InsertLevelOrder(arr, 2*i+2, n)
	return root
}

func PrintLevelOrder(root *TreeNode) {
	if root == nil {
		return
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		fmt.Print(node.Val, " ")

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leafArr1 := []int{}
	getLeafArray(root1, &leafArr1)
	leafArr2 := []int{}
	getLeafArray(root2, &leafArr2)
	fmt.Println(leafArr1)
	fmt.Println(leafArr2)
	if len(leafArr1) != len(leafArr2) {
		return false
	}

	for i := 0; i < len(leafArr1); i++ {
		if leafArr1[i] != leafArr2[i] {
			return false
		}
	}
	return true
}

func getLeafArray(root *TreeNode, leafArr *[]int) {
	fmt.Println("inside getleaf ", root.Val, root.Left, root.Right)
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		*leafArr = append(*leafArr, root.Val)
		fmt.Println("arrauy updated ", leafArr)
		return
	}

	if root.Left != nil {
		getLeafArray(root.Left, leafArr)
	}

	if root.Right != nil {
		getLeafArray(root.Right, leafArr)
	}

	return
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
