package main

import (
	"fmt"
	"github.com/jhungerford/adventofcode-2023/internal/day8"
)

func main() {
	network, err := day8.LoadNetwork("input/day8.txt")
	if err != nil {
		fmt.Printf("failed to load network: %v", err)
	}

	fmt.Println("Part 1:", day8.Part1(network))
}
