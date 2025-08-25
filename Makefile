# Makefile for second-brain-automatization

# Module path for go mod init (override with: make setup MODULE=github.com/you/repo)
MODULE ?= second-brain-automatization

.PHONY: setup start lint tests test-html clean help

# Default target
help:
	@echo "Available commands:"
	@echo "  setup     - Install dependencies and development tools"
	@echo "  start     - Start the application with hot reload using Air"
	@echo "  lint      - Run golint on the codebase"
	@echo "  tests     - Run unit tests"
	@echo "  test-html - Run tests with coverage and open HTML report"
	@echo "  clean     - Clean build artifacts and coverage files"
	@echo "  help      - Show this help message"

# Setup development environment
setup:
	@echo "Installing Go dependencies..."
	go mod tidy
	@echo "Installing Air for hot reload..."
	go install github.com/air-verse/air@latest
	@echo "Installing golint..."
	go install golang.org/x/lint/golint@latest
	@echo "Setup complete!"

# Start application with hot reload
start:
	@echo "Starting application with Air..."
	air

# Run linter
lint:
	@echo "Running golint..."
	golint ./...
	@echo "Running go vet..."
	go vet ./...
	@echo "Running go fmt check..."
	@test -z "$$(gofmt -l .)" || (echo "Code is not formatted. Run 'go fmt ./...' to fix." && exit 1)

# Run unit tests
tests:
	@echo "Running unit tests..."
	go test -v ./...

# Run tests with coverage and open HTML report
test-html:
	@echo "Running tests with coverage..."
	go test -coverprofile=coverage.out ./...
	@echo "Generating HTML coverage report..."
	go tool cover -html=coverage.out -o coverage.html
	@echo "Opening coverage report in browser..."
	@if command -v xdg-open > /dev/null; then \
		xdg-open coverage.html; \
	elif command -v open > /dev/null; then \
		open coverage.html; \
	elif command -v start > /dev/null; then \
		start coverage.html; \
	else \
		echo "Please open coverage.html manually in your browser"; \
	fi

# Clean build artifacts and coverage files
clean:
	@echo "Cleaning up..."
	rm -f coverage.out coverage.html
	go clean
	@echo "Clean complete!"