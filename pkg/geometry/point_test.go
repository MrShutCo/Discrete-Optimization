package geometry

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SortPoints(t *testing.T) {
	unsorted := []Point{NewPoint(3, 4), NewPoint(1, 2), NewPoint(1.5, 1.5), NewPoint(0, 0), NewPoint(0, 5), NewPoint(5, 0), NewPoint(5, 5), NewPoint(-5, 6)}
	uc := []Point{NewPoint(3, 4), NewPoint(1, 2), NewPoint(1.5, 1.5), NewPoint(0, 0), NewPoint(0, 5), NewPoint(5, 0), NewPoint(5, 5), NewPoint(-5, 6)}
	sorted := SortPoints(unsorted)
	for x := range sorted {
		fmt.Println(sorted[x].String())
	}
	assert.Equal(t, []Point{uc[7], uc[3], uc[4], uc[1], uc[2], uc[0], uc[5], uc[6]}, sorted)
}
