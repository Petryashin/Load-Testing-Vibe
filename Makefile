.PHONY: setup run

# Default values for load test parameters
URL ?= http://localhost:8080
WORKERS ?= 10
DURATION ?= 30s

# Setup project structure
setup:
	@echo "Setting up project structure..."
	@mkdir -p cmd internal/loadtest internal/config pkg/metrics
	@mv main.go cmd/

# Run the load test
run:
	@echo "Running load test..."
	@go run cmd/main.go -url="$(URL)" -workers=$(WORKERS) -duration=$(DURATION)

# Default target
.DEFAULT_GOAL := help

# Help
help:
	@echo "Available commands:"
	@echo "  make setup                                    - Create project structure and move main.go"
	@echo "  make run                                      - Run with default parameters"
	@echo "  make run URL=http://your-url.com             - Run with custom URL"
	@echo "  make run WORKERS=50                          - Run with custom number of workers"
	@echo "  make run DURATION=1m                         - Run with custom duration"
	@echo "  make run URL=http://your-url.com WORKERS=50 DURATION=1m  - Run with all custom parameters"