package day13

import (
	"fmt"
	"testing"
)

func TestPart1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		filename string
		want     int
	}{
		{"day13_sample.txt", 405},
	}

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			notes, err := LoadNotes(tt.filename)
			if err != nil {
				t.Fatalf("failed to load notes from '%s': %v", tt.filename, err)
			}

			if got := Part1(notes); got != tt.want {
				t.Errorf("Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	t.Parallel()

	tests := []struct {
		filename string
		want     int
	}{
		{"day13_sample.txt", 400},
	}

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			notes, err := LoadNotes(tt.filename)
			if err != nil {
				t.Fatalf("failed to load notes from '%s': %v", tt.filename, err)
			}

			if got := Part2(notes); got != tt.want {
				t.Errorf("Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reflections(t *testing.T) {
	t.Parallel()

	tests := []struct {
		wantAbove, wantLeft int
	}{
		{0, 5},
		{4, 0},
	}

	filename := "day13_sample.txt"

	notes, err := LoadNotes(filename)
	if err != nil {
		t.Fatalf("failed to load notes from '%s': %v", filename, err)
	}

	if len(tests) != len(notes.patterns) {
		t.Fatalf("Wrong number of tests.  %d patterns, %d tests", len(notes.patterns), len(tests))
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("pattern %d", i), func(t *testing.T) {
			got := notes.patterns[i].findReflection()

			if got.above != tt.wantAbove {
				t.Errorf("findReflection().above = %d, want %d", got.above, tt.wantAbove)
			}

			if got.left != tt.wantLeft {
				t.Errorf("findReflection().left = %d, want %d", got.left, tt.wantLeft)
			}
		})
	}
}
