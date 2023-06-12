package pkg

type Greedy struct{}

func (g Greedy) Solve(graph Graph) Path {
	path := make(Path, len(graph))

	path[0] = 0
	visited := map[int]bool{0: true}
	lastIndex := 0
	i := 1

	for len(visited) != len(graph) {
		// Going from lastIndex -> nextEdge
		nextEdge := 1000000.0
		nextEdgeIndex := -1
		for j := 0; j < len(graph); j++ {
			if graph[lastIndex][j] < nextEdge && !visited[j] && lastIndex != j {
				nextEdge = graph[lastIndex][j]
				nextEdgeIndex = j
			}
		}
		path[i] = nextEdgeIndex
		i++
		visited[nextEdgeIndex] = true
		lastIndex = nextEdgeIndex
	}

	return path
}
