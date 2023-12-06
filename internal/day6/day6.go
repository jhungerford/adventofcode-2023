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
		w := wins(races.times[i], races.distances[i])
		if w > 0 {
			if product == 0 {
				product = 1
			}

			product *= w
		}
	}

	return product
}

// Part2 returns the number of ways to win a single race with times and distances given by removing all of the spaces
// from the times in the input.
func Part2(races Races) int {
	time := combineNumbers(races.times)
	distance := combineNumbers(races.distances)

	return wins(time, distance)
}

// combineDigits concatenates the numbers to form a single number.  For example, [1, 23, 4] becomes 1234.
func combineNumbers(nums []int) int {
	s := ""

	for _, num := range nums {
		s = fmt.Sprintf("%s%d", s, num)
	}

	return util.MustAtoi(s)
}

// wins returns the number of ways to win race at the given index.
func wins(time, dist int) int {
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

	raceTime := float64(time)
	raceDist := float64(dist)

	minHoldTime := math.Ceil(0.5 * (raceTime - math.Sqrt(math.Pow(raceTime, 2)-4.0*raceDist)))
	maxHoldTime := math.Floor(0.5 * (raceTime + math.Sqrt(math.Pow(raceTime, 2)-4.0*raceDist)))

	// Have to win the race, not just tie the distance
	if (time-int(minHoldTime))*int(minHoldTime) == dist {
		minHoldTime += 1
	}

	if (time-int(maxHoldTime))*int(maxHoldTime) == dist {
		maxHoldTime -= 1
	}

	return int(maxHoldTime) - int(minHoldTime) + 1
}
