package aocutils

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

type Grid struct {
	g      map[Vector2D]rune
	Width  int
	Height int

	Formats map[rune]func(string, ...interface{}) string
}

// NewGrid makes a new grid from a multiline grid string.
// e.g.:
// /------\
// | o   x|
// \------/
func NewGrid(s string) Grid {
	var width, height int

	lines := strings.Split(strings.TrimSpace(s), "\n")
	height = len(lines)

	g := make(map[Vector2D]rune)
	for y, line := range lines {
		width = MaxInt(width, len(line))
		for x, char := range line {
			g[Vector2D{x, y}] = char
		}
	}

	return Grid{g, width, height, nil}
}

func (g Grid) Set(pos Vector2D, r rune) {
	g.g[pos] = r
}

// Cell returns a cell with the given coordinates.
func (g Grid) Cell(pos Vector2D) rune {
	r, exists := g.g[pos]
	if !exists {
		return '.'
	}
	return r
}

// CountChars counts the amount of cells that contain a given rune in the grid.
func (g Grid) CountChars(r rune) int {
	var total int
	for _, cell := range g.g {
		if cell == r {
			total++
		}
	}

	return total
}

// Display prints the grid.
func (g Grid) Display() {
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			g.PrintCell(x, y)
		}
		fmt.Println()
	}
}

// PrintCell prints a cell using custom formatting if any is specified.
func (g Grid) PrintCell(x, y int) {
	cell := g.Cell(Vector2D{x, y})
	format := color.WhiteString

	if g.Formats != nil {
		f, exists := g.Formats[cell]
		if exists {
			format = f
		}
	}

	fmt.Print(format("%c", cell))
}
