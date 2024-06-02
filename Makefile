.PHONY: build test install publish help release

repo_root=$(shell git rev-parse --show-toplevel)
version=$(shell cat $(repo_root)/version.txt)

# Passed as: make release tag=v0.0.1
tag ?=

all: build

help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  build     Build codetalks"
	@echo "  test      Run tests"
	@echo "  install   Install codetalks"
	@echo "  release   Release new version"
	@echo "  publish   Publish codetalks to go remote packages index"
	@echo "  help      Show this help message"

release:
	@echo "Release new version..."
	@bash ./scripts/release.sh $(tag)

build:
	@echo "Building..."
	@mkdir -p bin
	@go build -o bin/ ./cmd/...

test:
	@echo "Testing..."
	@go test -v ./...

install:
	@echo "Installing..."
	@go install ./cmd/...

publish: release
	@echo "Publishing to go packages..."
	GOPROXY=proxy.golang.org go list -m github.com/92hackers/codetalks@$(version)
