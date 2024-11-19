package day16

import (
	"fmt"
	"github.com/jhungerford/adventofcode-2023/internal/util"
)

func Part1(m *Map) string {
	return fmt.Sprintf("%d", m.Energize())
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

func (m *Map) Energize() int {
	positions := map[position]interface{}{}
	beams := map[beam]interface{}{}

	toVisit := []beam{
		{
			pos: position{0, 0},
			dir: east,
		},
	}

	for len(toVisit) > 0 {
		at := toVisit[0]
		toVisit = toVisit[1:]

		// TODO: remove debugging
		fmt.Println(at.String())

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

	for rowNum := 0; rowNum < len(m.rows); rowNum++ {
		for colNum := 0; colNum < len(m.rows[0]); colNum++ {
			if _, ok := positions[position{rowNum, colNum}]; ok {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
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
