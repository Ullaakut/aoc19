package main

import (
	"fmt"
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

type Wire struct {
	ID        int
	Positions []aocutils.Vector2D
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
	grid.g[aocutils.NewVector2D(0, 0)] = Cell{-1, 'o'}

	var xMax, xMin, yMax, yMin int
	var wires []Wire
	for wireID, wirePath := range strings.Split(content, "\n") {
		var positions []aocutils.Vector2D
		var circuitPosition aocutils.Vector2D
		for _, pathPart := range strings.Split(wirePath, ",") {
			// Compute which direction to go towards and which
			// character to use to print it in the grid, depending
			// on the direction of the instruction.
			direction, r := computeDirection(rune(pathPart[0]))

			// If we are changing directions, set a + at the previous position
			// before moving the circuit further.
			if !circuitPosition.IsUnset() {
				grid.g[circuitPosition] = Cell{wireID, '+'}
			}

			// Get distance to travel in direction
			distance := aocutils.Atoi(pathPart[1:])

			// Iterate on each cell to travel through.
			for i := 0; i < distance; i++ {
				// Get cell position from current circuit position + direction
				// multiplied by the iteration number.
				cellPosition := circuitPosition.Add(direction.Mul(i))
				if cellPosition.IsUnset() {
					continue
				}

				positions = append(positions, cellPosition)

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

		wires = append(wires, Wire{
			ID:        wireID,
			Positions: positions,
		})
	}

	// Display the grid.
	// It is recommended to disable it for large inputs.
	grid.DisplaySquare(xMax, xMin, yMax, yMin)

	// Find the closest intersection to the origin point.
	return findClosestIntersection(wires, grid)
}

// findClosestIntersection looks through the wires, and stores the
// positions where they intersect and their distance relative to
// the beginning of the wire. It then returns the intersection with
// the smallest relative distance.
func findClosestIntersection(wires []Wire, grid Grid) int {
	intersections := make(map[aocutils.Vector2D]int)
	for _, wire := range wires {
		for dist, pos := range wire.Positions {
			if grid.Cell(pos) == 'X' {
				intersections[pos] = intersections[pos] + dist + 1
			}
		}
	}

	var closestIntersection int
	for _, distance := range intersections {
		if closestIntersection == 0 {
			closestIntersection = distance
		}
		fmt.Println("Intersection found with distance", distance)
		closestIntersection = aocutils.MinInt(closestIntersection, distance)
	}

	return closestIntersection
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
