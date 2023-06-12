package geometry

import "fmt"

type Polygon struct {
	vertices []Point
}

func ConvexHull(points []Point) Polygon {
	SortPoints(points)
	fmt.Println(points)
	n := len(points) - 1

	upper_haul := []Point{points[0], points[1]}
	for i := 2; i <= n; i++ {
		upper_haul = append(upper_haul, points[i])
		last_i := len(upper_haul) - 1
		for len(upper_haul) > 2 && PointAboveLine(upper_haul[last_i-2], upper_haul[last_i-1], upper_haul[last_i]) {
			upper_haul = removeSecondLast(upper_haul)
			last_i = len(upper_haul) - 1
		}
	}

	fmt.Println("Computing Lower haul")

	lower_haul := []Point{points[n], points[n-1]}
	for i := n - 1; i >= 0; i-- {
		lower_haul = append(lower_haul, points[i])
		last_i := len(lower_haul) - 1
		for len(lower_haul) > 2 && PointBelowLine(lower_haul[last_i-2], lower_haul[last_i-1], lower_haul[last_i]) {
			lower_haul = removeSecondLast(lower_haul)
			last_i = len(lower_haul) - 1
		}
	}
	lower_haul = lower_haul[1 : len(lower_haul)-2]

	return Polygon{append(upper_haul, lower_haul...)}
}

func removeSecondLast(p []Point) []Point {
	return append(p[:len(p)-2], p[len(p)-1:]...)
}

func PointAboveLine(l1, l2, p Point) bool {
	m, b := CalculateLine(l1, l2)
	y := m*p.X + b
	return y-p.Y < 0
}

func PointBelowLine(l1, l2, p Point) bool {
	m, b := CalculateLine(l1, l2)
	y := m*p.X + b
	return y-p.Y > 0
}

func CalculateSlope(l1, l2 Point) float64 {
	return (l2.Y - l1.Y) / (l2.X - l1.X)
}

func CalculateLine(l1, l2 Point) (m float64, b float64) {
	m = CalculateSlope(l1, l2)
	b = l2.Y - m*l2.X
	return
}
