package pkg

import (
	"math"
	"math/rand"
)

type PercentMutator struct {
	// Percent chance that a mutation can occur. Needs to be in [0.0,1.0]
	ChanceToMutate float64
}

// Mutate swaps two nodes in the path with a given percentage. Used to get out of local optima
func (pm PercentMutator) Mutate(path Path) Path {
	for i := range path {
		if rand.Float64() <= pm.ChanceToMutate {
			swap := generateDiffIndex(path, i)
			path[i], path[swap] = path[swap], path[i]
		}
	}
	return path
}

type OrderCrossover struct{}

// fixedReproduce creates a new path consisting of:
//   the span of [start:end] of p1
//   all remaining nodes in p2 not in that span, in order they appear
func (c OrderCrossover) fixedReproduce(p1 Path, p2 Path, start int, end int) Path {
	child := make(Path, len(p1))

	// Copy span from first gene here
	for i := start; i < end; i++ {
		child[i] = p1[i]
	}

	p2Index := 0
	newP2 := remove(p2, p1[start:end])
	for i := 0; i < start; i++ {
		child[i] = newP2[p2Index]
		p2Index++
	}

	for i := end; i < len(p1); i++ {
		child[i] = newP2[p2Index]
		p2Index++
	}

	return child
}

// Reproduce chooses two random unique points to perform an ordered crossover with
func (c OrderCrossover) Reproduce(p1 Path, p2 Path) Path {
	a := (float64)(rand.Intn(len(p1)))
	b := (float64)(generateDiffIndex(p1, int(a)))
	startGene := int(math.Min(a, b))
	endGene := int(math.Max(a, b))

	return c.fixedReproduce(p1, p2, startGene, endGene)
}

// remove creates a new path p1 with all the subPath removed
func remove(p1 Path, subPath Path) Path {
	newPath := make(Path, len(p1)-len(subPath))
	index := 0
	for i := range p1 {
		if !isIn(subPath, p1[i]) {
			newPath[index] = p1[i]
			index++
		}
	}

	return newPath
}

// isIn determines if a node is in a given path
func isIn(p Path, node int) bool {
	for i := range p {
		if p[i] == node {
			return true
		}
	}
	return false
}

// generateDiffIndex chooses a random point on the given path that isnt the excluded one
func generateDiffIndex(path Path, excluded int) int {
	x := rand.Intn(len(path))
	for x == excluded {
		x = rand.Intn(len(path))
	}
	return x
}
