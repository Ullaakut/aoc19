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

	disgo.Infoln(style.Success(style.SymbolCheck, " Module weight requirement:"), result)

	os.Exit(0)
}

func computeFuelRequirement(mass int) int {
	return mass/3 - 2
}

func solve(content string) int {
	disgo.StartStep("Computing initial weight requirement")
	defer disgo.EndStep()

	modules := strings.Split(strings.TrimSpace(content), "\n")

	var fuelRequirement int
	for _, module := range modules {
		moduleWeight := aocutils.Atoi(module)

		moduleFuelRequirement := computeFuelRequirement(moduleWeight)

		fuelRequirement = fuelRequirement + moduleFuelRequirement
		disgo.Infoln("Found module with weight", moduleWeight, "which needs", moduleFuelRequirement, "fuel")
	}

	return fuelRequirement
}
