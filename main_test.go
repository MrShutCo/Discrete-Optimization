package main

import (
	"math/rand"
	"testing"
	"tsp/pkg"

	"github.com/stretchr/testify/assert"
)

func Test_GeneticAlgorithm(t *testing.T) {
	rand.Seed(0)
	mat := generateProblem(100)

	ge := pkg.NewGeneticAlgorithm(mat, pkg.OrderCrossover{}, pkg.PercentMutator{0.01})
	optValue, optPath := ge.Simulate(60, 100)

	assert.Equal(t, 0, optValue)
	assert.Equal(t, pkg.Path{}, optPath)
}

/*
func Test_TwoOpt(t *testing.T) {
	rand.Seed(0)
	mat := generateProblem(100)

	ge := pkg.NewTwoOpt(mat)
	ge.Solve()
}
*/
