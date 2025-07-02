package main

// REVISIT : MEDIUM : GRAPH

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	visited := make(map[*Node]*Node)

	var clone func(node *Node) *Node

	clone = func(node *Node) *Node {
		if node == nil {
			return nil
		}

		if n, ok := visited[node]; ok {
			return n
		}

		copy := &Node{Val: node.Val}
		visited[node] = copy

		for _, neighbor := range node.Neighbors {
			copy.Neighbors = append(copy.Neighbors, clone(neighbor))
		}

		return copy
	}

	return clone(node)
}
