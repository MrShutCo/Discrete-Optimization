package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"tsp/pkg"
)

func readTSPProblem(file string) (int, pkg.Graph) {
	data, _ := os.ReadFile(file)

	lines := strings.Split(string(data), "\n")

	nodes := len(lines) - 2 - 6

	pointX := make([]float64, nodes)
	pointY := make([]float64, nodes)

	for i := 0; i < nodes; i++ {
		line := lines[i+6]
		xStr := strings.ReplaceAll(line[4:7], " ", "")
		yStr := strings.ReplaceAll(line[7:], " ", "")
		x, _ := strconv.ParseInt(xStr, 10, 32)
		y, _ := strconv.ParseInt(yStr, 10, 32)
		pointX[i], pointY[i] = float64(x), float64(y)

	}
	fmt.Printf("Nodes %d\n", nodes)
	return nodes, calculateDistances(pointX, pointY)
}

// generateProblem generates n random points in R^2 in the bounds x: [0,100) y: [0,100)
// and calculates the distance matrix between those points
func generateProblem(n int) pkg.Graph {
	pointX := make([]float64, n)
	pointY := make([]float64, n)

	for i := 0; i < n; i++ {
		pointX[i] = rand.Float64() * 100
		pointY[i] = rand.Float64() * 100
	}

	return calculateDistances(pointX, pointY)
}

func calculateDistances(xCoords []float64, yCoords []float64) pkg.Graph {
	n := len(xCoords)
	g := make(pkg.Graph, n)
	for i := 0; i < n; i++ {
		g[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			g[i][j] = dist(xCoords[i], xCoords[j], yCoords[i], yCoords[j])
		}
	}
	return g
}

// dist calculates the euclidean distance between two points in R^2
func dist(x1, x2, y1, y2 float64) float64 {
	return math.Sqrt((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2))
}

var exampleGraph = pkg.Graph{{0.0, 8.1, 9.2, 7.7, 9.3, 2.3, 5.1, 10.2, 6.1, 7.0},
	{8.1, 0.0, 12.0, 0.9, 12.0, 9.5, 10.1, 12.8, 2.0, 1.0},
	{9.2, 12.0, 0.0, 11.2, 0.7, 11.1, 8.1, 1.1, 10.5, 11.5},
	{7.7, 0.9, 11.2, 0.0, 11.2, 9.2, 9.5, 12.0, 1.6, 1.1},
	{9.3, 12.0, 0.7, 11.2, 0.0, 11.2, 8.5, 1.0, 10.6, 11.6},
	{2.3, 9.5, 11.1, 9.2, 11.2, 0.0, 5.6, 12.1, 7.7, 8.5},
	{5.1, 10.1, 8.1, 9.5, 8.5, 5.6, 0.0, 9.1, 8.3, 9.3},
	{10.2, 12.8, 1.1, 12.0, 1.0, 12.1, 9.1, 0.0, 11.4, 12.4},
	{6.1, 2.0, 10.5, 1.6, 10.6, 7.7, 8.3, 11.4, 0.0, 1.1},
	{7.0, 1.0, 11.5, 1.1, 11.6, 8.5, 9.3, 12.4, 1.1, 0.0}}

func badRoute(n int) pkg.Path {
	p := make(pkg.Path, n)
	for i := 0; i < n; i++ {
		p[i] = i
	}
	return p
}

func main() {
	// Used for testing purposes
	rand.Seed(0)
	//mat := generateProblem(1000)
	n, mat := readTSPProblem("./testdata/dsj1000.tsp")

	// Modify the given graph to test others
	ge := pkg.NewGeneticAlgorithm(mat, pkg.OrderCrossover{}, pkg.PercentMutator{0.02})
	//to := pkg.NewTwoOpt(mat)
	//fmt.Printf("Matrix side: %f\n", mat[1][0])

	println("Outcome of trivial route\n==================\n")
	fmt.Printf("%f\n", 1/ge.EvaluateFitness(badRoute(n)))

	_, _ = ge.Simulate(10, 100)

	fmt.Println()

	//to.Solve(badRoute(n))
}
