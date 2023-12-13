package day13

import (
	"fmt"
	"github.com/jhungerford/adventofcode-2023/internal/util"
)

// LoadNotes returns notes in the given file
func LoadNotes(filename string) (Notes, error) {
	type parseNotes struct {
		patterns []pattern
		current  pattern
	}

	sectionParsers := map[string]util.SectionLineParser[parseNotes]{
		"pattern": func(line string, pn *parseNotes) (string, error) {
			if line == "" {
				pn.patterns = append(pn.patterns, pn.current)
				pn.current = pattern{}
			} else {
				pn.current = append(pn.current, []byte(line))
			}

			return "pattern", nil
		},
	}

	pn, err := util.ParseInputLinesSections(filename, "pattern", parseNotes{}, false, sectionParsers)
	if err != nil {
		return Notes{}, fmt.Errorf("failed to parse '%s': %w", filename, err)
	}

	if len(pn.current) > 0 {
		pn.patterns = append(pn.patterns, pn.current)
	}

	return Notes{
		patterns: pn.patterns,
	}, nil
}

// Part1 returns the sum of the reflection value of each pattern in the notes.
func Part1(notes Notes) int {
	sum := 0

	for _, p := range notes.patterns {
		sum += p.findReflection().value()
	}

	return sum
}

// Part2 returns the sum of the reflection value of each pattern in the notes, with a single smudge fixed.
func Part2(notes Notes) int {
	sum := 0

	for _, p := range notes.patterns {
		sum += p.findSmudgedReflection().value()
	}

	return sum
}

type Notes struct {
	patterns []pattern
}

type pattern [][]byte

// findReflection returns the reflection in this pattern.
func (p pattern) findReflection() reflection {
	reflect := reflection{}

	for row := 0; row < len(p)-1; row++ {
		if p.isRowReflection(row) {
			reflect.above = row + 1
		}
	}

	for col := 0; col < len(p[0])-1; col++ {
		if p.isColReflection(col) {
			reflect.left = col + 1
		}
	}

	return reflect
}

func (p pattern) isRowReflection(row int) bool {
	check := min(row+1, len(p)-row-1)

	for i := 0; i < check; i++ {
		for col := range p[row] {
			if p[row-i][col] != p[row+i+1][col] {
				return false
			}
		}
	}

	return true
}

func (p pattern) isColReflection(col int) bool {
	check := min(col+1, len(p[0])-col-1)

	for i := 0; i < check; i++ {
		for row := range p {
			if p[row][col-i] != p[row][col+i+1] {
				return false
			}
		}
	}

	return true
}

func (p pattern) findSmudgedReflection() reflection {
	reflect := reflection{}

	for row := 0; row < len(p)-1; row++ {
		if p.isRowSmudgedReflection(row) {
			reflect.above = row + 1
		}
	}

	for col := 0; col < len(p[0])-1; col++ {
		if p.isColSmudgedReflection(col) {
			reflect.left = col + 1
		}
	}

	return reflect
}

func (p pattern) isRowSmudgedReflection(row int) bool {
	foundSmudge := false

	check := min(row+1, len(p)-row-1)

	for i := 0; i < check; i++ {
		for col := range p[row] {
			if p[row-i][col] != p[row+i+1][col] {
				if foundSmudge {
					return false
				}

				foundSmudge = true
			}
		}
	}

	return foundSmudge
}

func (p pattern) isColSmudgedReflection(col int) bool {
	foundSmudge := false

	check := min(col+1, len(p[0])-col-1)

	for i := 0; i < check; i++ {
		for row := range p {
			if p[row][col-i] != p[row][col+i+1] {
				if foundSmudge {
					return false
				}

				foundSmudge = true
			}
		}
	}

	return foundSmudge
}

// reflection captures the rows to the left or above a reflection in a pattern
type reflection struct {
	left, above int
}

// value returns the value of this reflection - the number of columns to the left of each vertical line of reflection,
// and 100 * the number of rows above each horizontal line.
func (r reflection) value() int {
	return r.left + 100*r.above
}
