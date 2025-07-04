package main

func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := make(map[int][]int)
	inDegree := make([]int, numCourses)

	for _, pre := range prerequisites {
		a, b := pre[0], pre[1]
		graph[b] = append(graph[b], a)
		inDegree[a]++
	}

	queue := []int{}

	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	processed := 0
	ans := []int{}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		ans = append(ans, curr)
		processed++

		for _, n := range graph[curr] {
			inDegree[n]--
			if inDegree[n] == 0 {
				queue = append(queue, n)
			}
		}
	}

	if processed == numCourses {
		return ans
	} else {
		return []int{}
	}

}
