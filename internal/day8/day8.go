package day8

import "github.com/jhungerford/adventofcode-2023/internal/util"

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

	return util.ParseInputLinesSections(filename, "instructions", Network{nodes: map[string]node{}}, parsers)
}

// Part1 calculates the number of steps it takes to go from AAA to ZZZ in the network.  Directions are repeated as
// many times as needed to traverse the network.
func Part1(network Network) int {
	steps := 0

	for n := "AAA"; n != "ZZZ"; {
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
