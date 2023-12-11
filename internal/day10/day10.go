package day10

import (
	"container/heap"
	"fmt"
	"github.com/jhungerford/adventofcode-2023/internal/util"
	"slices"
)

// LoadMap loads a Map from the given file.
func LoadMap(filename string) (Map, error) {
	lines, err := util.ReadInputLines(filename)
	if err != nil {
		return Map{}, fmt.Errorf("failed to load map: %w", err)
	}

	m := Map{}

	for row, line := range lines {
		m.tiles = append(m.tiles, []byte(line))

		for col, tile := range line {
			if tile == 'S' {
				m.start = position{row: row, col: col}
			}
		}
	}

	return m, nil
}

// Part1 returns the furthest distance from the starting position.
func Part1(m Map) int {
	queue := distanceQueue{
		{pos: m.start, dist: 0},
	}
	heap.Init(&queue)

	visited := map[position]interface{}{
		m.start: nil,
	}

	longestDist := 0

	for len(queue) > 0 {
		tileDist := heap.Pop(&queue).(*distance)

		longestDist = max(longestDist, tileDist.dist)

		for _, neighbor := range m.neighbors(tileDist.pos) {
			if _, ok := visited[neighbor]; !ok {
				heap.Push(&queue, &distance{
					pos:  neighbor,
					dist: tileDist.dist + 1,
				})

				visited[neighbor] = nil
			}
		}
	}

	return longestDist
}

type distance struct {
	pos  position
	dist int
}

type distanceQueue []*distance

func (dq *distanceQueue) Len() int {
	return len(*dq)
}

func (dq *distanceQueue) Less(i, j int) bool {
	di, dj := (*dq)[i], (*dq)[j]

	// Pop should return elements from lowest to highest distance
	if di.dist != dj.dist {
		return di.dist > dj.dist
	}

	if di.pos.row != dj.pos.row {
		return di.pos.row < dj.pos.row
	}

	return di.pos.col < dj.pos.col
}

func (dq *distanceQueue) Swap(i, j int) {
	(*dq)[i], (*dq)[j] = (*dq)[j], (*dq)[i]
}

func (dq *distanceQueue) Push(x any) {
	*dq = append(*dq, x.(*distance))
}

func (dq *distanceQueue) Pop() any {
	old := *dq
	item := old[0]
	old[0] = nil
	*dq = old[1:]

	return item
}

type Map struct {
	tiles [][]byte
	start position
}

type position struct {
	row, col int
}

// add returns a new position by adding the other position's row and column to this position.
func (p position) add(other position) position {
	return position{
		row: p.row + other.row,
		col: p.col + other.col,
	}
}

// tile types:
// | is a vertical pipe connecting north and south.
// - is a horizontal pipe connecting east and west.
// L is a 90-degree bend connecting north and east.
// J is a 90-degree bend connecting north and west.
// 7 is a 90-degree bend connecting south and west.
// F is a 90-degree bend connecting south and east.
// . is ground; there is no pipe in this tile.
// S is the starting position of the animal; there is a pipe on this tile,
// but your sketch doesn't show what shape the pipe has.
var neighborOffsets = map[byte][]position{
	'|': {{row: -1, col: 0}, {row: 1, col: 0}},
	'-': {{row: 0, col: -1}, {row: 0, col: 1}},
	'L': {{row: -1, col: 0}, {row: 0, col: 1}},
	'J': {{row: -1, col: 0}, {row: 0, col: -1}},
	'7': {{row: 1, col: 0}, {row: 0, col: -1}},
	'F': {{row: 1, col: 0}, {row: 0, col: 1}},
	'.': {},
}

// neighbors returns a list of positions that this tile's pipe is connected to.
func (m Map) neighbors(pos position) []position {
	tile := m.tiles[pos.row][pos.col]

	if tile == 'S' {
		tile = m.startTile()
	}

	offsets := neighborOffsets[tile]
	neighbors := make([]position, 0, len(offsets))

	for _, offset := range offsets {
		neighbor := pos.add(offset)

		if m.inBounds(neighbor) {
			neighbors = append(neighbors, neighbor)
		}
	}

	return neighbors
}

// inBounds returns whether the given position is inside of this map.
func (m Map) inBounds(pos position) bool {
	return pos.row >= 0 && pos.col >= 0 && pos.row < len(m.tiles) && pos.col < len(m.tiles[0])
}

// startTile returns the tile type under the starting position.
func (m Map) startTile() byte {
TileType:
	for startTile, startOffsets := range neighborOffsets {
		if len(startOffsets) == 0 {
			continue TileType
		}

		// Try the tile for the starting position - if the tiles it's connected to point back at the start position,
		// the type is correct.
		for _, startOffset := range startOffsets {
			neighbor := m.start.add(startOffset)
			if !m.inBounds(neighbor) || !slices.Contains(m.neighbors(m.start.add(startOffset)), m.start) {
				continue TileType
			}
		}

		return startTile
	}

	return '.'
}
