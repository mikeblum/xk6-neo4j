MAKEFLAGS += --silent

all: help

## help: Prints a list of available build targets.
help:
	echo "Usage: make <OPTIONS> ... <TARGETS>"
	echo ""
	echo "Available targets are:"
	echo ''
	sed -n 's/^##//p' ${PWD}/Makefile | column -t -s ':' | sed -e 's/^/ /'
	echo
	echo "Targets run by default are: `sed -n 's/^all: //p' ./Makefile | sed -e 's/ /, /g' | sed -e 's/\(.*\), /\1, and /'`"

build:
	go mod download
	GOARCH=amd64 GOOS=linux xk6 build --with github.com/mikeblum/xk6-neo4j=.

## clean: Removes any previously created build artifacts.
clean:
	rm -f ./k6

test:
	go test -cover -race ./... -coverprofile=coverage.out

## lint: Lint with golangci-lint
lint:
	golangci-lint run ./...

## vendor: Compile upstream dependencies
vendor:
	go mod vendor

## version: Check tooling versions
version:
	go version
	k6 version
	xk6 version

.PHONY: help
