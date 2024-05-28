.PHONY: build test install get-version publish help

repo_root=$(shell git rev-parse --show-toplevel)
version=$(shell cat $(repo_root)/version.txt)

all: build

help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  build     Build codetalks"
	@echo "  test      Run tests"
	@echo "  install   Install codetalks"
	@echo "  publish   Publish codetalks to go remote packages index"
	@echo "  help      Show this help message"

get-version:
	@echo "Version..."
	@bash ./scripts/get-version.sh

update-version: get-version
	@echo "Updating version..."
	@bash ./scripts/update-go-version.sh

build:
	@echo "Building..."
	@mkdir -p bin
	@go build -o bin/ ./cmd/...

test:
	@echo "Testing..."
	@go test -v ./...

install: update-version
	@echo "Installing..."
	@go install ./cmd/...

publish: update-version
	@echo "Publishing to go packages..."
	GOPROXY=proxy.golang.org go list -m github.com/92hackers/codetalks@$(version)
