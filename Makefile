.PHONY: build-validator validate validate-all summary touch reset-retry clean

# Default target
all: build-validator validate

# Build the Go-native state and integrity validator binary
build-validator:
	@echo "Building Go-native validator binary..."
	go build -ldflags="-s -w" -o agent/tools/validate_state agent/tools/validator/main.go
	@chmod +x agent/tools/validate_state
	@echo "Binary built successfully at agent/tools/validate_state"

# Run state validation
validate:
	@./agent/tools/validate_state

# Run full integrity check (rules, skills, templates, state)
validate-all:
	@./agent/tools/validate_state --check-all

# Display visual system dashboard
summary:
	@./agent/tools/validate_state --summary

# Sync timestamps in handoff.json and STATE.md
touch:
	@./agent/tools/validate_state --touch

# Safely reset QA retry counter to 0
reset-retry:
	@./agent/tools/validate_state --reset-retry

# Clean built binary
clean:
	@rm -f agent/tools/validate_state
