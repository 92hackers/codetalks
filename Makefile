.PHONY: build test install

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
