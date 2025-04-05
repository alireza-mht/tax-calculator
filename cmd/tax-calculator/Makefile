APP_NAME := tc
BUILD_DIR := ../../bin

.PHONY: build test lint doc

build:
	@echo "Building $(APP_NAME)..."
	go build -o $(BUILD_DIR)/$(APP_NAME) .

test:
	@echo "Running tests for $(APP_NAME)..."
	go test ./...

lint:
	@echo "Linting $(APP_NAME)..."
	golangci-lint run ./...

doc:
	@echo "Generating docs for $(APP_NAME)..."
	go doc ./...