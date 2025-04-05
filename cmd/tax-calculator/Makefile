APP_NAME := tc
BUILD_DIR := ../../bin
API_DIR := ../../api


.PHONY: build test lint doc


############# Build Application #############################

build: build-api build-go

build-api:
	$(MAKE) -C $(API_DIR)

build-go:
	@echo "Building $(APP_NAME)..."
	go build -o $(BUILD_DIR)/$(APP_NAME) .


############# Test Application #############################

test:
	@echo "Running tests for $(APP_NAME)..."


################## Linter ##################################

lint:
	@echo "Linting $(APP_NAME)..."


################## Doc ##################################

doc:
	@echo "Generating docs for $(APP_NAME)..."