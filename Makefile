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
	go install go.k6.io/xk6/cmd/xk6@latest
	GOARCH=amd64 GOOS=linux \
	XK6_RACE_DETECTOR=1 \
	xk6 build \
		--with github.com/grafana/xk6-client-prometheus-remote@latest \
		--with github.com/mikeblum/xk6-neo4j=.

## clean: Removes any previously created build artifacts.
clean:
	rm -f ./k6

test:
	go test -cover -race ./... -coverprofile=coverage.out

test-integration: build
	./k6 run k6-tests/verify_test.js
	./k6 run k6-tests/conf_test.js
	./k6 run k6-tests/read_test.js

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

.PHONY: build
