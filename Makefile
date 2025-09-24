.PHONY: build clean serve dev wasm-support test help

# Default target
all: dev

# Build the WASM module
build:
	GOOS=js GOARCH=wasm go build -o web/main.wasm go/main.go

# Start development server
serve:
	@echo "Starting server and opening browser..."
	@cd web && python3 -m http.server 8080

# Development mode: build and serve
dev: build serve

# Clean build artifacts
clean:
	rm -f web/main.wasm

# Update dependencies
deps:
	go mod tidy
	go mod download

# Help
help:
	@echo "Available targets:"
	@echo "  build        - Build the WASM module"
	@echo "  serve        - Start development server on :8080 and open browser"
	@echo "  dev          - Build and serve in development mode"
	@echo "  clean        - Remove build artifacts"
	@echo "  deps         - Update and download dependencies"
	@echo "  help         - Show this help message"