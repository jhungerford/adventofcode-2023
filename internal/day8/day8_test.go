package day8

import (
	"testing"
)

func TestPart1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		filename string
		want     int
	}{
		{"day8_sample.txt", 2},
		{"day8_sample2.txt", 6},
	}

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			network, err := LoadNetwork(tt.filename)
			if err != nil {
				t.Fatalf("failed to load network from '%s': %v", tt.filename, err)
			}

			if got := Part1(network); got != tt.want {
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
		{"day8_sample3.txt", 6},
	}

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			network, err := LoadNetwork(tt.filename)
			if err != nil {
				t.Fatalf("failed to load network from '%s': %v", tt.filename, err)
			}

			if got := Part2(network); got != tt.want {
				t.Errorf("Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
