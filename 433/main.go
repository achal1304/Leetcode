package main

func minMutation(startGene string, endGene string, bank []string) int {
	bankSet := make(map[string]bool)
	for _, gene := range bank {
		bankSet[gene] = true
	}

	if !bankSet[endGene] {
		return -1
	}

	genes := []byte{'A', 'C', 'G', 'T'}
	visited := make(map[string]bool)
	queue := []string{startGene}
	steps := 0

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]

			if curr == endGene {
				return steps
			}

			for j := 0; j < len(curr); j++ {
				for _, g := range genes {
					if curr[j] == g {
						continue
					}
					mutated := curr[:j] + string(g) + curr[j+1:]
					if bankSet[mutated] && !visited[mutated] {
						visited[mutated] = true
						queue = append(queue, mutated)
					}
				}
			}
		}
		steps++
	}

	return -1
}
