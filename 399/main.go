package main

func calcEquation(
	equations [][]string,
	values []float64,
	queries [][]string,
) []float64 {
	// Step 1: Build the graph
	graph := make(map[string]map[string]float64)

	for i, eq := range equations {
		a, b := eq[0], eq[1]
		val := values[i]

		if graph[a] == nil {
			graph[a] = make(map[string]float64)
		}
		if graph[b] == nil {
			graph[b] = make(map[string]float64)
		}

		graph[a][b] = val
		graph[b][a] = 1 / val
	}

	// Step 2: Process each query
	var results []float64
	for _, query := range queries {
		start, end := query[0], query[1]
		if graph[start] == nil || graph[end] == nil {
			results = append(results, -1.0)
		} else if start == end {
			results = append(results, 1.0)
		} else {
			visited := make(map[string]bool)
			result := dfs(graph, start, end, 1.0, visited)
			results = append(results, result)
		}
	}

	return results
}

func dfs(
	graph map[string]map[string]float64,
	curr, target string,
	product float64,
	visited map[string]bool,
) float64 {
	if curr == target {
		return product
	}

	visited[curr] = true

	for neighbor, val := range graph[curr] {
		if !visited[neighbor] {
			res := dfs(graph, neighbor, target, product*val, visited)
			if res != -1.0 {
				return res
			}
		}
	}

	return -1.0
}
