# Variables
APP_NAME := starter
BUILD_DIR := ./bin
GO_FILES := $(shell find . -type f -name '*.go')

# Default goal
.DEFAULT_GOAL := help

# Targets
.PHONY: help run build test clean lint fmt vet

## Display help information
help:
	@echo "Usage:"
	@echo "  make <target>"
	@echo ""
	@echo "Targets:"
	@echo "  run        Run the application"
	@echo "  build      Build the binary"
	@echo "  test       Run tests"
	@echo "  lint       Lint the code"
	@echo "  fmt        Format the code"
	@echo "  vet        Check for common issues"
	@echo "  clean      Clean build files"

## Run the application
run:
	# go run main.go
	air


## Build the binary
build: $(GO_FILES)
	@echo "Building $(APP_NAME)..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) main.go
	@echo "Build completed! Binary located at $(BUILD_DIR)/$(APP_NAME)"

## Run tests
test:
	@echo "Running tests..."
	go test ./...

## Lint the code
lint:
	@echo "Linting code..."
	golangci-lint run

## Format the code
fmt:
	@echo "Formatting code..."
	go fmt ./...

## Check for common issues
vet:
	@echo "Running go vet..."
	go vet ./...

## Clean build files
clean:
	@echo "Cleaning build files..."
	@rm -rf $(BUILD_DIR)
	@echo "Clean complete!"

