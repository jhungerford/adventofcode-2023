package main

import (
	"fmt"
	"github.com/jhungerford/adventofcode-2023/internal/day7"
)

func main() {
	hands, err := day7.LoadHands("input/day7.txt")
	if err != nil {
		fmt.Printf("failed to load hands: %v", err)
	}

	fmt.Println("Part 1:", day7.Part1(hands))
}
