package main

import (
	"fmt"
	"github.com/jhungerford/adventofcode-2023/internal/day11"
)

func main() {
	img, err := day11.LoadImage("input/day11.txt")
	if err != nil {
		fmt.Printf("failed to load image: %v", err)
	}

	fmt.Println("Part 1:", day11.Part1(img))
	fmt.Println("Part 2:", day11.Part2(img))
}
