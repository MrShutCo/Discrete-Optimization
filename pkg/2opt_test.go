package pkg

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

var exampleGraph = Graph{{0.0, 8.1, 9.2, 7.7, 9.3, 2.3, 5.1, 10.2, 6.1, 7.0},
	{8.1, 0.0, 12.0, 0.9, 12.0, 9.5, 10.1, 12.8, 2.0, 1.0},
	{9.2, 12.0, 0.0, 11.2, 0.7, 11.1, 8.1, 1.1, 10.5, 11.5},
	{7.7, 0.9, 11.2, 0.0, 11.2, 9.2, 9.5, 12.0, 1.6, 1.1},
	{9.3, 12.0, 0.7, 11.2, 0.0, 11.2, 8.5, 1.0, 10.6, 11.6},
	{2.3, 9.5, 11.1, 9.2, 11.2, 0.0, 5.6, 12.1, 7.7, 8.5},
	{5.1, 10.1, 8.1, 9.5, 8.5, 5.6, 0.0, 9.1, 8.3, 9.3},
	{10.2, 12.8, 1.1, 12.0, 1.0, 12.1, 9.1, 0.0, 11.4, 12.4},
	{6.1, 2.0, 10.5, 1.6, 10.6, 7.7, 8.3, 11.4, 0.0, 1.1},
	{7.0, 1.0, 11.5, 1.1, 11.6, 8.5, 9.3, 12.4, 1.1, 0.0}}

func Test_reverse(t *testing.T) {
	p := Path{1, 2, 3, 4, 5}
	r := reverse(p)
	assert.Equal(t, Path{5, 4, 3, 2, 1}, r)

	assert.Equal(t, Path{1, 2, 3, 4, 5}, p)
}

func Test_TwoOpt_Crossover(t *testing.T) {
	to := TwoOpt{}

	p := Path{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	cross := to.crossover(p, 0, 2)
	assert.Equal(t, Path{0, 2, 1, 3, 4, 5, 6, 7, 8, 9}, cross)
	assert.Equal(t, Path{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, p)
}

func Test_TwoOpt_Eval(t *testing.T) {
	to := TwoOpt{
		graph: exampleGraph,
		size:  10,
	}

	length := to.evaluatePath(Path{0, 9, 1, 3, 8, 4, 7, 2, 6, 5})
	assert.Equal(t, 39.2, length)
}

func Test_evalutePathDiff(t *testing.T) {
	to := TwoOpt{
		graph: exampleGraph,
		size:  10,
	}

	p1 := Path{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	p2 := Path{0, 8, 7, 6, 5, 4, 3, 2, 1, 9}

	diff1 := math.Abs(to.evaluatePathDiff(p1, 0, 8))
	diff2 := math.Abs(to.evaluatePathDiff(p2, 0, 8))
	realDiff := math.Abs(to.evaluatePath(p1) - to.evaluatePath(p2))

	assert.InDelta(t, diff1, diff2, 0.01)
	assert.InDelta(t, diff1, realDiff, 0.01)
	assert.InDelta(t, diff2, realDiff, 0.01)
}
