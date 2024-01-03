# First attempt at a makefile, refrences:
# - https://devblogs.microsoft.com/cppblog/now-announcing-makefile-support-in-visual-studio-code/
# - https://makefiletutorial.com/#makefile-cookbook
# - https://blog.djnavarro.net/posts/2023-06-30_makefiles/

# Variables
BIN=bin

# Build all target
all: dir Build link

# Create bin directory
dir:
	mkdir -p $(BIN)

# Compile and links cpp to an executable with debug (-g) information
build: dir day00.cpp
	clang++ --std=c++17 day00.cpp -g -o $(BIN)/day00

# Launch the executable
run: build
	$(BIN)/day00

# Debug the executable
debug: build

# Delete bin directory
clean:
	rm -rf $(BIN)