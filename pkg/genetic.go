package pkg

import (
	"fmt"
	"math/rand"
)

type Population []Path
type PopulationFitness []float64

type Reproducer interface {
	Reproduce(Path, Path) Path
}

type Mutator interface {
	Mutate(Path) Path
}

type GeneticAlgorithm struct {
	graph     Graph
	nodeCount int

	population       Population
	currentFitnesses PopulationFitness

	absoluteBest     float64
	absoluteBestPath Path

	reproducer     Reproducer
	mutator        Mutator
	localOptimizer TwoOpt

	cumSumFitness []float64
}

func NewGeneticAlgorithm(graph Graph, r Reproducer, m Mutator) GeneticAlgorithm {
	return GeneticAlgorithm{
		graph:          graph,
		nodeCount:      len(graph),
		reproducer:     r,
		mutator:        m,
		localOptimizer: NewTwoOpt(graph),
	}
}

// createPopulation creates a randomly generated starting population of valid paths
func (ge *GeneticAlgorithm) createPopulation(size int) {
	population := make(Population, size)
	ge.currentFitnesses = make(PopulationFitness, size)
	for i := 0; i < size; i++ {
		population[i] = rand.Perm(ge.nodeCount)
	}
	ge.population = population
}

// evaluatePopulation assigns a fitness to every path in the population, updating any optimal paths found
func (ge *GeneticAlgorithm) evaluatePopulation() PopulationFitness {
	fitness := make(PopulationFitness, len(ge.population))
	for i, path := range ge.population {
		fitness[i] = ge.EvaluateFitness(path)
		if fitness[i] > ge.absoluteBest {
			ge.absoluteBest = fitness[i]
			ge.absoluteBestPath = ge.population[i]
		}
	}
	return fitness
}

// evaluate determines the inverse length of a given path
//
// The inverse is taken to account for the fact this is a minimization problem,
// and standard genetic algorithms are for maximization
func (ge GeneticAlgorithm) EvaluateFitness(p Path) float64 {
	length := 0.0
	for _, i := range p {
		length += ge.graph[p[i]][p.next(i)]
		length += ge.graph[p[i]][p.next(i, ge.nodeCount)]
	}
	return 1 / length
}

func (ge *GeneticAlgorithm) ApplyLocalOptimization() {
	for i := 0; i < len(ge.population); i++ {
		ge.currentFitnesses[i], ge.population[i] = ge.localOptimizer.Solve(ge.population[i])
		fmt.Println(i)
	}
}

// chooseParent uses Roulette selection to choose a parent, based proportionally on how high
// the parents fitness is
func (ga *GeneticAlgorithm) chooseParent(pop Population, fitness PopulationFitness) Path {
	pick := rand.Float64() * ga.cumSumFitness[len(ga.population)-1]
	return pop[findInInterval(ga.cumSumFitness, pick)]
}

func findInInterval(values []float64, target float64) int {
	if target < values[0] {
		return 0
	}
	left := 0
	right := len(values)
	for left <= right {
		middle := left + (right-left)/2
		//println(middle)
		// We have found our target
		if target >= values[middle] && target < values[middle+1] {
			return middle + 1
		}
		// Need to look in left half of search space
		if target < values[middle] {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left - 1
}

// reproduce creates a new generation of paths based on their fitness with roulette selection
// for choosing parents, and the given reproduce and mutate algorithms
func (ge *GeneticAlgorithm) reproduce() {
	newPopulation := make(Population, len(ge.population))

	// Calculate running sum to use in roulette selection
	ge.cumSumFitness = make([]float64, len(ge.population))
	ge.cumSumFitness[0] = ge.cumSumFitness[0]
	for i := 1; i < len(ge.population); i++ {
		ge.cumSumFitness[i] = ge.currentFitnesses[i-1] + ge.currentFitnesses[i]
	}

	for i := 0; i < len(ge.population); i++ {
		parent1 := ge.chooseParent(ge.population, ge.currentFitnesses)
		parent2 := ge.chooseParent(ge.population, ge.currentFitnesses)

		child := ge.reproducer.Reproduce(parent1, parent2)
		child = ge.mutator.Mutate(child)

		_, child = ge.localOptimizer.Solve(child)
		newPopulation[i] = child
	}
	ge.population = newPopulation
}

// Simulate runs the generic algorithm with the specified popSize and for the given generations
// Prints out the minimum cost path found.
// Note: algorithm has no guarantee on optimality
func (ge *GeneticAlgorithm) Simulate(popSize, generationCount int) (float64, Path) {
	ge.createPopulation(popSize)

	fmt.Println("Starting Local Optimization")
	ge.ApplyLocalOptimization()
	fmt.Println("Finished Applying Local Optimization")

	for i := 0; i < generationCount; i++ {
		ge.currentFitnesses = ge.evaluatePopulation()
		ge.reproduce()
		fmt.Printf("Finished Gen %d\n", i)
	}

	fmt.Println("\nOutcome of Genetic algorithm\n===============================")
	fmt.Printf("Min cost: %f\n", 1/ge.absoluteBest)
	fmt.Printf("Best Path: %v", ge.absoluteBestPath)
	return 1 / ge.absoluteBest, ge.absoluteBestPath
}
