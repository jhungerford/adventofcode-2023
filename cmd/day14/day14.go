package main

import (
	"fmt"
	"github.com/jhungerford/adventofcode-2023/internal/day14"
)

func main() {
	platform, err := day14.LoadPlatform("input/day14.txt")
	if err != nil {
		fmt.Printf("failed to load platform: %v", err)
		return
	}

	fmt.Println("Part 1:", day14.Part1(platform))
}
