package main

import (
	"fmt"
	"github.com/jhungerford/adventofcode-2023/internal/day3"
)

func main() {
	schematic, err := day3.LoadSchematic("input/day3.txt")
	if err != nil {
		fmt.Printf("failed to load schematic: %v", err)
		return
	}

	fmt.Println("Part 1:", day3.Part1(schematic))
}
