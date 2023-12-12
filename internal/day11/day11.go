package day11

import (
	"fmt"
	"github.com/jhungerford/adventofcode-2023/internal/util"
)

// LoadImage loads a galaxy map from the given filename.  '.' is empty space, and '#' is a galaxy.
func LoadImage(filename string) (Image, error) {
	lines, err := util.ReadInputLines(filename)
	if err != nil {
		return Image{}, fmt.Errorf("failed to load image: %w", err)
	}

	img := Image{}

	for _, line := range lines {
		img.pixels = append(img.pixels, []byte(line))
	}

	return img, nil
}

// Part1 returns the sum of the shortest path between every pair of galaxies.
func Part1(img Image) int {
	expanded := img.expand()

	var galaxies []position

	for r, row := range expanded.pixels {
		for c, value := range row {
			if value == '#' {
				galaxies = append(galaxies, position{row: r, col: c})
			}
		}
	}

	sum := 0

	for i, galaxy := range galaxies {
		for j := i + 1; j < len(galaxies); j++ {
			sum += galaxy.distance(galaxies[j])
		}
	}

	return sum
}

type Image struct {
	pixels [][]byte
}

// expand returns a new image where any rows or columns that contain no galaxies are twice as big.
func (img Image) expand() Image {
	var emptyRows []int
	var emptyCols []int

Row:
	for row := range img.pixels {
		for col := 0; col < len(img.pixels[row]); col++ {
			if img.pixels[row][col] == '#' {
				continue Row
			}
		}

		emptyRows = append(emptyRows, row)
	}

Col:
	for col := range img.pixels[0] {
		for row := 0; row < len(img.pixels); row++ {
			if img.pixels[row][col] == '#' {
				continue Col
			}
		}

		emptyCols = append(emptyCols, col)
	}

	expanded := Image{
		pixels: make([][]byte, 0, len(img.pixels)+len(emptyRows)),
	}

	emptyRowsIdx := 0
	for r, row := range img.pixels {
		emptyColsIdx := 0
		expanded.pixels = append(expanded.pixels, make([]byte, 0, len(row)+len(emptyCols)))

		for c, col := range row {
			expanded.pixels[r+emptyRowsIdx] = append(expanded.pixels[r+emptyRowsIdx], col)
			if emptyColsIdx < len(emptyCols) && c == emptyCols[emptyColsIdx] {
				expanded.pixels[r+emptyRowsIdx] = append(expanded.pixels[r+emptyRowsIdx], col)
				emptyColsIdx++
			}
		}

		if emptyRowsIdx < len(emptyRows) && r == emptyRows[emptyRowsIdx] {
			expanded.pixels = append(expanded.pixels, row)
			for i := 0; i < len(emptyCols); i++ {
				expanded.pixels[len(expanded.pixels)-1] = append(expanded.pixels[len(expanded.pixels)-1], '.')
			}

			emptyRowsIdx++
		}
	}

	return expanded
}

type position struct {
	row, col int
}

func (pos position) distance(other position) int {
	return max(pos.row, other.row) - min(pos.row, other.row) + max(pos.col, other.col) - min(pos.col, other.col)
}
