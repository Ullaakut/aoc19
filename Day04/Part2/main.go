package main

import (
	"io/ioutil"
	"os"
	"strconv"
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

	disgo.Infoln(style.Success(style.SymbolCheck, " Amount of potential passwords:"), result)

	os.Exit(0)
}

func solve(content string) int {
	disgo.StartStep("Computing potential passwords")
	defer disgo.EndStep()

	parts := strings.Split(strings.TrimSpace(content), "-")
	begin := aocutils.Atoi(parts[0])
	end := aocutils.Atoi(parts[1])

	var validPasswords int
	for password := begin; password <= end; password++ {
		if isValid(password) {
			validPasswords++
		}
	}

	return validPasswords
}

func isValid(password int) bool {
	pass := strconv.FormatInt(int64(password), 10)

	if len(pass) != 6 {
		return false
	}

	// Verify that numbers are not decreasing.
	if pass[0] > pass[1] ||
		pass[1] > pass[2] ||
		pass[2] > pass[3] ||
		pass[3] > pass[4] ||
		pass[4] > pass[5] {
		return false
	}

	// Check duplicates.
	duplicates := make(map[rune]int, 6)
	for _, r := range pass {
		duplicates[r]++
	}

	// Verify that at least one number is present exactly twice.
	for _, duplicate := range duplicates {
		if duplicate == 2 {
			return true
		}
	}
	return false
}
