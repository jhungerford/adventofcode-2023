package day10

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPart1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		filename string
		want     int
	}{
		{"day10_sample.txt", 4},
		{"day10_sample2.txt", 8},
	}

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			m, err := LoadMap(tt.filename)
			if err != nil {
				t.Fatalf("failed to load map from '%s': %v", tt.filename, err)
			}

			if got := Part1(m); got != tt.want {
				t.Errorf("Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_neighbors(t *testing.T) {
	t.Parallel()

	filename := "day10_sample.txt"

	m, err := LoadMap(filename)
	if err != nil {
		t.Fatalf("failed to load map from '%s': %v", filename, err)
	}

	startTile := m.startTile()
	if startTile != 'F' {
		t.Errorf("startTile() = %c, want %c", startTile, 'F')
	}

	tests := []struct {
		pos  position
		want []position
	}{
		{position{row: 1, col: 1}, []position{{1, 2}, {2, 1}}},
		{position{row: 1, col: 2}, []position{{1, 1}, {1, 3}}},
		{position{row: 1, col: 3}, []position{{1, 2}, {2, 3}}},

		{position{row: 2, col: 1}, []position{{1, 1}, {3, 1}}},
		{position{row: 2, col: 3}, []position{{1, 3}, {3, 3}}},

		{position{row: 3, col: 1}, []position{{3, 2}, {2, 1}}},
		{position{row: 3, col: 2}, []position{{3, 1}, {3, 3}}},
		{position{row: 3, col: 3}, []position{{3, 2}, {2, 3}}},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("neighbors(%d,%d)", tt.pos.row, tt.pos.col), func(t *testing.T) {
			wantMap := map[position]interface{}{}
			for _, wantPos := range tt.want {
				wantMap[wantPos] = nil
			}

			got := m.neighbors(tt.pos)
			gotMap := map[position]interface{}{}
			for _, gotPos := range got {
				gotMap[gotPos] = nil
			}

			if !reflect.DeepEqual(wantMap, gotMap) {
				t.Errorf("neighbors(%d,%d) = %v, want %v", tt.pos.row, tt.pos.col, gotMap, wantMap)
			}
		})
	}
}
