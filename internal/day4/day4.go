package day4

import (
	"fmt"
	"github.com/jhungerford/adventofcode-2023/internal/util"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

// LoadCards loads cards from the given file.
func LoadCards(filename string) ([]Card, error) {
	return util.ParseInputLines(filename, func(line string) (Card, error) {
		// Line looks like 'Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53'

		re := regexp.MustCompile("  +")
		line = re.ReplaceAllString(line, " ")

		numIndex := strings.Index(line, ":")
		separatorIndex := strings.Index(line, "|")

		numStr := line[5:numIndex]

		num, err := strconv.Atoi(numStr)
		if err != nil {
			return Card{}, fmt.Errorf("failed to parse card number from '%s': %w", numStr, err)
		}

		winningNumStrs := strings.Split(strings.TrimSpace(line[numIndex+1:separatorIndex]), " ")
		winning := make([]int, 0, len(winningNumStrs))

		for _, winningNumStr := range winningNumStrs {
			winningNum, werr := strconv.Atoi(winningNumStr)
			if werr != nil {
				return Card{}, fmt.Errorf("failed to parse winning number from '%s': %w", winningNumStr, werr)
			}

			winning = append(winning, winningNum)
		}

		haveNumStrs := strings.Split(strings.TrimSpace(line[separatorIndex+1:]), " ")
		have := make([]int, 0, len(haveNumStrs))

		for _, haveNumStr := range haveNumStrs {
			haveNum, werr := strconv.Atoi(haveNumStr)
			if werr != nil {
				return Card{}, fmt.Errorf("failed to parse have number from '%s': %w", haveNumStr, werr)
			}

			have = append(have, haveNum)
		}

		return Card{
			num:     num,
			winning: winning,
			have:    have,
		}, nil
	})
}

// Part1 calculates the total scores in the pile of cards.
func Part1(cards []Card) int {
	score := 0

	for _, card := range cards {
		score += card.score()
	}

	return score
}

type Card struct {
	num     int
	winning []int
	have    []int
}

// score returns this card's score.  The first matching winning number is worth one point, and each match
// after the first doubles the card's score.
func (c Card) score() int {
	score := 0

	for _, n := range c.have {
		if slices.Contains(c.winning, n) {
			if score == 0 {
				score = 1
			} else {
				score *= 2
			}
		}
	}

	return score
}