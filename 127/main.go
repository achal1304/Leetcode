package main

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordSet := make(map[string]bool)
	for _, word := range wordList {
		wordSet[word] = true
	}

	if !wordSet[endWord] {
		return 0
	}

	queue := []string{beginWord}
	visited := make(map[string]bool)
	visited[beginWord] = true
	steps := 1

	for len(queue) > 0 {
		size := len(queue)

		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]

			if curr == endWord {
				return steps
			}

			neighbors := getNeighbors(curr, wordSet)
			for _, neighbor := range neighbors {
				if !visited[neighbor] {
					visited[neighbor] = true
					queue = append(queue, neighbor)
				}
			}
		}

		steps++
	}

	return 0
}

func getNeighbors(word string, wordSet map[string]bool) []string {
	var neighbors []string
	letters := []byte("abcdefghijklmnopqrstuvwxyz")

	for i := 0; i < len(word); i++ {
		for _, c := range letters {
			if word[i] == c {
				continue
			}
			newWord := word[:i] + string(c) + word[i+1:]
			if wordSet[newWord] {
				neighbors = append(neighbors, newWord)
			}
		}
	}

	return neighbors
}
