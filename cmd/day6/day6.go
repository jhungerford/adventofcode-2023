package main

import (
	"fmt"
	"github.com/jhungerford/adventofcode-2023/internal/day6"
)

func main() {
	races, err := day6.LoadRaces("input/day6.txt")
	if err != nil {
		fmt.Printf("failed to load races: %v", err)
	}

	fmt.Println("Part 1:", day6.Part1(races))
}
