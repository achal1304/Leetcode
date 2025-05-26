package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// REVISIT

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	element := scanner.Text()
	preOrder := strings.Split(element, " ")
	scanner.Scan()
	element = scanner.Text()
	inOrder := strings.Split(element, " ")
	fmt.Println(preOrder, inOrder)
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	ele := preorder[0]
	root := &TreeNode{Val: ele}
	splitIndex := 0

	for i := 0; i < len(inorder); i++ {
		if inorder[i] == ele {
			splitIndex = i
		}
	}

	root.Left = buildTree(preorder[1:splitIndex+1], inorder[:splitIndex])
	root.Right = buildTree(preorder[splitIndex+1:], inorder[splitIndex+1:])
	return root
}
