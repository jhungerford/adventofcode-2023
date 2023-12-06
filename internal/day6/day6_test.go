package day6

import (
	"fmt"
	"testing"
)

func TestPart1(t *testing.T) {
	t.Parallel()

	races, err := LoadRaces("day6_sample.txt")
	if err != nil {
		t.Fatalf("failed to load races: %v", err)
	}

	want := 288

	if got := Part1(races); got != want {
		t.Errorf("Part1() = %v, want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	t.Parallel()

	races, err := LoadRaces("day6_sample.txt")
	if err != nil {
		t.Fatalf("failed to load races: %v", err)
	}

	want := 71503

	if got := Part2(races); got != want {
		t.Errorf("Part2() = %v, want %v", got, want)
	}
}

func Test_wins(t *testing.T) {
	t.Parallel()

	races, err := LoadRaces("day6_sample.txt")
	if err != nil {
		t.Fatalf("failed to load races: %v", err)
	}

	tests := []struct {
		race int
		want int
	}{
		{0, 4},
		{1, 8},
		{2, 9},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("race %d", tt.race), func(t *testing.T) {
			got := wins(races.times[tt.race], races.distances[tt.race])
			if tt.want != got {
				t.Errorf("races.wins(%d) got = %v, want %v", tt.race, got, tt.want)
			}
		})
	}
}
