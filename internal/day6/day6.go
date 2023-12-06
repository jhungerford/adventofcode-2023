package day6

import (
	"errors"
	"fmt"
	"github.com/jhungerford/adventofcode-2023/internal/util"
	"math"
)

type Races struct {
	times     []int
	distances []int
}

func LoadRaces(filename string) (Races, error) {
	lines, err := util.ReadInputLines(filename)
	if err != nil {
		return Races{}, fmt.Errorf("failed to read race lines from '%s': %w", filename, err)
	}

	// File contains two lines like:
	// Time:      7  15   30
	// Distance:  9  40  200
	if len(lines) < 2 {
		return Races{}, errors.New("input file too short")
	}

	times, err := util.IntList(util.CondenseSpaces(lines[0][len("Time:")+1:]))
	if err != nil {
		return Races{}, fmt.Errorf("failed to read times from file '%s': %v", filename, err)
	}

	distances, err := util.IntList(util.CondenseSpaces(lines[1][len("Distance:")+1:]))
	if err != nil {
		return Races{}, fmt.Errorf("failed to read distances from file '%s': %v", filename, err)
	}

	if len(times) != len(distances) {
		return Races{}, errors.New("mismatched times and distances")
	}

	return Races{times, distances}, nil
}

// Part1 returns the margin that you have in the given races, which is the product of the number of ways you can beat
// the record in each race.
func Part1(races Races) int {
	product := 0

	for i := 0; i < len(races.times); i++ {
		wins := races.wins(i)
		if wins > 0 {
			if product == 0 {
				product = 1
			}

			product *= wins
		}
	}

	return product
}

// wins returns the number of ways to win race at the given index.
func (races Races) wins(i int) int {
	// The boat has a starting speed of zero mm per ms, which increases by 1 mm/ms for every ms you hold the button.
	// The shortest and longest times you can hold the button to win are given by:
	//
	//   d = (T-h) * h
	//
	// solving for h using the quadratic equation, we get:
	//
	//     T ± √(T² - 4*d)
	// h = ---------------
	//            2
	//
	// solving for h where d=D, we get the min and max hold times to win the race
	//
	// where:
	//   T = race time
	//   D = race distance record
	//   d = boat distance
	//   h = hold time

	raceTime := float64(races.times[i])
	raceDist := float64(races.distances[i])

	minHoldTime := math.Ceil(0.5 * (raceTime - math.Sqrt(math.Pow(raceTime, 2)-4.0*raceDist)))
	maxHoldTime := math.Floor(0.5 * (raceTime + math.Sqrt(math.Pow(raceTime, 2)-4.0*raceDist)))

	// Have to win the race, not just tie the distance
	if (races.times[i]-int(minHoldTime))*int(minHoldTime) == races.distances[i] {
		minHoldTime += 1
	}

	if (races.times[i]-int(maxHoldTime))*int(maxHoldTime) == races.distances[i] {
		maxHoldTime -= 1
	}

	return int(maxHoldTime) - int(minHoldTime) + 1
}
