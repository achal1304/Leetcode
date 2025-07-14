package main

type TrieNode struct {
	children map[byte]*TrieNode
	isWord   bool
}

type WordDictionary struct {
	root *TrieNode
}

func Constructor() WordDictionary {
	return WordDictionary{root: &TrieNode{children: make(map[byte]*TrieNode)}}
}

func (wd *WordDictionary) AddWord(word string) {
	node := wd.root
	for i := 0; i < len(word); i++ {
		c := word[i]
		if node.children[c] == nil {
			node.children[c] = &TrieNode{children: make(map[byte]*TrieNode)}
		}
		node = node.children[c]
	}
	node.isWord = true
}

func (wd *WordDictionary) Search(word string) bool {
	return dfs(word, 0, wd.root)
}

func dfs(word string, index int, node *TrieNode) bool {
	if node == nil {
		return false
	}
	if index == len(word) {
		return node.isWord
	}

	c := word[index]
	if c == '.' {
		for _, child := range node.children {
			if dfs(word, index+1, child) {
				return true
			}
		}
		return false
	} else {
		return dfs(word, index+1, node.children[c])
	}
}
