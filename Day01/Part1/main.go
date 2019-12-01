package main

import (
	"io/ioutil"
	"math"
	"os"
	"strconv"
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

	disgo.Infoln(style.Success(style.SymbolCheck, " Module weight requirement:"), result)

	os.Exit(0)
}

func solve(content string) int {
	disgo.StartStep("Computing initial weight requirement")
	defer disgo.EndStep()

	modules := strings.Split(content, "\n")

	var fuelRequirement int
	for _, module := range modules {
		moduleWeight, _ := strconv.ParseFloat(module, 10)
		if moduleWeight == 0 {
			continue
		}

		moduleFuelRequirement := int(math.Floor(moduleWeight/3)) - 2

		fuelRequirement = fuelRequirement + moduleFuelRequirement
		disgo.Infoln("Found module with weight", moduleWeight, "which needs", moduleFuelRequirement, "fuel")
	}

	return fuelRequirement
}
