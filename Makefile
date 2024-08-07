.PHONY: build test install publish help release format dev

repo_root=$(shell git rev-parse --show-toplevel)
version=$(shell cat $(repo_root)/version.txt)

# Passed as: make release tag=v0.0.1
tag ?=

all: build

help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  dev       Init development env"
	@echo "  build     Build codetalks"
	@echo "  format    Format Go code"
	@echo "  test      Run tests"
	@echo "  install   Install codetalks"
	@echo "  release   Release new version"
	@echo "  publish   Publish codetalks to go remote packages index"
	@echo "  help      Show this help message"

release:
	@echo "Release new version..."
	@bash ./scripts/release.sh $(tag)

dev:
	@echo "Initializing development environment..."
	@bash ./init-env.sh

format:
	@echo "Formatting..."
	@go fmt ./...

vet:
	@echo "Vetting..."
	@go vet ./...

# @go build -o bin/ ./cmd/...
# -ldflags '-extldflags "-static"' means that the binary will be statically linked
build: vet
	@echo "Building..."
	@mkdir -p bin
	@go build -ldflags '-extldflags "-static"' -o bin/ ./cmd/...

# For a verbose output, use: go test -v ./... instead.
test: vet
	@echo "Testing..."
	@go test ./...

install: vet
	@echo "Installing..."
	@go install ./cmd/...

publish: release
	@echo "Publishing to go packages..."
	GOPROXY=proxy.golang.org go list -m github.com/92hackers/codetalks@$(version)
