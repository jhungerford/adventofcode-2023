package day1

import (
	"github.com/jhungerford/adventofcode-2023/internal/util"
	"unicode"
)

// Part1 sums calibration values in the given file, where a calibration value is the first and last digit in a line.
// Panics if an error occurs while reading the input file.
func Part1(file string) int {
	lines, err := util.ReadFileLines(file)
	if err != nil {
		panic(err)
	}

	sum := 0

	for _, line := range lines {
		sum += calibrationValue(line)
	}

	return sum
}

// calibrationValue returns the first and last digit of the given line.
func calibrationValue(line string) int {
	firstDigit, lastDigit := -1, -1

	for _, r := range line {
		if unicode.IsNumber(r) {
			lastDigit = util.MustAtoi(string(r))
			if firstDigit == -1 {
				firstDigit = lastDigit
			}
		}
	}

	return firstDigit*10 + lastDigit
}
