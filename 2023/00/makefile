# First attempt at a makefile, refrences:
# - https://devblogs.microsoft.com/cppblog/now-announcing-makefile-support-in-visual-studio-code/
# - https://makefiletutorial.com/#makefile-cookbook
# - https://blog.djnavarro.net/posts/2023-06-30_makefiles/

# Variables
BIN=bin

# Build all target
.PHONY: all
all: clean dir build run

# Create bin directory
.PHONY: dir
dir:
	mkdir -p $(BIN)

# Compile and links cpp to an executable with debug (-g) information
.PHONY: build
build: dir day00.cpp inputs.h
	clang++ --std=c++20 day00.cpp -I inputs.h -g -o $(BIN)/day00

# Launch the executable
.PHONY: run
run: build
	$(BIN)/day00

# Delete bin directory
.PHONY: clean
clean:
	rm -rf $(BIN)