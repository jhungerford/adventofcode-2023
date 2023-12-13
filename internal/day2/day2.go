package day2

import (
	"errors"
	"fmt"
	"github.com/jhungerford/adventofcode-2023/internal/util"
	"strconv"
	"strings"
)

type Game struct {
	num   int
	pulls []pull
}

type pull struct {
	red, green, blue int
}

// LoadGames loads a list of games from the given input file.
func LoadGames(file string) ([]Game, error) {
	return util.ParseInputLines(file, true, parseGame)
}

func parseGame(line string) (Game, error) {
	// Game looks like 'Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green'

	separatorIdx := strings.Index(line, ":")
	if separatorIdx == -1 {
		return Game{}, errors.New("invalid game - no ':' separator")
	}

	gameNumStr := line[5:separatorIdx]

	num, err := strconv.Atoi(gameNumStr)
	if err != nil {
		return Game{}, fmt.Errorf("invalid game number '%s'", gameNumStr)
	}

	pullStrs := strings.Split(line[separatorIdx+1:], ";")

	pulls := make([]pull, 0, len(pullStrs))

	for _, pullStr := range pullStrs {
		colors := strings.Split(strings.TrimSpace(pullStr), ",")

		p := pull{}

		for _, color := range colors {
			color = strings.TrimSpace(color)

			spaceIdx := strings.Index(color, " ")
			if spaceIdx == -1 {
				return Game{}, errors.New(fmt.Sprintf(
					"invalid color - '%s' does not have a space between the count and the color", color))
			}

			count, cerr := strconv.Atoi(color[:spaceIdx])
			if cerr != nil {
				return Game{}, fmt.Errorf("cannot parse '%s' as a color count: %v", color[:spaceIdx], err)
			}

			switch color[spaceIdx+1:] {
			case "red":
				p.red = count
			case "green":
				p.green = count
			case "blue":
				p.blue = count
			}

		}

		pulls = append(pulls, p)
	}

	game := Game{
		num:   num,
		pulls: pulls,
	}

	return game, nil
}

// Part1 returns the sum of ids of games that are possible if the bag is loaded with 12 red, 13 green, and 14 blue cubes.
func Part1(games []Game) int {
	var sum int

	for _, g := range games {
		possible := true

		for _, p := range g.pulls {
			possible = possible && p.red <= 12 && p.green <= 13 && p.blue <= 14
		}

		if possible {
			sum += g.num
		}
	}

	return sum
}

// Part2 finds the power set of each game, which is the product of the minimum number of cubes, and returns the sum.
func Part2(games []Game) int {
	var sum int

	for _, g := range games {
		need := pull{}

		for _, p := range g.pulls {
			need.red = max(need.red, p.red)
			need.green = max(need.green, p.green)
			need.blue = max(need.blue, p.blue)
		}

		sum += need.red * need.green * need.blue
	}

	return sum
}
