.PHONY: build test install publish help release format dev

repo_root=$(shell git rev-parse --show-toplevel)
version=$(shell cat $(repo_root)/version.txt)

# For building with cgo, you need to set the following environment variables:
# export CGO_ENABLED=1
# export CGO_LDFLAGS="-L/home/cy/projects/rust-exp/regex/target/release"
# export LD_LIBRARY_PATH="/home/cy/projects/rust-exp/regex/target/release"
#
# Note: We must set above env vars at shell level, not in Makefile.
#

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

# Build options:
#
# @go build -o bin/ ./cmd/...
# 1. -ldflags '-extldflags "-static"' means that the binary will be statically linked
# 2. -s -w flags will strip the debug information from the binary
# 3. -gcflags=-m will print the escape analysis information [Optimize tips]
# 4. `upx -9 bin/codetalks` can be used to compress the binary further more, which can reduce the binary size by 50%.
build: vet
	@echo "Building..."
	@mkdir -p bin
	@go build -gcflags=-m -ldflags '-extldflags "-static" -s -w' -o bin/ ./cmd/...

compress: build
	@upx -9 bin/codetalks

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
