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

Row:
	for row := range img.pixels {
		for col := 0; col < len(img.pixels[row]); col++ {
			if img.pixels[row][col] == '#' {
				continue Row
			}
		}

		img.emptyRows = append(img.emptyRows, row)
	}

Col:
	for col := range img.pixels[0] {
		for row := 0; row < len(img.pixels); row++ {
			if img.pixels[row][col] == '#' {
				continue Col
			}
		}

		img.emptyCols = append(img.emptyCols, col)
	}

	for r, row := range img.pixels {
		for c, value := range row {
			if value == '#' {
				img.galaxies = append(img.galaxies, position{row: r, col: c})
			}
		}
	}

	return img, nil
}

// Part1 returns the sum of the shortest path between every pair of galaxies, where each empty row or column is
// expanded one time.
func Part1(img Image) int {
	return img.galaxyDistances(2)
}

// Part2 returns the sum of the shortest path between every pair of galaxies, where empty rows and columns expand
// to 1,000,000.
func Part2(img Image) int {
	return img.galaxyDistances(1_000_000)
}

// galaxyDistances returns the sum of the shortest path between every pair of galaxies, where empty rows and columns
// are replaced by the given number of expansion rows / cols.
func (img Image) galaxyDistances(expansion int) int {
	sum := 0

	for i, galaxy := range img.galaxies {
		for j := i + 1; j < len(img.galaxies); j++ {
			sum += img.distance(galaxy, img.galaxies[j], expansion)
		}
	}

	return sum
}

type Image struct {
	pixels    [][]byte
	galaxies  []position
	emptyRows []int
	emptyCols []int
}

type position struct {
	row, col int
}

// distance returns the shortest path between this position and the other position, where empty rows and columns
// expand by the given amount.
func (img Image) distance(a, b position, expansion int) int {
	emptyRows, emptyCols := 0, 0

	for _, emptyRow := range img.emptyRows {
		if emptyRow > min(a.row, b.row) && emptyRow < max(a.row, b.row) {
			emptyRows++
		}
	}

	for _, emptyCol := range img.emptyCols {
		if emptyCol > min(a.col, b.col) && emptyCol < max(a.col, b.col) {
			emptyCols++
		}
	}

	return max(a.row, b.row) - min(a.row, b.row) +
		max(a.col, b.col) - min(a.col, b.col) +
		emptyRows*(expansion-1) +
		emptyCols*(expansion-1)
}
