package day14

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/jhungerford/adventofcode-2023/internal/util"
	"slices"
)

func Part1(platform *Platform) string {
	platform.Tilt(platform.North())

	return fmt.Sprintf("%d", platform.Load())
}

func Part2(platform *Platform) string {
	platform.NumCycles(1000000000)

	return fmt.Sprintf("%d", platform.Load())
}

func LoadPlatform(filename string) (*Platform, error) {
	lines, err := util.ReadInputLines(filename)
	if err != nil {
		return nil, err
	}

	rows := make([][]byte, 0, len(lines))

	for _, line := range lines {
		rows = append(rows, []byte(line))
	}

	return &Platform{rows}, nil
}

type Position struct {
	row, col int
}

type Direction struct {
	// move moves a rock in the direction
	move func(Position) Position
	// traverse iterates over the platform in the opposite direction so rocks move out of the way of each other.
	traverse func(func(Position, byte) bool)
}

type Platform struct {
	rows [][]byte
}

func (p *Platform) Tilt(dir Direction) {
	for pos, value := range dir.traverse {
		if value == 'O' {
			fromPos := pos
			toPos := dir.move(fromPos)

			for p.inBounds(toPos) && p.get(toPos) == '.' {
				p.set(toPos, 'O')
				p.set(fromPos, '.')

				fromPos = toPos
				toPos = dir.move(fromPos)
			}
		}
	}
}

func (p *Platform) Cycle() {
	p.Tilt(p.North())
	p.Tilt(p.West())
	p.Tilt(p.South())
	p.Tilt(p.East())
}

func (p *Platform) NumCycles(num int) {
	var hashes []string

	hashes = append(hashes, p.hash())

	cycles := 0
	cycleStart := -1

	for cycles < num {
		p.Cycle()
		cycles++

		cycleHash := p.hash()

		cycleStart = slices.Index(hashes, cycleHash)
		if cycleStart != -1 {
			break
		}

		hashes = append(hashes, cycleHash)
	}

	cycleLen := cycles - cycleStart
	remainingCycles := (num - cycles) % cycleLen

	for i := 0; i < remainingCycles; i++ {
		p.Cycle()
	}
}

func (p *Platform) North() Direction {
	return Direction{
		move: func(p Position) Position {
			return Position{
				row: p.row - 1,
				col: p.col,
			}
		},
		traverse: func(yield func(Position, byte) bool) {
			// North is top-to-bottom
			for rowNum, row := range p.rows {
				for colNum := range row {
					pos := Position{row: rowNum, col: colNum}
					if !yield(pos, p.get(pos)) {
						return
					}
				}
			}
		},
	}
}

func (p *Platform) South() Direction {
	return Direction{
		move: func(p Position) Position {
			return Position{
				row: p.row + 1,
				col: p.col,
			}
		},
		traverse: func(yield func(Position, byte) bool) {
			// South is bottom-to-top
			for rowNum := len(p.rows) - 1; rowNum >= 0; rowNum-- {
				for colNum := range p.rows[rowNum] {
					pos := Position{row: rowNum, col: colNum}
					if !yield(pos, p.get(pos)) {
						return
					}
				}
			}
		},
	}
}

func (p *Platform) East() Direction {
	return Direction{
		move: func(p Position) Position {
			return Position{
				row: p.row,
				col: p.col + 1,
			}
		},
		traverse: func(yield func(Position, byte) bool) {
			// East is right-to-left
			for colNum := len(p.rows[0]) - 1; colNum >= 0; colNum-- {
				for rowNum := range p.rows {
					pos := Position{row: rowNum, col: colNum}
					if !yield(pos, p.get(pos)) {
						return
					}
				}
			}
		},
	}
}

func (p *Platform) West() Direction {
	return Direction{
		move: func(p Position) Position {
			return Position{
				row: p.row,
				col: p.col - 1,
			}
		},
		traverse: func(yield func(Position, byte) bool) {
			// West is left-to-right
			for colNum := 0; colNum < len(p.rows[0]); colNum++ {
				for rowNum := range p.rows {
					pos := Position{row: rowNum, col: colNum}
					if !yield(pos, p.get(pos)) {
						return
					}
				}
			}
		},
	}
}

func (p *Platform) inBounds(pos Position) bool {
	return pos.row >= 0 &&
		pos.col >= 0 &&
		pos.row < len(p.rows) &&
		pos.col < len(p.rows[0])
}

func (p *Platform) get(pos Position) byte {
	return p.rows[pos.row][pos.col]
}

func (p *Platform) set(pos Position, value byte) {
	p.rows[pos.row][pos.col] = value
}

func (p *Platform) Load() int {
	load := 0

	for rowNum, row := range p.rows {
		for _, col := range row {
			if col == 'O' {
				load += len(p.rows) - rowNum
			}
		}
	}

	return load
}

func (p *Platform) String() string {
	s := ""

	for _, row := range p.rows {
		s += fmt.Sprintf("%s\n", string(row))
	}

	return s
}

func (p *Platform) hash() string {
	hash := sha256.New()

	for _, row := range p.rows {
		hash.Write(row)
	}

	return base64.StdEncoding.EncodeToString(hash.Sum(nil))
}
