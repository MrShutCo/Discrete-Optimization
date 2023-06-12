package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Greedy(t *testing.T) {
	g := Greedy{}
	mat := Graph{
		{0, 2, 2, 1},
		{2, 0, 1, 1},
		{2, 1, 0, 2},
		{1, 1, 2, 0}}
	path := g.Solve(mat)
	assert.Equal(t, Path{0, 3, 1, 2}, path)
}
