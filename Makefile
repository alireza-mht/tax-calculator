APP := tax-calculator
CMD_DIR := cmd/$(APP)

TARGETS := build test lint doc deps clean

.PHONY: $(TARGETS) help

# Generic rule to forward targets
$(TARGETS):
	@echo "==> Running '$@' for $(APP)"
	$(MAKE) -C $(CMD_DIR) $@

help:
	@echo "Usage: make [target]"
	@echo "Targets:"
	@echo "  build   - Build the Go application"
	@echo "  lint    - Run linter"
	@echo "  api     - Generate api required files"
	@echo "  clean   - Clean the added artifacts"