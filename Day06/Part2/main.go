package main

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/Ullaakut/disgo"
	"github.com/Ullaakut/disgo/style"
)

func main() {
	disgo.StartStep("Reading input file")
	content, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		_ = disgo.FailStepf("unable to read file ../input.txt: %w", err)
		os.Exit(1)
	}
	disgo.EndStep()

	result := solve(string(content))

	disgo.Infoln(style.Success(style.SymbolCheck, " Orbital transfers:"), result)

	os.Exit(0)
}

func solve(content string) int {
	disgo.StartStep("Computing minimum orbital transfers")
	defer disgo.EndStep()

	orbits := make(map[string]string)
	for _, orbit := range strings.Split(strings.TrimSpace(content), "\n") {
		bodies := strings.Split(orbit, ")")

		mainObject := bodies[0]
		orbitingObject := bodies[1]
		orbits[orbitingObject] = mainObject
	}

	// Loop through all of the bodies and store distance from begin
	// position for each one of them.
	distances := make(map[string]int)
	orbiting, exists := orbits["YOU"]
	for {
		if exists != true {
			break
		}

		// Store distance.
		distances[orbiting] = len(distances)

		// The value of 'orbiting' will go through
		// all of the bodies that directly/indirectly orbit 'orbit'.
		orbiting, exists = orbits[orbiting]
	}

	// Calculate total distance from YOU to SAN using the previously computed
	// distances.
	var totalDistance int
	orbiting, exists = orbits["SAN"]
	for {
		if exists != true {
			break
		}

		// If we know the distance of the orbiting body
		// we can already calculate the total distance.
		if _, exists := distances[orbiting]; exists {
			return totalDistance + distances[orbiting]
		}

		totalDistance++
		orbiting, exists = orbits[orbiting]
	}

	return -1
}
