package day11

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
		{"day11_sample.txt", 374},
	}

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			img, err := LoadImage(tt.filename)
			if err != nil {
				t.Fatalf("failed to load image from '%s': %v", tt.filename, err)
			}

			if got := Part1(img); got != tt.want {
				t.Errorf("Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	t.Parallel()

	tests := []struct {
		filename  string
		expansion int
		want      int
	}{
		{"day11_sample.txt", 10, 1030},
		{"day11_sample.txt", 100, 8410},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s: %d", tt.filename, tt.expansion), func(t *testing.T) {
			img, err := LoadImage(tt.filename)
			if err != nil {
				t.Fatalf("failed to load image from '%s': %v", tt.filename, err)
			}

			if got := img.galaxyDistances(tt.expansion); got != tt.want {
				t.Errorf("Part2(%d) = %v, want %v", tt.expansion, got, tt.want)
			}
		})
	}
}
