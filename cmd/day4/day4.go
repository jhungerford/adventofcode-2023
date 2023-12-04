package main

import (
	"fmt"
	"github.com/jhungerford/adventofcode-2023/internal/day4"
)

func main() {
	cards, err := day4.LoadCards("input/day4.txt")
	if err != nil {
		fmt.Printf("failed to load cards: %v", err)
		return
	}

	fmt.Println("Part 1:", day4.Part1(cards))
}
