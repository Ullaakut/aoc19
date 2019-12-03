package main

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/Ullaakut/aoc19/pkg/aocutils"
	"github.com/Ullaakut/disgo"
	"github.com/Ullaakut/disgo/style"
	"github.com/fatih/color"
)

var format = map[rune]func(string, ...interface{}) string{
	'o': color.RedString,
	'X': color.GreenString,
	'+': color.WhiteString,
	'-': color.WhiteString,
	'|': color.WhiteString,
}

func main() {
	disgo.StartStep("Reading input file")
	content, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		_ = disgo.FailStepf("unable to read file ../input.txt: %w", err)
		os.Exit(1)
	}
	disgo.EndStep()

	result := solve(strings.TrimSpace(string(content)))

	disgo.Infoln(style.Success(style.SymbolCheck, " Closest intersection to central port:"), result)

	os.Exit(0)
}

func solve(content string) int {
	grid := NewGrid(format)

	// Set central port.
	grid.g[aocutils.NewVector2D(0, 0)] = Cell{0, 'o'}

	var xMax, xMin, yMax, yMin int
	for wireID, wirePath := range strings.Split(content, "\n") {
		var circuitPosition aocutils.Vector2D
		for _, pathPart := range strings.Split(wirePath, ",") {
			// Compute which direction to go towards and which
			// character to use to print it in the grid, depending
			// on the direction of the instruction.
			direction, r := computeDirection(rune(pathPart[0]))

			// If we are changing directions, set a '+' at the previous position
			// before moving the circuit further.
			if !circuitPosition.IsUnset() {
				grid.g[circuitPosition] = Cell{wireID, '+'}
			}

			// Get distance to travel in direction
			distance := aocutils.Atoi(pathPart[1:])

			// Iterate on each cell to travel through.
			for i := 1; i < distance; i++ {
				// Get cell position from current circuit position + direction
				// multiplied by the iteration number.
				cellPosition := circuitPosition.Add(direction.Mul(i))
				if cellPosition.IsUnset() {
					continue
				}

				// Set boundaries to be able to render the grid later on.
				xMax, xMin, yMax, yMin = checkLimits(cellPosition, xMax, xMin, yMax, yMin)

				// There's already another circuit going through here!
				cell, exists := grid.g[cellPosition]
				if exists && cell.wireID != wireID {
					grid.g[cellPosition] = Cell{wireID, 'X'}
					continue
				}

				// Set the cell in the grid with the right rune depending on
				// which direction the wire is going.
				grid.g[cellPosition] = Cell{wireID, r}
			}

			// Update the circuit position.
			circuitPosition = circuitPosition.Add(direction.Mul(distance))
		}
	}

	// Display the grid.
	// It is recommended to disable it for large inputs.
	grid.DisplaySquare(xMax, xMin, yMax, yMin)

	// Find the closest intersection to the origin point.
	return grid.FindClosest('o', 'X')
}

func computeDirection(dirRune rune) (aocutils.Vector2D, rune) {
	switch dirRune {
	case 'R':
		return aocutils.NewVector2D(1, 0), '-'
	case 'L':
		return aocutils.NewVector2D(-1, 0), '-'
	case 'U':
		return aocutils.NewVector2D(0, 1), '|'
	case 'D':
		return aocutils.NewVector2D(0, -1), '|'
	default:
		return aocutils.NewVector2D(0, 0), '$'
	}
}

func checkLimits(pos aocutils.Vector2D, xMax, xMin, yMax, yMin int) (int, int, int, int) {
	xMax = aocutils.MaxInt(xMax, pos.X())
	xMin = aocutils.MinInt(xMin, pos.X())
	yMax = aocutils.MaxInt(yMax, pos.Y())
	yMin = aocutils.MinInt(yMin, pos.Y())
	return xMax, xMin, yMax, yMin
}
