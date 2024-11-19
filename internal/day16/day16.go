package day16

import (
	"fmt"
	"github.com/jhungerford/adventofcode-2023/internal/util"
)

func Part1(m *Map) string {
	return fmt.Sprintf("%d", m.part1Energize())
}

func Part2(m *Map) string {
	return fmt.Sprintf("%d", m.part2Energize())
}

type Map struct {
	rows [][]byte
}

type direction int

const (
	north direction = iota
	south
	east
	west
)

type position struct {
	row, col int
}

func (p position) move(dir direction) position {
	switch dir {
	case north:
		return position{
			row: p.row - 1,
			col: p.col,
		}
	case south:
		return position{
			row: p.row + 1,
			col: p.col,
		}
	case east:
		return position{
			row: p.row,
			col: p.col + 1,
		}
	case west:
		return position{
			row: p.row,
			col: p.col - 1,
		}
	}

	return p
}

type beam struct {
	pos position
	dir direction
}

func (b beam) String() string {
	strDir := ""

	switch b.dir {
	case north:
		strDir = "north"
	case south:
		strDir = "south"
	case east:
		strDir = "east"
	case west:
		strDir = "west"
	}

	return fmt.Sprintf("%d, %d - %s", b.pos.row, b.pos.col, strDir)
}

func (b beam) move(dir direction) beam {
	return beam{
		pos: b.pos.move(dir),
		dir: dir,
	}
}

func LoadMap(filename string) (*Map, error) {
	lines, err := util.ReadInputLines(filename)
	if err != nil {
		return nil, err
	}

	rows := make([][]byte, 0, len(lines))

	for _, line := range lines {
		rows = append(rows, []byte(line))
	}

	return &Map{rows}, nil
}

func (m *Map) part1Energize() int {
	return m.energize(beam{pos: position{0, 0}, dir: east})
}

func (m *Map) part2Energize() int {
	// Corners
	allFrom := []beam{
		// top-left
		{pos: position{0, 0}, dir: east},
		{pos: position{0, 0}, dir: south},
		// bottom-left
		{pos: position{len(m.rows) - 1, 0}, dir: east},
		{pos: position{len(m.rows) - 1, 0}, dir: north},
		// bottom-right
		{pos: position{len(m.rows) - 1, len(m.rows[0]) - 1}, dir: west},
		{pos: position{len(m.rows) - 1, len(m.rows[0]) - 1}, dir: north},
		// top-right
		{pos: position{0, len(m.rows[0]) - 1}, dir: west},
		{pos: position{0, len(m.rows[0]) - 1}, dir: south},
	}

	// Top + Bottom
	for col := 1; col < len(m.rows[0])-1; col++ {
		allFrom = append(allFrom, beam{pos: position{0, col}, dir: south})
		allFrom = append(allFrom, beam{pos: position{len(m.rows) - 1, col}, dir: north})
	}

	// Left + Right
	for row := 1; row < len(m.rows)-1; row++ {
		allFrom = append(allFrom, beam{pos: position{row, 0}, dir: east})
		allFrom = append(allFrom, beam{pos: position{row, len(m.rows[0]) - 1}, dir: west})
	}

	maxEnergy := 0

	for _, from := range allFrom {
		maxEnergy = max(maxEnergy, m.energize(from))
	}

	return maxEnergy
}

func (m *Map) energize(from beam) int {
	positions := map[position]interface{}{}
	beams := map[beam]interface{}{}

	toVisit := []beam{
		from,
	}

	for len(toVisit) > 0 {
		at := toVisit[0]
		toVisit = toVisit[1:]

		// Check for beam loops
		if _, ok := beams[at]; ok {
			continue
		}

		// Mark the position as visited.
		beams[at] = nil
		positions[at.pos] = nil

		for _, next := range m.travel(at) {
			if m.inBounds(next) {
				toVisit = append(toVisit, next)
			}
		}
	}

	return len(positions)
}

func (m *Map) inBounds(b beam) bool {
	return b.pos.row >= 0 &&
		b.pos.col >= 0 &&
		b.pos.row < len(m.rows) &&
		b.pos.col < len(m.rows[0])
}

func (m *Map) travel(b beam) []beam {
	switch m.rows[b.pos.row][b.pos.col] {
	case '.':
		return []beam{
			b.move(b.dir),
		}
	case '|':
		if b.dir == north || b.dir == south {
			return []beam{
				b.move(b.dir),
			}
		} else {
			return []beam{
				b.move(north),
				b.move(south),
			}
		}
	case '-':
		if b.dir == east || b.dir == west {
			return []beam{
				b.move(b.dir),
			}
		} else {
			return []beam{
				b.move(east),
				b.move(west),
			}
		}
	case '\\':
		switch b.dir {
		case north:
			return []beam{
				b.move(west),
			}
		case south:
			return []beam{
				b.move(east),
			}
		case east:
			return []beam{
				b.move(south),
			}
		case west:
			return []beam{
				b.move(north),
			}
		}
	case '/':
		switch b.dir {
		case north:
			return []beam{
				b.move(east),
			}
		case south:
			return []beam{
				b.move(west),
			}
		case east:
			return []beam{
				b.move(north),
			}
		case west:
			return []beam{
				b.move(south),
			}
		}
	}

	return nil
}
