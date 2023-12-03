package day3

import "testing"

func TestPart1(t *testing.T) {
	t.Parallel()

	schematic, err := LoadSchematic("day3_sample.txt")
	if err != nil {
		t.Fatalf("failed to load schematic: %v", err)
	}

	want := 4361

	if got := Part1(schematic); got != want {
		t.Errorf("Part1() = %v, want %v", got, want)
	}
}
