package models

type Node struct {
	Name  string
	index int
}

func (n Node) GetIndex() int {
	return n.index
}

type Graph struct {
	adjacencyList [][]float64

	nodes []Node
}

func (g Graph) IsNeighbour(v1, v2 int) bool {
	return g.adjacencyList[v1][v2] != 0
}

func (g Graph) GetNeighbours(v1 int) []Node {
	neighbours := []Node{}
	for v2 := 0; v2 < len(g.nodes); v2++ {
		if g.adjacencyList[v1][v2] != 0 {
			neighbours = append(neighbours, g.nodes[v2])
		}
	}
	return neighbours
}
