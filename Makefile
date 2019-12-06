GO_FILES=$(shell find . -name '*.go' | grep -v 'vendor/')

# Formatting
H=$(shell tput -Txterm setaf 3; tput bold)
B=$(shell tput bold; tput smul)
X=$(shell tput sgr0)

# Removes leading zero from given day
SHORT_DAY := $(shell echo ${DAY} | awk 'sub(/^0+/, "", $$1)')
YEAR ?= 2019

default: help

## Format the go source
fmt:
	@echo "${H}=== Running gofmt ===${X}"
	@gofmt -s -w $(GO_FILES)

## Test all solvers
test:
	@echo "${H}=== Running go test ===${X}"
	@go test -v ./...

## Run all solvers
run:
	@echo "${H}=== Solving all challenges ===${X}"
	@for f in $(shell find * -mindepth 1 -maxdepth 1 -type d -not -empty | grep Day | sort); \
		do echo "\n\t${B}$${f}${X}" && (cd $${f} && go build -o solver main.go && ./solver); \
	done

## Downloads the instructions and inputs for a day (e.g. DAY=02)
download: Day${DAY} Day${DAY}/challenge.md Day${DAY}/input.txt
	@tree Day${DAY}

Day${DAY}:
	@mkdir -p Day${DAY}/Part1 Day${DAY}/Part2
	@cp templates/main.go Day${DAY}/Part1/main.go
	@cp templates/main.go Day${DAY}/Part2/main.go
	@cp templates/main_test.go Day${DAY}/Part1/main_test.go
	@cp templates/main_test.go Day${DAY}/Part2/main_test.go

Day${DAY}/input.txt:
	@echo "${H}=== Downloading input for day ${SHORT_DAY} ===${X}"
	@curl -s -H "cookie: ${AOC_COOKIE}" https://adventofcode.com/${YEAR}/day/${SHORT_DAY}/input > Day${DAY}/input.txt

Day${DAY}/challenge.md: Day${DAY}/challenge.html
	@echo "${H}=== Parsing input ===${X}"
	@./scripts/parse_challenge.sh ${DAY}

## The AOC_COOKIE environment variable should contain a complete session cookie in order to be able to use the make download target
Day${DAY}/challenge.html:
	@[ "${AOC_COOKIE}" ] || ( echo "AOC_COOKIE is not set, please specify your Advent of Code session cookie in order to download challenge and input files"; exit 1 )
	@echo "${H}=== Downloading challenge for day ${SHORT_DAY} ===${X}"
	@curl -s -H "cookie: ${AOC_COOKIE}" https://adventofcode.com/${YEAR}/day/${SHORT_DAY} > Day${DAY}/challenge.html

## Print this message
help:
	@./scripts/help.sh $(abspath $(lastword $(MAKEFILE_LIST)))
