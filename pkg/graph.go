package pkg

type Graph [][]float64
type Path []int

func (p Path) next(i int) int {
	return p[(i+1)%len(p)]
}
