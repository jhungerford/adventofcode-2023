package main

import (
	"fmt"
	"github.com/jhungerford/adventofcode-2023/internal/day1"
	"github.com/jhungerford/adventofcode-2023/internal/util"
)

func main() {
	lines, err := util.ReadInputLines("input/day1.txt")
	if err != nil {
		fmt.Println("failed to read lines", err)
	}

	fmt.Println("Part 1: ", day1.Part1(lines))
	fmt.Println("Part 2: ", day1.Part2(lines))
}
