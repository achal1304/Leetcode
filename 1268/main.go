package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(suggestedProducts([]string{"mobile", "mouse", "moneypot", "monitor", "mousepad"}, "mouse"))
}

type TrieNode struct {
	children [26]*TrieNode
	words    []string
}

type Trie struct {
	root *TrieNode
}

func (t *Trie) Insert(word string) {
	node := t.root
	for _, ch := range word {
		index := ch - 'a'
		if node.children[index] == nil {
			node.children[index] = &TrieNode{}
		}
		node = node.children[index]

		if len(node.words) < 3 {
			node.words = append(node.words, word)
		}
	}
}

func suggestedProducts(products []string, searchWord string) [][]string {
	trie := &Trie{root: &TrieNode{}}
	sort.Strings(products)
	for _, ele := range products {
		trie.Insert(ele)
	}

	result := [][]string{}
	prefix := []rune{}
	for _, ele := range searchWord {
		prefix = append(prefix, ele)
		result = append(result, trie.searchPrefix(string(prefix)))
	}
	return result
}

func (t *Trie) searchPrefix(prefix string) []string {
	node := t.root
	for _, ch := range prefix {
		index := ch - 'a'
		if node.children[index] == nil {
			return nil
		}
		node = node.children[index]
	}
	return node.words
}
