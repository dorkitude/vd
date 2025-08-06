.PHONY: build dev install clean test help

# Binary name
BINARY_NAME=vd
# Installation directory for dev
USER_BIN_DIR=$(HOME)/bin

# Default target
help:
	@echo "Available targets:"
	@echo "  make dev     - Build and install to ~/bin for development"
	@echo "  make build   - Build the binary"
	@echo "  make clean   - Remove built binary"
	@echo "  make test    - Run tests"

# Build the binary
build:
	@echo "🔨 Building $(BINARY_NAME)..."
	@go build -o $(BINARY_NAME) cmd/vd/main.go
	@echo "✅ Build complete: ./$(BINARY_NAME)"

# Install for development - builds and copies to user's bin
dev: build
	@echo "📦 Installing to $(USER_BIN_DIR)..."
	@mkdir -p $(USER_BIN_DIR)
	@cp $(BINARY_NAME) $(USER_BIN_DIR)/
	@chmod +x $(USER_BIN_DIR)/$(BINARY_NAME)
	@echo "✅ Installed! You can now use 'vd' command"
	@echo ""
	@echo "⚠️  Make sure $(USER_BIN_DIR) is in your PATH:"
	@echo "    Add this to your ~/.zshrc or ~/.bashrc if needed:"
	@echo "    export PATH=\"$(USER_BIN_DIR):$$PATH\""

# Install system-wide (requires sudo)
install: build
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