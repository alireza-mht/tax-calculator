APP_NAME := tc
BUILD_DIR := ../../bin
API_DIR := ../../api
PROJECT_ROOT := ../..

ifndef GOPATH 
GOPATH := $(PROJECT_ROOT)/.gopath
$(shell mkdir -p $(GOPATH))
$(warning nWARNING: GOPATH not defined, using $(GOPATH) instead)
endif

############# Build Application #############################
.PHONY: build 
build: build-api build-prepare build-go

.PHONY: build-api 
build-api: deps
	$(MAKE) -C $(API_DIR)

.PHONY: build-go 
build-go:
	@echo "Building $(APP_NAME)..."
	go build -o $(BUILD_DIR)/$(APP_NAME) .

# Get code dependencies
.PHONY: build-prepare
build-prepare:
	go mod tidy
	go fmt ./...
############# Install deps #############################

.PHONY: deps
deps:
	@if ! which oapi-codegen; then \
		go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@v2.1.0; \
	fi
	@if ! which golangci-lint; then \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOPATH)/bin v1.64.8; \
	fi


############# Test Application #############################

.PHONY: test
test:
	@echo "Running tests for $(APP_NAME)..."


################## Linter ##################################

.PHONY: lint
lint: deps
	@echo "Linting $(APP_NAME)..."
	golangci-lint -j 1 run --timeout 10m -c .golangci.yml || true


################## Doc ##################################

.PHONY: doc
doc:
	@echo "Generating docs for $(APP_NAME)..."

############# Clean #############################

.PHONY: clean
clean: clean-api
	@echo "Cleaning $(APP_NAME)..."
	rm -rf $(BUILD_DIR)/$(APP_NAME)*
	go clean

.PHONY: clean-api 
clean-api:
	$(MAKE) -C $(API_DIR) clean
