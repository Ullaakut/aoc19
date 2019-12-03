package main

import (
	"fmt"

	"github.com/Ullaakut/aoc19/pkg/aocutils"
	"github.com/fatih/color"
)

type Cell struct {
	wireID int
	r      rune
}

type Grid struct {
	g map[aocutils.Vector2D]Cell

	formats map[rune]func(string, ...interface{}) string
}

func NewGrid(formats map[rune]func(string, ...interface{}) string) Grid {
	return Grid{
		g:       make(map[aocutils.Vector2D]Cell),
		formats: formats,
	}
}

// Cell returns a cell with the given coordinates.
func (g Grid) Cell(pos aocutils.Vector2D) rune {
	cell, exists := g.g[pos]
	if !exists {
		return '.'
	}
	return cell.r
}

// DisplaySquare prints a square of the grid, with padding.
func (g Grid) DisplaySquare(xMax, xMin, yMax, yMin int) {
	negativePadding := 1
	positivePaddidng := 2

	for y := yMin - negativePadding; y < yMax+positivePaddidng; y++ {
		for x := xMin - negativePadding; x < xMax+positivePaddidng; x++ {
			g.PrintCell(x, y)
		}
		fmt.Println()
	}
}

// PrintCell prints a cell using custom formatting if any is specified.
func (g Grid) PrintCell(x, y int) {
	cell := g.Cell(aocutils.NewVector2D(x, y))
	format := color.WhiteString

	if g.formats != nil {
		f, exists := g.formats[cell]
		if exists {
			format = f
		}
	}

	fmt.Print(format("%c", cell))
}

func (g Grid) FindClosest(originRune, otherRune rune) int {
	var (
		originPos aocutils.Vector2D
		otherPos  []aocutils.Vector2D
	)

	for pos, cell := range g.g {
		// This is the origin
		if cell.r == originRune {
			originPos = pos
			continue
		}

		if cell.r == otherRune {
			otherPos = append(otherPos, pos)
		}
	}

	var closest int
	for _, pos := range otherPos {
		if closest == 0 {
			closest = originPos.ManhattanDistance(pos)
			continue
		}

		closest = aocutils.MinInt(closest, originPos.ManhattanDistance(pos))
	}

	return closest
}
