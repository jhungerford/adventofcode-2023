package main

import (
	"fmt"
	"github.com/jhungerford/adventofcode-2023/internal/day9"
)

func main() {
	readings, err := day9.LoadReadings("input/day9.txt")
	if err != nil {
		fmt.Printf("failed to load readings: %v", err)
	}

	fmt.Println("Part 1:", day9.Part1(readings))
	fmt.Println("Part 2:", day9.Part2(readings))
}
