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

	disgo.Infoln(style.Success(style.SymbolCheck, " Orbits:"), result)

	os.Exit(0)
}

func solve(content string) int {
	disgo.StartStep("Computing direct and indirect orbits")
	defer disgo.EndStep()

	orbits := make(map[string]string)
	for _, orbit := range strings.Split(strings.TrimSpace(content), "\n") {
		bodies := strings.Split(orbit, ")")

		mainObject := bodies[0]
		orbitingObject := bodies[1]
		orbits[orbitingObject] = mainObject
	}

	var total int
	for orbit := range orbits {
		// Loop through all of the bodies that this body orbits, directly and indirectly.
		orbiting, exists := orbits[orbit]
		for {
			if exists != true {
				break
			}

			// The value of 'orbiting' will go through
			// all of the bodies that directly/indirectly orbit 'orbit'.
			orbiting, exists = orbits[orbiting]
			total++
		}
	}

	return total
}
