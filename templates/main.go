package main

import (
	"io/ioutil"
	"os"

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

	disgo.Infoln(style.Success(style.SymbolCheck, " <RESULT>:"), result)

	os.Exit(0)
}

func solve(content string) int {
	disgo.StartStep("Computing <THING>")
	defer disgo.EndStep()

	var result int

	return result
}
