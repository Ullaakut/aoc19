package main

import (
	"io/ioutil"
	"os"

	"github.com/Ullaakut/disgo"
	"github.com/Ullaakut/disgo/style"
)

func main() {
	if err := solve("input.txt"); err != nil {
		disgo.Errorln(style.Failure(style.SymbolCross, " Unable to solve challenge:"), err)
		os.Exit(1)

	}

	os.Exit(0)
}

func solve(inputPath string) error {
	disgo.StartStep("Reading input file")
	content, err := ioutil.ReadFile(inputPath)
	if err != nil {
		return disgo.FailStepf("unable to read file %q: %w", inputPath, err)
	}
	disgo.EndStep()

	disgo.Infoln(style.Success(style.SymbolCheck, " Result:"), content)

	return nil
}
