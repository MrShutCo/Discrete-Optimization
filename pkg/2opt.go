package pkg

type TwoOpt struct {
	graph Graph
	size  int
}

func NewTwoOpt(graph Graph) TwoOpt {
	return TwoOpt{
		graph: graph,
		size:  len(graph),
	}
}

// initialSolutions produces the path 0->1->2->...->n as a starting point for the algorithm
func (to TwoOpt) initialSolution() Path {
	p := make(Path, to.size)
	for i := 0; i < to.size; i++ {
		p[i] = i
	}
	return p
}

// Solve uses the 2-Opt method which consists of: generating a path, trying all permutations of swapping two
// edges, and continuing until a local optimum can be found (i.e swapping any two edges does not produce a better solution)
// Note: This does not produce optimal solutions
func (to TwoOpt) Solve(initialPath Path) (float64, Path) {
	//currentPath := to.initialSolution()
	currentPath := initialPath
	minCost := to.evaluatePath(currentPath)
	canImprove := true
	for canImprove {
		canImprove = false
		for i := 0; i < to.size-1; i++ {
			for j := i + 1; j < to.size; j++ {
				// We dont need to recalculate path length, just the difference
				pathDelta := to.evaluatePathDiff(currentPath, i, j)
				if pathDelta > 0 {
					currentPath = to.crossover(currentPath, i, j)
					canImprove = true
				}
			}
		}
	}

	minCost = to.evaluatePath(currentPath)

	//fmt.Println("\nOutcome of 2-Opt\n===============================")
	//fmt.Printf("Min cost: %f\n", minCost)
	//fmt.Printf("Path found: %v\n", currentPath)
	return minCost, currentPath
}

// evaluatePath calculates the length of a given path
func (to TwoOpt) evaluatePath(p Path) float64 {
	length := 0.0
	for _, i := range p {
		length += to.graph[p[i]][p.next(i, to.size)]
	}
	return length
}

// evaluatePathDiff calculates (Length of current edges) - (length of swapped edges)
// used for efficient calculation of paths
func (to TwoOpt) evaluatePathDiff(p Path, i, j int) float64 {
	return (to.graph[p[i]][p.next(i, to.size)] + to.graph[p[j]][p.next(j, to.size)]) - (to.graph[p[i]][p[j]] + to.graph[p.next(i, to.size)][p.next(j, to.size)])
}

// crossover swaps two edges that start with v1 and v2 (as in, they come first in the path)
// and then reverse the ordering of the path inbetween these two edges to make it a valid path again
func (to TwoOpt) crossover(path Path, v1, v2 int) Path {
	reversed_path := reverse(path[v1+1 : v2+1])
	rest_of_path := append(reversed_path, path[v2+1:]...)
	return append(path[0:v1+1], rest_of_path...)
}

// reverse reverses a given path
func reverse(p Path) Path {
	newPath := make(Path, len(p))
	for i := 0; i < len(p); i++ {
		newPath[len(p)-1-i] = p[i]
	}
	return newPath
}
