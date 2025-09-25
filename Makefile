# Makefile for mass-git-cloner

.PHONY: build clean test run install help

# Variables
BINARY_NAME=gclone
BUILD_DIR=bin
CMD_DIR=cmd/git-clone

# Default target
all: build

# Build the application
build:
	@echo "🔨 Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) ./$(CMD_DIR)
	@echo "✅ Build complete: $(BUILD_DIR)/$(BINARY_NAME)"

# Clean build artifacts
clean:
	@echo "🧹 Cleaning..."
	@rm -rf $(BUILD_DIR)
	@go clean
	@echo "✅ Clean complete"

# Run tests
test:
	@echo "🧪 Running tests..."
	@go test -v ./...
	@echo "✅ Tests complete"

# Run the application
run: build
	@echo "🚀 Running $(BINARY_NAME)..."
	@./$(BUILD_DIR)/$(BINARY_NAME)

# Install the application globally
install:
	@echo "📦 Installing $(BINARY_NAME)..."
	@go install ./$(CMD_DIR)
	@echo "✅ Installation complete"

# Format code
fmt:
	@echo "📝 Formatting code..."
	@go fmt ./...
	@echo "✅ Formatting complete"

# Run linter
lint:
	@echo "🔍 Running linter..."
	@golangci-lint run
	@echo "✅ Linting complete"

# Download dependencies
deps:
	@echo "📥 Downloading dependencies..."
	@go mod download
	@go mod tidy
	@echo "✅ Dependencies updated"

# Run the demo
demo: build
	@echo "🎨 Starting Mass Git Cloner demo..."
	@./demo.sh

# Show help
help:
	@echo "Available targets:"
	@echo "  build     - Build the application"
	@echo "  clean     - Clean build artifacts"
	@echo "  test      - Run tests"
	@echo "  run       - Build and run the application"
	@echo "  demo      - Build and run the demo"
	@echo "  install   - Install the application globally"
	@echo "  fmt       - Format code"
	@echo "  lint      - Run linter"
	@echo "  deps      - Download and tidy dependencies"
	@echo "  help      - Show this help message"