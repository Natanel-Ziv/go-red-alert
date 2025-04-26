SHELL := /bin/bash

# Binary name and command path
BIN        := oref-alerts-go
CMD_DIR    := ./cmd/alerts-proxy

# Default port (can be overridden e.g. `make run PORT=8080`)
PORT       ?= 9001

.PHONY: all deps build run run-local test lint docker-build docker-run clean

all: build

# Download deps & tidy go.mod
deps:
	go mod tidy

# Build the Go binary
build: deps
	go build -o $(BIN) $(CMD_DIR)

# Run the compiled binary
run: build
	@echo "Starting $(BIN) on port $(PORT)..."
	PORT=$(PORT) ./$(BIN)

# Run directly via `go run`, loading .env if present
run-local:
	@echo "Running via go run with .env"
	@if [ -f .env ]; then \
	  set -o allexport; source .env; set +o allexport; \
	fi; \
	go run $(CMD_DIR)

# Run all tests
test:
	go test ./... -v

# Lint with golangci-lint (if you have it installed)
lint:
	golangci-lint run

# Build the Docker image
docker-build:
	docker build -t oref-alerts-go .

# Run the Docker container (requires a .env file)
docker-run:
	docker run --rm -p $(PORT):$(PORT) --env-file .env oref-alerts-go

# Clean up binaries
clean:
	rm -f $(BIN)
