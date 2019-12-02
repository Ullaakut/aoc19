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

func computeOpcode(opcodes []int) int {
	var idx, arg1, arg2, resultPos int
	for {
		switch opcodes[idx] {
		case 1:
			arg1 = opcodes[idx+1]
			arg2 = opcodes[idx+2]
			resultPos = opcodes[idx+3]
			opcodes[resultPos] = opcodes[arg1] + opcodes[arg2]
		case 2:
			arg1 = opcodes[idx+1]
			arg2 = opcodes[idx+2]
			resultPos = opcodes[idx+3]
			opcodes[resultPos] = opcodes[arg1] * opcodes[arg2]
		case 99:
			return opcodes[0]
		default:
			disgo.Errorln("UNKNOWN OPCODE:", opcodes[idx], "AT POSITION", idx)
			disgo.Errorln(opcodes)
			return 0
		}

		idx = idx + 4
	}
}

func solve(content string) int {
	disgo.StartStep("Computing opcode")
	defer disgo.EndStep()

	opcodesStr := strings.Split(strings.TrimSpace(content), ",")
	var opcodes []int
	for _, opcode := range opcodesStr {
		opcodes = append(opcodes, aocutils.Atoi(opcode))
	}

	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			memory := make([]int, len(opcodes))
			copy(memory, opcodes)

			memory[1] = noun
			memory[2] = verb
			result := computeOpcode(memory)
			if result == 19690720 {
				disgo.Infoln(memory)
				return 100*noun + verb
			}
		}
	}
	return -1
}
