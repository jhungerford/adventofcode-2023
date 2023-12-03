package day3

import (
	"fmt"
	"github.com/jhungerford/adventofcode-2023/internal/util"
)

// LoadSchematic loads a schematic from the given file.  The schematic contains numbers, symbols, and periods.
func LoadSchematic(filename string) (Schematic, error) {
	lines, err := util.ReadInputLines(filename)
	if err != nil {
		return Schematic{}, fmt.Errorf("failed to read schematic from %s: %v", filename, err)
	}

	s := Schematic{
		numbers: map[Position]int{},
		symbols: map[Position]byte{},
	}

	for row, line := range lines {
		currentNum := -1
		currentNumPos := Position{}

		for col, value := range []byte(line) {
			if value >= '0' && value <= '9' {
				if currentNum == -1 {
					currentNum = int(value - '0')
					currentNumPos = Position{row, col}
				} else {
					currentNum = currentNum*10 + int(value-'0')
				}
			} else {
				if currentNum != -1 {
					s.numbers[currentNumPos] = currentNum
					currentNum = -1
				}

				if value != '.' {
					s.symbols[Position{row, col}] = value
				}
			}
		}

		if currentNum != -1 {
			s.numbers[currentNumPos] = currentNum
		}
	}

	return s, nil
}

// Part1 sums the numbers that are adjacent to at least one symbol, even diagonally.
func Part1(schematic Schematic) int {
	sum := 0

	for pos, number := range schematic.numbers {
		if schematic.hasAdjacentSymbol(pos, number) {
			sum += number
		}
	}

	return sum
}

// Part2 sums the gear ratios in the engine, where a gear ratio is the product of the two numbers adjacent to a '*'.
// '*'s that don't have exactly two adjacent numbers don't have a gear ratio.
func Part2(schematic Schematic) int {
	gearNumbers := map[Position][]int{}

	for numberPos, number := range schematic.numbers {
		for symbolPos, symbol := range schematic.getAdjacentSymbols(numberPos, number) {
			if symbol == '*' {
				gearNumbers[symbolPos] = append(gearNumbers[symbolPos], number)
			}
		}
	}

	sum := 0

	for _, numbers := range gearNumbers {
		if len(numbers) == 2 {
			sum += numbers[0] * numbers[1]
		}
	}

	return sum
}

// Position is a row and column in the schematic.  It can indicate the start of a number, or the position of a symbol.
type Position struct {
	row, col int
}

// Schematic is a parsed schematic containing numbers, symbols, and their positions
type Schematic struct {
	numbers map[Position]int
	symbols map[Position]byte
}

// hasAdjacentSymbol returns whether there's at least one adjacent symbol to this number.
// Symbols can be diagonal to the number.
func (s Schematic) hasAdjacentSymbol(pos Position, number int) bool {
	for plusRow := -1; plusRow <= 1; plusRow++ {
		for plusCol := -1; plusCol <= len(fmt.Sprintf("%d", number)); plusCol++ {
			checkPos := Position{row: pos.row + plusRow, col: pos.col + plusCol}
			if _, ok := s.symbols[checkPos]; ok {
				return true
			}
		}
	}

	return false
}

// getAdjacentSymbols returns a map of symbol positions to symbol that are adjacent to this number.
func (s Schematic) getAdjacentSymbols(pos Position, number int) map[Position]byte {
	adjacentSymbols := map[Position]byte{}

	for plusRow := -1; plusRow <= 1; plusRow++ {
		for plusCol := -1; plusCol <= len(fmt.Sprintf("%d", number)); plusCol++ {
			checkPos := Position{row: pos.row + plusRow, col: pos.col + plusCol}
			if symbol, ok := s.symbols[checkPos]; ok {
				adjacentSymbols[checkPos] = symbol
			}
		}
	}

	return adjacentSymbols
}
