.PHONY: build dev install clean test help

# Binary name
BINARY_NAME=vd
# Installation directory for dev
USER_BIN_DIR=$(HOME)/bin

# Default target
help:
	@echo "Available targets:"
	@echo "  make dev     - Install wrapper script that runs latest code via go run"
	@echo "  make build   - Build the binary"
	@echo "  make install - Build and install compiled binary to ~/bin"
	@echo "  make clean   - Remove built binary"
	@echo "  make test    - Run tests"

# Build the binary
build:
	@echo "🔨 Building $(BINARY_NAME)..."
	@go build -o $(BINARY_NAME) cmd/vd/main.go
	@echo "✅ Build complete: ./$(BINARY_NAME)"

# Install development wrapper script that uses go run
dev:
	@echo "🚀 Installing development wrapper to $(USER_BIN_DIR)/vd..."
	@mkdir -p $(USER_BIN_DIR)
	@echo '#!/bin/bash' > $(USER_BIN_DIR)/vd
	@echo 'cd $(shell pwd) && go run cmd/vd/main.go "$$@"' >> $(USER_BIN_DIR)/vd
	@chmod +x $(USER_BIN_DIR)/vd
	@echo "✅ Installed dev wrapper! The 'vd' command will now run the latest code."
	@echo ""
	@echo "⚠️  Make sure $(USER_BIN_DIR) is in your PATH:"
	@echo "    export PATH=\"$(USER_BIN_DIR):$$PATH\""

# Install compiled binary to user's bin
install: build
	@echo "📦 Installing compiled binary to $(USER_BIN_DIR)..."
	@mkdir -p $(USER_BIN_DIR)
	@cp $(BINARY_NAME) $(USER_BIN_DIR)/
	@chmod +x $(USER_BIN_DIR)/$(BINARY_NAME)
	@echo "✅ Installed compiled binary!"

# Install system-wide (requires sudo)
install-system: build
	@echo "📦 Installing to /usr/local/bin (requires sudo)..."
	@sudo cp $(BINARY_NAME) /usr/local/bin/
	@sudo chmod +x /usr/local/bin/$(BINARY_NAME)
	@echo "✅ Installed system-wide!"

# Clean built files
clean:
	@echo "🧹 Cleaning..."
	@rm -f $(BINARY_NAME)
	@echo "✅ Clean complete"

# Run tests
test:
	@echo "🧪 Running tests..."
	@go test ./...
	@echo "✅ Tests complete"