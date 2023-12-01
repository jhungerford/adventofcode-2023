package day1

import (
	"github.com/jhungerford/adventofcode-2023/internal/util"
	"unicode"
)

// Part1 sums calibration values in the given lines, where a calibration value is the first and last digit in a line.
// Panics if an error occurs while reading the input file.
func Part1(lines []string) int {
	sum := 0

	for _, line := range lines {
		sum += calibrationValue(line)
	}

	return sum
}

// Part2 sums real values in the given lines, where a real value is the first and last digit in a line.
// Digits can either be numeric, or spelled out with letters like 'one'
// Panics if an error occurs while reading the input file.
func Part2(lines []string) int {
	sum := 0

	for _, line := range lines {
		sum += realValue(line)
	}

	return sum
}

// calibrationValue returns the first and last numeric digit in the given line.
func calibrationValue(line string) int {
	return lineValue(line, calibrationNumber)
}

// realValue returns the first and last digit in the given line, where a digit can either be numeric
// or spelled out like 'one'.
func realValue(line string) int {
	return lineValue(line, realNumber)
}

// lineValue returns the value of the line, using the given strategy to extract numbers.
func lineValue(line string, toNumber func(string, int, rune) (int, bool)) int {
	firstDigit, lastDigit := -1, -1

	for i, r := range line {
		if number, ok := toNumber(line, i, r); ok {
			if firstDigit == -1 {
				firstDigit = number
			}

			lastDigit = number
		}
	}

	return firstDigit*10 + lastDigit
}

// calibrationNumber returns the digit at the given index in the line.
func calibrationNumber(line string, i int, r rune) (int, bool) {
	if unicode.IsNumber(r) {
		return util.MustAtoi(string(r)), true
	}

	return -1, false
}

// realNumber returns the digit or spelled out number at the given index in the line.
func realNumber(line string, i int, r rune) (int, bool) {
	numberStrs := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	if unicode.IsNumber(r) {
		return util.MustAtoi(string(r)), true
	}

	for n, numberStr := range numberStrs {
		endIndex := i + len(numberStr)

		if endIndex <= len(line) && numberStr == string([]rune(line)[i:endIndex]) {
			return n, true
		}
	}

	return -1, false
}
