.PHONY: all build build-mak build-validator test validate validate-all summary stats touch reset-retry init-prd clean

# Default target
all: build test validate

# Build both Go-native CLI binaries (mak and validate_state)
build: build-mak build-validator

build-mak:
	@echo "Building Go-native MultiAgent-Kit CLI (agent/tools/mak)..."
	go build -ldflags="-s -w" -o agent/tools/mak agent/tools/validator/main.go
	@chmod +x agent/tools/mak
	@echo "Binary built successfully at agent/tools/mak"

build-validator:
	@echo "Building Go-native validator binary (agent/tools/validate_state)..."
	go build -ldflags="-s -w" -o agent/tools/validate_state agent/tools/validator/main.go
	@chmod +x agent/tools/validate_state
	@echo "Binary built successfully at agent/tools/validate_state"

# Run Go unit tests
test:
	@echo "Running Go tests for kit tooling..."
	cd agent/tools/validator && go test -v ./...

# Run fast state validation
validate:
	@./agent/tools/mak validate

# Run full integrity audit (rules, roles, skills, templates, PRD, state)
validate-all:
	@./agent/tools/mak check-all

# Display visual system dashboard
summary:
	@./agent/tools/mak summary

# Display milestone execution metrics & progress bar
stats:
	@./agent/tools/mak stats

# Sync timestamps in handoff.json and STATE.md
touch:
	@./agent/tools/mak touch

# Safely reset QA retry counter to 0 and unblock
reset-retry:
	@./agent/tools/mak reset-retry

# Initialize default PRD.md template
init-prd:
	@./agent/tools/mak init-prd

# Clean built binaries
clean:
	@rm -f agent/tools/mak agent/tools/validate_state
