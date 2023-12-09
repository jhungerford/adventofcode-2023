package day9

import "github.com/jhungerford/adventofcode-2023/internal/util"

func LoadReadings(filename string) (Readings, error) {
	values, err := util.ParseInputLines(filename, util.IntList)
	if err != nil {
		return Readings{}, err
	}

	return Readings{values}, nil
}

// Part1 returns the sum of the extrapolated next value in each history reading.
func Part1(readings Readings) int {
	sum := 0

	for _, history := range readings.values {
		delta := history
		allZero := false

		for !allZero {
			sum += delta[len(delta)-1]

			delta, allZero = difference(delta)
		}

	}

	return sum
}

// difference returns the difference between each value, and whether all differences were 0.
// The returned list has one fewer element than the values since differences are between values.
func difference(values []int) ([]int, bool) {
	differences := make([]int, 0, len(values)-1)
	allZero := true

	for i, value := range values[1:] {
		diff := value - values[i]
		differences = append(differences, diff)

		if diff != 0 {
			allZero = false
		}
	}

	return differences, allZero
}

type Readings struct {
	values [][]int
}
