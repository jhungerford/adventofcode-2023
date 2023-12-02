package day2

import (
	"reflect"
	"testing"
)

func Test_Part1(t *testing.T) {
	t.Parallel()

	games, err := LoadGames("day2_sample.txt")
	if err != nil {
		t.Fatalf("failed to load games: %v", err)
	}

	want := 8

	if got := Part1(games); got != want {
		t.Errorf("Part1() = %v, want %v", got, want)
	}
}

func Test_parseGame(t *testing.T) {
	t.Parallel()

	tests := []struct {
		line string
		want Game
	}{
		{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", Game{
			num: 1,
			pulls: []pull{
				{red: 4, blue: 3},
				{red: 1, green: 2, blue: 6},
				{green: 2},
			},
		}},
		{"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", Game{
			num: 2,
			pulls: []pull{
				{green: 2, blue: 1},
				{red: 1, green: 3, blue: 4},
				{green: 1, blue: 1},
			},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.line, func(t *testing.T) {
			got, err := parseGame(tt.line)
			if err != nil {
				t.Fatalf("parseGame() error = %v", err)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseGame() got = %v, want %v", got, tt.want)
			}
		})
	}
}
