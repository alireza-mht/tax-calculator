APP := tax-calculator
CMD_DIR := cmd/$(APP)
MAKECMDGOALS ?= help
TARGETS := build test lint doc

$(TARGETS):
	@echo "==> Running '$@' for $(APP)"
	$(MAKE) -C $(CMD_DIR) $@

help:
	@echo "Usage: make [target]"
	@echo "Targets:"
	@echo "  build   - Build the Go application"
	@echo "  test    - Run tests"
	@echo "  lint    - Run linter"
	@echo "  doc     - Generate documentation"