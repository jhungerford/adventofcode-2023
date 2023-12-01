package day1

import (
	"github.com/jhungerford/adventofcode-2023/internal/util"
	"testing"
)

func Test_Part1(t *testing.T) {
	t.Parallel()

	lines, err := util.ReadInputLines("day1_sample.txt")
	if err != nil {
		t.Fatalf("failed to read lines: %v", err)
	}

	want := 142

	if got := Part1(lines); got != want {
		t.Errorf("Part1() = %v, want %v", got, want)
	}
}

func Test_Part2(t *testing.T) {
	t.Parallel()

	lines, err := util.ReadInputLines("day1_sample2.txt")
	if err != nil {
		t.Fatalf("failed to read lines: %v", err)
	}

	want := 281

	if got := Part2(lines); got != want {
		t.Errorf("Part2() = %v, want %v", got, want)
	}
}

func Test_calibrationValue(t *testing.T) {
	t.Parallel()

	tests := []struct {
		line string
		want int
	}{
		{"1abc2", 12},
		{"pqr3stu8vwx", 38},
		{"a1b2c3d4e5f", 15},
		{"treb7uchet", 77},
	}

	for _, tt := range tests {
		if got := calibrationValue(tt.line); got != tt.want {
			t.Errorf("calibrationValue('%s') = %v, want %v", tt.line, got, tt.want)
		}
	}
}

func Test_realValue(t *testing.T) {
	t.Parallel()

	tests := []struct {
		line string
		want int
	}{
		{"two1nine", 29},
		{"eightwothree", 83},
		{"abcone2threexyz", 13},
		{"xtwone3four", 24},
		{"4nineeightseven2", 42},
		{"zoneight234", 14},
		{"7pqrstsixteen", 76},
	}

	for _, tt := range tests {
		if got := realValue(tt.line); got != tt.want {
			t.Errorf("realValue('%s') = %v, want %v", tt.line, got, tt.want)
		}
	}
}
