package main

type TrieNode struct {
	children map[byte]*TrieNode
	word     string
}

func findWords(board [][]byte, words []string) []string {
	// Step 1: Build the Trie
	root := &TrieNode{children: make(map[byte]*TrieNode)}
	for _, word := range words {
		node := root
		for i := 0; i < len(word); i++ {
			c := word[i]
			if node.children[c] == nil {
				node.children[c] = &TrieNode{children: make(map[byte]*TrieNode)}
			}
			node = node.children[c]
		}
		node.word = word // store word at the end node
	}

	var res []string
	m, n := len(board), len(board[0])

	// Step 2: DFS for each cell
	var dfs func(i, j int, node *TrieNode)
	dfs = func(i, j int, node *TrieNode) {
		c := board[i][j]
		child := node.children[c]
		if child == nil {
			return
		}

		if child.word != "" {
			res = append(res, child.word)
			child.word = "" // avoid duplicate
		}

		board[i][j] = '#' // mark visited

		directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
		for _, d := range directions {
			ni, nj := i+d[0], j+d[1]
			if ni >= 0 && nj >= 0 && ni < m && nj < n && board[ni][nj] != '#' {
				dfs(ni, nj, child)
			}
		}

		board[i][j] = c // backtrack
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if root.children[board[i][j]] != nil {
				dfs(i, j, root)
			}
		}
	}

	return res
}
