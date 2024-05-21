.PHONY: build test

build:
	@echo "Building..."
	@mkdir -p bin
	@go build -o bin/ ./cmd/...

test:
	@echo "Testing..."
	@go test -v ./...
