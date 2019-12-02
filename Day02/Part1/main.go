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

	disgo.Infoln(style.Success(style.SymbolCheck, " Position zero value:"), result)

	os.Exit(0)
}

func solve(content string) int {
	disgo.StartStep("Computing opcode")
	defer disgo.EndStep()

	opcodesStr := strings.Split(strings.TrimSpace(content), ",")
	var opcodes []int
	for _, opcode := range opcodesStr {
		opcodes = append(opcodes, aocutils.Atoi(opcode))
	}

	var idx, arg1, arg2, resultPos int
	for {
		switch opcodes[idx] {
		case 1:
			arg1 = opcodes[idx+1]
			arg2 = opcodes[idx+2]
			resultPos = opcodes[idx+3]
			disgo.Infof("Addition at index %d: %d + %d stored at index %d", idx, opcodes[arg1], opcodes[arg2], resultPos)
			opcodes[resultPos] = opcodes[arg1] + opcodes[arg2]
		case 2:
			arg1 = opcodes[idx+1]
			arg2 = opcodes[idx+2]
			resultPos = opcodes[idx+3]
			disgo.Infof("Multiplication at index %d: %d + %d stored at index %d", idx, opcodes[arg1], opcodes[arg2], resultPos)
			opcodes[resultPos] = opcodes[arg1] * opcodes[arg2]
		case 99:
			disgo.Infoln(opcodes)
			return opcodes[0]
		default:
			disgo.Errorln("UNKNOWN OPCODE:", opcodes[idx], "AT POSITION", idx)
			return 0
		}

		idx = idx + 4
	}
}
