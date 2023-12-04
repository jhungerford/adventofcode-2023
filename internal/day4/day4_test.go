package day4

import "testing"

func TestPart1(t *testing.T) {
	t.Parallel()

	cards, err := LoadCards("day4_sample.txt")
	if err != nil {
		t.Fatalf("failed to load cards: %v", err)
	}

	want := 13

	if got := Part1(cards); got != want {
		t.Errorf("Part1() = %v, want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	t.Parallel()

	cards, err := LoadCards("day4_sample.txt")
	if err != nil {
		t.Fatalf("failed to load cards: %v", err)
	}

	want := 30

	if got := Part2(cards); got != want {
		t.Errorf("Part2() = %v, want %v", got, want)
	}
}
