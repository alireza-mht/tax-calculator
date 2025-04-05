APP_NAME := tc
BUILD_DIR := ../../bin
API_DIR := ../../api

############# Build Application #############################
.PHONY: build 
build: build-api build-go

.PHONY: build-api 
build-api:
	$(MAKE) -C $(API_DIR)

.PHONY: build-go 
build-go:
	@echo "Building $(APP_NAME)..."
	go build -o $(BUILD_DIR)/$(APP_NAME) .

############# Install deps #############################

.PHONY: deps
deps:
	@if ! which oapi-codegen; then \
		go get github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@v2.1.0; \
	fi


############# Test Application #############################

.PHONY: test
test:
	@echo "Running tests for $(APP_NAME)..."


################## Linter ##################################

.PHONY: lint
lint:
	@echo "Linting $(APP_NAME)..."


################## Doc ##################################

.PHONY: doc
doc:
	@echo "Generating docs for $(APP_NAME)..."