package main

import (
	"fmt"
	"github.com/jhungerford/adventofcode-2023/internal/day13"
)

func main() {
	notes, err := day13.LoadNotes("input/day13.txt")
	if err != nil {
		fmt.Printf("failed to load notes: %v", err)
	}

	fmt.Println("Part 1:", day13.Part1(notes))
}
