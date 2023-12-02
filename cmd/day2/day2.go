package main

import (
	"fmt"
	"github.com/jhungerford/adventofcode-2023/internal/day2"
)

func main() {
	games, err := day2.LoadGames("input/day2.txt")
	if err != nil {
		fmt.Printf("failed to load games: %v", err)
		return
	}

	fmt.Println("Part1:", day2.Part1(games))
}
