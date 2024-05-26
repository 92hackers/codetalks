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

publish:
	@echo "Publishing to go packages..."
	GOPROXY=proxy.golang.org go list -m github.com/92hackers/codetalks@v0.1.0
