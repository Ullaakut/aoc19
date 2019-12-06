package main

import (
	"fmt"
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

	result := solve(string(content), 1)

	disgo.Infoln(style.Success(style.SymbolCheck, " Diagnostics string:"), aocutils.Atoi(result))

	os.Exit(0)
}

const (
	ModePosition  = 0
	ModeImmediate = 1
)

func computeOpcode(opcodes []int, input int) string {
	var result string
	var idx int
	for {
		var jump int
		switch opcodes[idx] % 100 {
		case 1: // adds x1 and x2 and writes to x3
			arg1, arg2, resultPos := parseOpcode(opcodes, idx)
			opcodes[resultPos] = arg1 + arg2
			jump = 4
		case 2: // multiplies x1 and x2 and writes to x3
			arg1, arg2, resultPos := parseOpcode(opcodes, idx)
			opcodes[resultPos] = arg1 * arg2
			jump = 4
		case 3: // writes input to x1
			resultPos := opcodes[idx+1]
			opcodes[resultPos] = input
			jump = 2
		case 4: // output x1
			switch opcodes[idx] / 100 {
			case ModePosition:
				result = result + fmt.Sprint(opcodes[opcodes[idx+1]])
			case ModeImmediate:
				result = result + fmt.Sprint(opcodes[idx+1])
			default:
				panic("invalid output result mode")
			}
			jump = 2
		case 99:
			return result
		default:
			panic("Unknown opcode")
		}

		idx = idx + jump
	}
}

func parseOpcode(opcodes []int, idx int) (arg1 int, arg2 int, resultPos int) {
	var (
		arg1Mode   int
		arg2Mode   int
		resultMode int
	)

	arg1Mode = (opcodes[idx] % 1000) / 100
	switch arg1Mode {
	case ModePosition:
		arg1 = opcodes[opcodes[idx+1]]
	case ModeImmediate:
		arg1 = opcodes[idx+1]
	default:
		panic(fmt.Sprintln("invalid arg1 mode", arg1Mode))
	}

	arg2Mode = (opcodes[idx] % 10000) / 1000
	switch arg2Mode {
	case ModePosition:
		arg2 = opcodes[opcodes[idx+2]]
	case ModeImmediate:
		arg2 = opcodes[idx+2]
	default:
		panic(fmt.Sprintln("invalid arg2 mode", arg2Mode))
	}

	resultMode = opcodes[idx] / 10000
	switch resultMode {
	case ModePosition:
		resultPos = opcodes[idx+3]
	case ModeImmediate:
		panic("invalid mode: write result can't be in immediate mode")
	default:
		panic(fmt.Sprintln("invalid result mode", resultMode))
	}

	return arg1, arg2, resultPos
}

func solve(content string, input int) string {
	opcodesStr := strings.Split(strings.TrimSpace(content), ",")
	var opcodes []int
	for _, opcode := range opcodesStr {
		opcodes = append(opcodes, aocutils.Atoi(opcode))
	}

	return computeOpcode(opcodes, input)
}
