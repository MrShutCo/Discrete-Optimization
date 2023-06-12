package main

import (
	"math/rand"
	"testing"
	"tsp/pkg"
)

/*func Test_GeneticAlgorithm(t *testing.T) {
	rand.Seed(0)
	//mat := generateProblem(100)
	_, mat := readTSPProblem("./testdata/pr1002.tsp")

	ge := pkg.NewGeneticAlgorithm(mat, pkg.OrderCrossover{}, pkg.PercentMutator{0.04})
	ge.Simulate(10, 100)

	fmt.Println("\n\nRandom permutation solution:\n===============================")
	fmt.Println(1 / ge.EvaluateFitness(rand.Perm(1001)))
}
*/

func Test_TwoOpt(t *testing.T) {
	rand.Seed(0)
	_, mat := readTSPProblem("./testdata/a280.tsp")

	ge := pkg.NewTwoOpt(mat)
	ge.Output = true

	//ge.Solve(pkg.Greedy{}.Solve(mat))
	ge.Solve(badRoute(280))
}
