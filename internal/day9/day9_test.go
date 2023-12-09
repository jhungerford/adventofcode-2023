package day9

import (
	"testing"
)

func TestPart1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		filename string
		want     int
	}{
		{"day9_sample.txt", 114},
	}

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			readings, err := LoadReadings(tt.filename)
			if err != nil {
				t.Fatalf("failed to load readings from '%s': %v", tt.filename, err)
			}

			if got := Part1(readings); got != tt.want {
				t.Errorf("Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}
