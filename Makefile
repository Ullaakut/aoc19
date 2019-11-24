GO_FILES=$(shell find . -name '*.go' | grep -v 'vendor/')

HEADER := $(shell tput -Txterm setaf 3; tput bold)

default: run

## Format the go source
fmt:
	@echo "${HEADER}=== Running gofmt ==="
	@gofmt -s -w $(GO_FILES)

## Test all solvers
test:
	@echo "${HEADER}=== Running go test ==="
	@go test ./...

## Run all solvers
run:
	@echo "${HEADER}=== Solving all challenges ==="
	@for f in $(shell find * -mindepth 0 -maxdepth 0 -type d -not -empty | grep Day | sort); \
		do \
		echo "\n$${f} Part 1:\n" && (cd $${f}/1 && go build -o solver main.go && ./solver) && \
		echo "\n$${f} Part 2:\n" && (cd $${f}/2 && go build -o solver main.go && ./solver); \
	done
