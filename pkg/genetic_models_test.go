package pkg

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Crossover(t *testing.T) {
	p1 := Path{1, 2, 3, 4, 5, 6, 7, 8, 9}
	p2 := Path{5, 7, 4, 9, 1, 3, 6, 2, 8}
	c := OrderCrossover{}
	newPath := c.fixedReproduce(p1, p2, 2, 6)
	assert.Equal(t, Path{7, 9, 3, 4, 5, 6, 1, 2, 8}, newPath)

	assert.Equal(t, Path{1, 2, 3, 4, 5, 6, 7, 8, 9}, p1)
	assert.Equal(t, Path{5, 7, 4, 9, 1, 3, 6, 2, 8}, p2)
}

func Test_Crossover_Start(t *testing.T) {
	c := OrderCrossover{}
	newPath := c.fixedReproduce(Path{1, 2, 3, 4, 5, 6, 7, 8, 9}, Path{5, 7, 4, 9, 1, 3, 6, 2, 8}, 0, 4)
	assert.Equal(t, Path{1, 2, 3, 4, 5, 7, 9, 6, 8}, newPath)
}

func Test_Crossover_End(t *testing.T) {
	c := OrderCrossover{}
	newPath := c.fixedReproduce(Path{1, 2, 3, 4, 5, 6, 7, 8, 9}, Path{5, 7, 4, 9, 1, 3, 6, 2, 8}, 5, 9)
	assert.Equal(t, Path{5, 4, 1, 3, 2, 6, 7, 8, 9}, newPath)
}

func Test_Unique(t *testing.T) {
	for i := 0; i < 100; i++ {
		index := generateDiffIndex(Path{0, 1}, 0)
		assert.Equal(t, 1, index)
	}

}

func Test_Remove(t *testing.T) {
	p := remove(Path{5, 7, 4, 9, 1, 3, 6, 2, 8}, Path{3, 4, 5, 6})
	assert.Equal(t, Path{7, 9, 1, 2, 8}, p)

}

func Test_IsIn(t *testing.T) {
	yes := isIn(Path{1, 2, 3}, 3)
	assert.True(t, yes)

	no := isIn(Path{1, 2, 3}, 4)
	assert.False(t, no)
}

func Test_findInInterval(t *testing.T) {

	values := []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 9.1}
	expectedIndex := []int{0, 0, 1, 1, 1, 2, 2, 2, 2, 2, 2}

	for i := 0; i < len(values); i++ {
		index := findInInterval([]float64{2, 5, 10}, values[i])
		assert.Equal(t, expectedIndex[i], index, fmt.Sprintf("Looking for value %f in [2,5,10]", values[i]))
	}

}
