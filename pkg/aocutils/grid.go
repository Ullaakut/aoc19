package aocutils

import (
	"fmt"
	"strings"

	"github.com/Ullaakut/disgo"
)

type Grid struct {
	g      map[Vector2D]rune
	Width  int
	Height int
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

	return Grid{g, width, height}
}

// Cell returns a cell with the give coordinates.
func (g Grid) Cell(x, y int) rune {
	r, exists := g.g[Vector2D{x, y}]
	if !exists {
		return ' '
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

func (g Grid) Display() {
	for y := 0; y < g.Height; y++ {
		disgo.Infoln()
		for x := 0; x < g.Width; x++ {
			disgo.Info(fmt.Sprintf("%c", g.Cell(x, y)))
		}
	}
}
