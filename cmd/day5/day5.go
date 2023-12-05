package main

import (
	"fmt"
	"github.com/jhungerford/adventofcode-2023/internal/day5"
)

func main() {
	almanac, err := day5.LoadAlmanac("input/day5.txt")
	if err != nil {
		fmt.Printf("failed to load almanac: %v", err)
	}

	fmt.Println("Part 1:", day5.Part1(almanac))
	fmt.Println("Part 2:", day5.Part2(almanac))
}
