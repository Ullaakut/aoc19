package main

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/Ullaakut/aoc19/pkg/aocutils"
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

	disgo.Infoln(style.Success(style.SymbolCheck, " Total fuel required:"), result)

	os.Exit(0)
}

func computeFuelRequirement(mass int) int {
	return mass/3 - 2
}

func solve(content string) int {
	disgo.StartStep("Computing fuel requirement")
	defer disgo.EndStep()

	modules := strings.Split(strings.TrimSpace(content), "\n")
	var fuelRequirement int
	for _, module := range modules {
		moduleWeight := aocutils.Atoi(module)

		disgo.Infoln("Found module with weight", moduleWeight)
		moduleFuelRequirement := computeFuelRequirement(moduleWeight)

		fuelWeightRequirement := moduleFuelRequirement
		totalModuleFuelRequirement := moduleFuelRequirement
		for {
			fuelWeightRequirement = computeFuelRequirement(fuelWeightRequirement)
			if fuelWeightRequirement < 0 {
				break
			}

			totalModuleFuelRequirement = totalModuleFuelRequirement + fuelWeightRequirement
		}

		disgo.Infoln("Total fuel requirement for module is", moduleFuelRequirement, "for weight and", totalModuleFuelRequirement, "for fuel")
		fuelRequirement = fuelRequirement + totalModuleFuelRequirement
	}

	return fuelRequirement
}
