package geometry

import (
	"math/rand"
	"testing"

	"github.com/esimov/gogu"
	"github.com/stretchr/testify/assert"
)

func Test_ConvexHull_Simple(t *testing.T) {
	points := []Point{{0, 0}, {0.5, 0.51}, {0.7, 0.1}, {1, 0}, {0.6, 0.65}, {1, 1}, {0.9, 0.8}, {0, 1}}
	convexPoly := ConvexHull(points)
	assert.Equal(t, Polygon{vertices: []Point{{0, 0}, {0, 1}, {1, 1}, {1, 0}}}, convexPoly)
}

func Test_ConvexHull_Triangle(t *testing.T) {
	points := []Point{{0, 0}, {0.5, 0.5}, {0.5, 0}}
	convexPoly := ConvexHull(points)
	assert.Equal(t, Polygon{vertices: []Point{{0, 0}, {0.5, 0.5}, {0.5, 0}}}, convexPoly)
}

func Benchmark_ConvexHull_RandomSquare(b *testing.B) {
	points := make([]Point, b.N+4)
	points[0], points[1], points[2], points[3] = Point{0, 0}, Point{1, 0}, Point{1, 1}, Point{0, 1}
	for i := 0; i < b.N; i++ {
		points[i+4] = NewPoint(rand.Float64(), rand.Float64())
	}
	pointsShuffled := gogu.Shuffle(points)
	convexPoly := ConvexHull(pointsShuffled)
	assert.Equal(b, Polygon{vertices: []Point{{0, 0}, {0, 1}, {1, 1}, {1, 0}}}, convexPoly)
}

func Test_RemoveSecondLast(t *testing.T) {
	assert.Equal(t, []Point{{1, 0}, {0, 1}, {0.5, 6}}, removeSecondLast([]Point{{1, 0}, {0, 1}, {0.5, 0.5}, {0.5, 6}}))
}

func Test_PointAboveLine(t *testing.T) {
	l1, l2 := NewPoint(1, 1), NewPoint(3, 2)
	t.Run("Point below line", func(t *testing.T) {
		assert.False(t, PointAboveLine(l1, l2, NewPoint(3, 0.5)))
	})
	t.Run("Point above line", func(t *testing.T) {
		assert.True(t, PointAboveLine(l1, l2, NewPoint(1.5, 2.5)))
	})
	t.Run("Point on line", func(t *testing.T) {
		assert.False(t, PointAboveLine(l1, l2, NewPoint(2, 1.5)))
	})
}
