# First attempt at a makefile, refrences:
# - https://devblogs.microsoft.com/cppblog/now-announcing-makefile-support-in-visual-studio-code/
# - https://makefiletutorial.com/#makefile-cookbook
# - https://blog.djnavarro.net/posts/2023-06-30_makefiles/

# Build directory; intermediate files, symbols and binary
TEMP=bin

# Name executable after parent directory e.g AoC-00, AoC-01, etc
PROG=AoC-$(notdir $(CURDIR))

# Compiler options
CC=/opt/homebrew/opt/llvm/bin/clang++
CFLAGS=--std=c++20 -Wall -Wextra -Wpedantic -Wshadow -Werror

# Compile databse
COMPDB=$(TEMP)/compile_commands.json

# Build all target
.PHONY: all
all: clean dir build run

# Create build directory
$(TEMP):
	mkdir -p $(TEMP)

# Generate compile commands database for clang-tidy
.PHONY: bear
bear: makefile $(TEMP)
	bear --output $(COMPDB) -- make build

# Compile and links cpp to an executable with debug (-g) information
.PHONY: build
build: $(TEMP) main.cpp inputs.h
	$(CC) $(CFLAGS) main.cpp -I inputs.h -g -o $(TEMP)/$(PROG)

# Launch the executable
.PHONY: run
run: build
	$(TEMP)/$(PROG)

# Delete build directory
.PHONY: clean
clean:
	rm -rf $(TEMP)