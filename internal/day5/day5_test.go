package day5

import (
	"fmt"
	"testing"
)

func TestPart1(t *testing.T) {
	t.Parallel()

	almanac, err := LoadAlmanac("day5_sample.txt")
	if err != nil {
		t.Fatalf("failed to load almanac: %v", err)
	}

	want := 35

	if got := Part1(almanac); got != want {
		t.Errorf("Part1() = %v, want %v", got, want)
	}
}

func Test_seedLocations(t *testing.T) {
	t.Parallel()

	almanac, err := LoadAlmanac("day5_sample.txt")
	if err != nil {
		t.Fatalf("failed to load almanac: %v", err)
	}

	tests := []struct {
		seed         int
		wantLocation int
	}{
		{79, 82},
		{14, 43},
		{55, 86},
		{13, 35},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("seed %d", tt.seed), func(t *testing.T) {
			gotLocation := almanac.resolveLocation(tt.seed)
			if tt.wantLocation != gotLocation {
				t.Errorf("resolveLocation(%d) got = %v, want %v", tt.seed, gotLocation, tt.wantLocation)
			}
		})
	}
}
