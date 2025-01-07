// Makefile (to automate tasks)

# Build variables
SRC := $(wildcard *.go)
TESTS := $(wildcard tests/*.go)
BINARY := networking_suite

all: build

# Build the application
build:
	go build -o $(BINARY) $(SRC)

# Run tests with coverage
coverage:
	go test -cover ./... > coverage.txt

# Run linters (requires golangci-lint)
lint:
	golangci-lint run ./...

# Run unit and integration tests
test:
	go test ./... -v

clean:
	rm -f $(BINARY) coverage.txt

.PHONY: build test lint clean coverage
