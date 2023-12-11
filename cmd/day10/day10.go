package main

import (
	"fmt"
	"github.com/jhungerford/adventofcode-2023/internal/day10"
)

func main() {
	m, err := day10.LoadMap("input/day10.txt")
	if err != nil {
		fmt.Printf("failed to load map: %v", err)
	}

	fmt.Println("Part 1:", day10.Part1(m))
}
