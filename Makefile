# Makefile

# Variables
APP_NAME = allOnlineTools
SRC_DIR = .
BUILD_DIR = $(SRC_DIR)

# Build the application
build:
	@echo "Building the application..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(APP_NAME) $(SRC_DIR)
	@echo "Build complete."

# Run the application
run: build
	@echo "Running the application..."
	@$(BUILD_DIR)/$(APP_NAME)

# Clean the build directory
clean:
	@echo "Cleaning up..."
	@rm -rf $(APP_NAME)
	@echo "Clean complete."

# Help
help:
	@echo "Makefile commands:"
	@echo "  make build  - Build the Go application"
	@echo "  make run    - Run the Go application"
	@echo "  make clean  - Clean the build directory"
	@echo "  make help   - Show this help message"

.PHONY: build run clean help
