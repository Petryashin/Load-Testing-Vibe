.PHONY: setup run

# Default values for load test parameters
URL ?= http://localhost:8080
WORKERS ?= 10
DURATION ?= 30s

# Run the load test
run:
	@echo "Running load test..."
	@go run cmd/main.go -url="$(URL)" -workers=$(WORKERS) -duration=$(DURATION)

# Default target
.DEFAULT_GOAL := help

# Help
help:
	@echo "Available commands:"
	@echo "  make run                                      - Run with default parameters"
	@echo "  make run URL=http://your-url.com             - Run with custom URL"
	@echo "  make run WORKERS=50                          - Run with custom number of workers"
	@echo "  make run DURATION=1m                         - Run with custom duration"
	@echo "  make run URL=http://your-url.com WORKERS=50 DURATION=1m  - Run with all custom parameters"