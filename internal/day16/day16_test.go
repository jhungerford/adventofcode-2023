package day16

import "testing"

func TestPart1(t *testing.T) {
	t.Parallel()

	m, err := LoadMap("day16_sample.txt")
	if err != nil {
		t.Fatalf("failed to load map: %v", err)
	}

	actual, expected := m.part1Energize(), 46

	if actual != expected {
		t.Errorf("Part1() = %v, want %v", actual, expected)
	}
}

func TestPart2(t *testing.T) {
	t.Parallel()

	m, err := LoadMap("day16_sample.txt")
	if err != nil {
		t.Fatalf("failed to load map: %v", err)
	}

	actual, expected := m.part2Energize(), 51

	if actual != expected {
		t.Errorf("Part2() = %v, want %v", actual, expected)
	}
}
