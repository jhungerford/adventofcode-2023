package day7

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPart1(t *testing.T) {
	t.Parallel()

	hands, err := LoadHands("day7_sample.txt")
	if err != nil {
		t.Fatalf("failed to load hands: %v", err)
	}

	origHands, err := LoadHands("day7_sample.txt")
	if err != nil {
		t.Fatalf("failed to load hands: %v", err)
	}

	want := 6440

	if got := Part1(hands); got != want {
		t.Errorf("Part1() = %v, want %v", got, want)
	}

	if !reflect.DeepEqual(hands, origHands) {
		t.Errorf("Part1() hands mutated.")
	}
}

func TestPart2(t *testing.T) {
	t.Parallel()

	hands, err := LoadHands("day7_sample.txt")
	if err != nil {
		t.Fatalf("failed to load hands: %v", err)
	}

	origHands, err := LoadHands("day7_sample.txt")
	if err != nil {
		t.Fatalf("failed to load hands: %v", err)
	}

	want := 5905

	if got := Part2(hands); got != want {
		t.Errorf("Part2() = %v, want %v", got, want)
	}

	if !reflect.DeepEqual(hands, origHands) {
		t.Errorf("Part2() hands mutated.")
	}
}

func TestHand_plainHandType(t *testing.T) {
	t.Parallel()

	tests := []struct {
		cards string
		bid   int
		want  string
	}{
		{"32T3K", 765, "one-pair"},
		{"T55J5", 684, "three-of-a-kind"},
		{"KK677", 28, "two-pair"},
		{"KTJJT", 220, "two-pair"},
		{"QQQJA", 483, "three-of-a-kind"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s %d", tt.cards, tt.bid), func(t *testing.T) {
			h := Hand{
				cards: []byte(tt.cards),
				bid:   tt.bid,
			}

			if got := h.plainHandType(); got != tt.want {
				t.Errorf("plainHandType(%s %d) = %v, want %v", tt.cards, tt.bid, got, tt.want)
			}

			if !reflect.DeepEqual([]byte(tt.cards), h.cards) {
				t.Errorf("plainHandType(%s %d) changed the cards order", tt.cards, tt.bid)
			}
		})
	}
}

func TestHand_jokerHandType(t *testing.T) {
	t.Parallel()

	tests := []struct {
		cards string
		bid   int
		want  string
	}{
		{"32T3K", 765, "one-pair"},
		{"T55J5", 684, "four-of-a-kind"},
		{"KK677", 28, "two-pair"},
		{"KTJJT", 220, "four-of-a-kind"},
		{"QQQJA", 483, "four-of-a-kind"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s %d", tt.cards, tt.bid), func(t *testing.T) {
			h := Hand{
				cards: []byte(tt.cards),
				bid:   tt.bid,
			}

			if got := h.jokerHandType(); got != tt.want {
				t.Errorf("jokerHandType(%s %d) = %v, want %v", tt.cards, tt.bid, got, tt.want)
			}

			if !reflect.DeepEqual([]byte(tt.cards), h.cards) {
				t.Errorf("jokerHandType(%s %d) changed the cards order", tt.cards, tt.bid)
			}
		})
	}
}
