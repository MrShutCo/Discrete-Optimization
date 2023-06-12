package geometry

import (
	"fmt"
	"sort"
)

type Point struct {
	X float64
	Y float64
}

func NewPoint(x, y float64) Point {
	return Point{X: x, Y: y}
}

func (p Point) Equal(other Point) bool {
	return p.X == other.X && p.Y == other.Y
}

func (p Point) String() string {
	return fmt.Sprintf("(%f, %f)", p.X, p.Y)
}

func (p Point) LessThan(other Point) bool {
	if p.X == other.X {
		return p.Y < other.Y
	}
	return p.X < other.X
}

func SortPoints(points []Point) []Point {
	sort.SliceStable(points, func(i, j int) bool {
		return points[i].LessThan(points[j])
	})
	return points
}
