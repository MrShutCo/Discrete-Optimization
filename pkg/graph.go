package pkg

const NODE_COUNT = 280

type Graph [][]float64
type Path []int

// next determines the next node in a path using the modulo operator.
//    NOTE: this for some reason is 2x as fast when using % CONST instead of % VAR
func (p Path) next(i int, length int) int {
	return p[(i+1)%NODE_COUNT]
}

/*
type Path struct {
	path   []int
	length int
}

func NewPath(size int) Path {
	return Path{
		path:   make([]int, size),
		length: size,
	}
}

*/
