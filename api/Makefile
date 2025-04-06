
SERVER_DIR := ../internal/server
FILES = $(SERVER_DIR)/api/spec.go $(SERVER_DIR)/api/stubs.go

.PHONY: all
all: $(FILES)

$(FILES): openapi.yml
	mkdir -p $(SERVER_DIR)/api
	oapi-codegen --config spec-config.yml openapi.yml
	oapi-codegen --config stubs-config.yml openapi.yml

.PHONY: clean
clean:
	rm -rf ../internal/server/api