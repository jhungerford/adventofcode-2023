package main

import (
	"fmt"
	"github.com/jhungerford/adventofcode-2023/internal/day16"
)

func main() {
	m, err := day16.LoadMap("input/day16.txt")
	if err != nil {
		fmt.Printf("failed to load map: %v", err)
		return
	}

	fmt.Println("Part 1: ", day16.Part1(m))
	fmt.Println("Part 2: ", day16.Part2(m))
}
