# First attempt at a makefile, refrences:
# - https://devblogs.microsoft.com/cppblog/now-announcing-makefile-support-in-visual-studio-code/
# - https://makefiletutorial.com/#makefile-cookbook
# - https://blog.djnavarro.net/posts/2023-06-30_makefiles/

# Build directory; intermediate files, symbols and binary
TEMP=bin

# Name executable after parent directory e.g AoC-00, AoC-01, etc
PROG=AoC-$(notdir $(CURDIR))

# Build all target
.PHONY: all
all: clean dir build run

# Create build directory
.PHONY: dir
dir: 
	mkdir -p $(TEMP)

# Compile and links cpp to an executable with debug (-g) information
.PHONY: build
build: dir main.cpp inputs.h
	clang++ --std=c++20 main.cpp -I inputs.h -g -o $(TEMP)/$(PROG)

# Launch the executable
.PHONY: run
run: build
	$(TEMP)/$(PROG)

# Delete build directory
.PHONY: clean
clean:
	rm -rf $(TEMP)