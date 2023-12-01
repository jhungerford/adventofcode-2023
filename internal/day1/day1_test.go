package day1

import "testing"

func Test_Part1(t *testing.T) {
	t.Parallel()

	want := 142

	if got := Part1("input/day1_sample.txt"); got != want {
		t.Errorf("Part1() = %v, want %v", got, want)
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
