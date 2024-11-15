package day14

import (
	"testing"
)

func TestPart1(t *testing.T) {
	t.Parallel()

	platform, err := LoadPlatform("day14_sample.txt")
	if err != nil {
		t.Fatalf("failed to load platform: %v", err)
	}

	platform.Tilt(platform.North())

	actualLoad, expectedLoad := platform.Load(), 136

	if actualLoad != expectedLoad {
		t.Errorf("Part1() = %v, want %v\nplatform:\n%s", actualLoad, expectedLoad, platform)
	}
}

func TestPart2(t *testing.T) {
	t.Parallel()

	platform, err := LoadPlatform("day14_sample.txt")
	if err != nil {
		t.Fatalf("failed to load platform: %v", err)
	}

	platform.NumCycles(1000000000)

	actualLoad, expectedLoad := platform.Load(), 64

	if actualLoad != expectedLoad {
		t.Errorf("Part2() = %v, want %v\nplatform (%s):\n%s", actualLoad, expectedLoad, platform.hash(), platform)
	}
}
