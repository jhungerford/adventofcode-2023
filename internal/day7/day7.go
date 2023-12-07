package day7

import (
	"fmt"
	"github.com/jhungerford/adventofcode-2023/internal/util"
	"slices"
	"strconv"
)

func LoadHands(filename string) ([]Hand, error) {
	return util.ParseInputLines(filename, func(line string) (Hand, error) {
		// Hand looks like '32T3K 765', where the first five digits are cards, and the second number is a bid.
		bid, err := strconv.Atoi(line[6:])
		if err != nil {
			return Hand{}, fmt.Errorf("failed to parse bid from '%s': %w", line[6:], err)
		}

		return Hand{
			cards: []byte(line)[:5],
			bid:   bid,
		}, nil
	})
}

// Part1 returns the total winnings
func Part1(hands []Hand) int {
	return winnings(hands, part1CardStrengths, Hand.plainHandType)
}

// Part2 returns the total winnings where 'J' cards are treated as jokers.  Jokers are wildcards that make the
// strongest possible hand, but have the weakest strength when compared individually.
func Part2(hands []Hand) int {
	return winnings(hands, part2CardStrengths, Hand.jokerHandType)
}

// winnings returns the total winnings of this set of hands by adding up each hand's bid multiplied by its rank.
// Hands have a type like 'full house' that determines its rank, and a bid.
// cardStrengths and handType determine the game's rules.
func winnings(hands []Hand, cardStrengths []byte, handType func(Hand) string) int {
	sortedHands := make([]Hand, len(hands))
	copy(sortedHands, hands)

	cmp := cmpHands(cardStrengths, handType)
	slices.SortFunc(sortedHands, cmp)

	total := 0
	rank := 1

	for i, hand := range sortedHands[:len(sortedHands)-1] {
		total += rank * hand.bid

		if cmp(hand, sortedHands[i+1]) < 0 {
			rank++
		}
	}

	total += rank * sortedHands[len(sortedHands)-1].bid

	return total
}

type Hand struct {
	cards []byte
	bid   int
}

var part1CardStrengths = []byte{'2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A'}
var part2CardStrengths = []byte{'J', '2', '3', '4', '5', '6', '7', '8', '9', 'T', 'Q', 'K', 'A'}
var typeStrengths = []string{"high-card", "one-pair", "two-pair", "three-of-a-kind", "full-house", "four-of-a-kind", "five-of-a-kind"}

// cmpHands returns the ordering of the given hands, from weakest to strongest.
// Hands are primarily sorted by their type, like 'Full House'.  If two hands have the same rank, cards
// are compared from left to right.
func cmpHands(cardStrengths []byte, handType func(Hand) string) func(a, b Hand) int {
	return func(a, b Hand) int {
		aType := slices.Index(typeStrengths, handType(a))
		bType := slices.Index(typeStrengths, handType(b))

		if aType != bType {
			return aType - bType
		}

		for i, aCard := range a.cards {
			aStr := slices.Index(cardStrengths, aCard)
			bStr := slices.Index(cardStrengths, b.cards[i])

			if aStr != bStr {
				return aStr - bStr
			}
		}

		return 0
	}
}

func (h Hand) plainHandType() string {
	// sort the cards, and assign them numbers
	cards := make([]byte, len(h.cards))
	copy(cards, h.cards)
	slices.Sort(cards)

	cardNums := []byte{'1'}

	cardNum := 1

	for i, card := range cards[1:] {
		if card != cards[i] {
			cardNum++
		}

		cardNums = append(cardNums, '0'+byte(cardNum))
	}

	// pattern match the normalized hand against the hand types.
	handTypes := map[string]string{
		"11111": "five-of-a-kind",
		"11112": "four-of-a-kind",
		"12222": "four-of-a-kind",
		"11122": "full-house",
		"11222": "full-house",
		"11123": "three-of-a-kind",
		"12223": "three-of-a-kind",
		"12333": "three-of-a-kind",
		"11223": "two-pair",
		"11233": "two-pair",
		"12233": "two-pair",
		"11234": "one-pair",
		"12234": "one-pair",
		"12334": "one-pair",
		"12344": "one-pair",
		"12345": "high-card",
	}

	handType, ok := handTypes[string(cardNums)]
	if !ok {
		panic(fmt.Sprintf("unknown hand type: %s", string(cardNums)))
	}

	return handType
}

func (h Hand) jokerHandType() string {
	// sort the cards, and assign them numbers.  Jokers get '*'s
	cards := make([]byte, len(h.cards))
	copy(cards, h.cards)
	slices.Sort(cards)

	firstCard, cardNum := '1', 1
	if cards[0] == 'J' {
		firstCard, cardNum = '*', 0
	}

	cardNums := []byte{byte(firstCard)}

	for i, card := range cards[1:] {
		if card == 'J' {
			cardNums = append(cardNums, '*')
		} else {
			if card != cards[i] {
				cardNum++
			}

			cardNums = append(cardNums, '0'+byte(cardNum))
		}
	}

	// sort the numbered cards _again_ so jokers are at the start.
	slices.Sort(cardNums)

	// pattern match the normalized hand against the hand types.
	handTypes := map[string]string{
		"11111": "five-of-a-kind",
		"*1111": "five-of-a-kind",
		"**111": "five-of-a-kind",
		"***11": "five-of-a-kind",
		"****1": "five-of-a-kind",
		"*****": "five-of-a-kind",

		"11112": "four-of-a-kind",
		"12222": "four-of-a-kind",
		"*1112": "four-of-a-kind",
		"*1222": "four-of-a-kind",
		"**112": "four-of-a-kind",
		"**122": "four-of-a-kind",
		"***12": "four-of-a-kind",

		"11122": "full-house",
		"11222": "full-house",
		"*1122": "full-house",

		"11123": "three-of-a-kind",
		"12223": "three-of-a-kind",
		"12333": "three-of-a-kind",
		"*1123": "three-of-a-kind",
		"*1223": "three-of-a-kind",
		"*1233": "three-of-a-kind",
		"**123": "three-of-a-kind",

		"11223": "two-pair",
		"11233": "two-pair",
		"12233": "two-pair",

		"11234": "one-pair",
		"12234": "one-pair",
		"12334": "one-pair",
		"12344": "one-pair",
		"*1234": "one-pair",

		"12345": "high-card",
	}

	handType, ok := handTypes[string(cardNums)]
	if !ok {
		panic(fmt.Sprintf("unknown hand type: %s", string(cardNums)))
	}

	return handType
}
