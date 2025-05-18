# Use makefile to wrap CMake for a smoother workflow!

# Build directory
BUILD_DIR := build
EXEC_NAME := AoC-$(notdir $(CURDIR))

.PHONY: all configure build run clean tidy

# Default target
all: build

# Configure CMake if needed
$(BUILD_DIR)/Makefile:
	cmake -S . -B $(BUILD_DIR)

# Build project (runs configure if needed)
build: $(BUILD_DIR)/Makefile
	cmake --build $(BUILD_DIR)

# Run the binary
run: build
	./$(BUILD_DIR)/$(EXEC_NAME)

# Run clang-tidy manually
lint:
	clang-tidy -p build \
		--extra-arg=-isysroot \
		--extra-arg=$(shell xcrun --show-sdk-path) \
		main.cpp

# Remove build artifacts
clean:
	rm -rf $(BUILD_DIR)
