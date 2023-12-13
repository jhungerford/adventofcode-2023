package day8

import (
	"github.com/jhungerford/adventofcode-2023/internal/util"
	"sync"
)

// LoadNetwork loads a network from the given file.  The first line of the file is a list of L and R instructions,
// and subsequent lines are nodes in the network.
func LoadNetwork(filename string) (Network, error) {
	parsers := map[string]util.SectionLineParser[Network]{
		"instructions": func(line string, network *Network) (string, error) {
			network.directions = []byte(line)

			return "nodes", nil
		},
		"nodes": func(line string, network *Network) (string, error) {
			// A node line looks like 'AAA = (BBB, CCC)'
			network.nodes[line[:3]] = node{
				left:  line[7:10],
				right: line[12:15],
			}

			return "nodes", nil
		},
	}

	return util.ParseInputLinesSections(filename, "instructions", Network{nodes: map[string]node{}}, true, parsers)
}

// Part1 calculates the number of steps it takes to go from AAA to ZZZ in the network.  Directions are repeated as
// many times as needed to traverse the network.
func Part1(network Network) int {
	return network.steps("AAA", func(n string) bool {
		return n == "ZZZ"
	})
}

// Part2 starts at every node ending in 'A' simultaneously, and returns the number of steps until all nodes end with Z.
func Part2(network Network) int {
	var starts []string

	for n, _ := range network.nodes {
		if n[2] == 'A' {
			starts = append(starts, n)
		}
	}

	endsWithZ := func(n string) bool {
		return n[2] == 'Z'
	}

	// Find the path length for each node that ends with 'A' in parallel.
	lengths := make(chan int, len(starts))
	var wg sync.WaitGroup

	for _, start := range starts {
		wg.Add(1)

		go func(n string) {
			defer wg.Done()

			lengths <- network.steps(n, endsWithZ)
		}(start)
	}

	wg.Wait()
	close(lengths)

	// Total steps is the LCM of each path's steps.
	var pathSteps []int

	for length := range lengths {
		pathSteps = append(pathSteps, length)
	}

	return util.LCM(pathSteps...)
}

// steps returns the number of steps it takes to go from the given start position to a node satisfying atEnd.
func (network Network) steps(start string, atEnd func(string) bool) int {
	steps := 0

	for n := start; !atEnd(n); {
		n = network.nodes[n].next(network.directions[steps%len(network.directions)])

		steps++
	}

	return steps
}

// Network consists of a list of directions and a graph of nodes.
type Network struct {
	directions []byte
	nodes      map[string]node
}

// node contains left and right instructions.
type node struct {
	left  string
	right string
}

// next returns the next left or right direction of this node.
func (n node) next(direction byte) string {
	if direction == 'L' {
		return n.left
	}

	return n.right
}
