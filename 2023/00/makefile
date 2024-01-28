# First attempt at a makefile, refrences:
# - https://devblogs.microsoft.com/cppblog/now-announcing-makefile-support-in-visual-studio-code/
# - https://makefiletutorial.com/#makefile-cookbook
# - https://blog.djnavarro.net/posts/2023-06-30_makefiles/

# Insert homebrew installed LLVM at front of path, mainly to get bear intercepts to work
export PATH := /opt/homebrew/opt/llvm/bin/:$(PATH)

# Build directory; intermediate files, symbols and binary
TEMP=bin

# Source and header filesets
SRC=$(wildcard *.cpp)
HDR=$(wildcard *.h)

# Name executable after parent directory e.g AoC-00, AoC-01, etc
PROG=AoC-$(notdir $(CURDIR))
PROG_OUT=$(TEMP)/$(PROG)

# Compiler options
# - C++ 20
# - Clang standard library
# - Warnings as errors
# - Extra warnings
# - Debug symbols
CXX=clang++
CFLAGS=--std=c++20 -stdlib=libc++ -Werror -Wall -Wextra -Wpedantic -Wshadow -g

# Linker options, specify the homebrew install C++ libraries
LDFLAGS=-L/opt/homebrew/opt/llvm/lib/c++ -Wl,-rpath,/opt/homebrew/opt/llvm/lib/c++

# Compile databse required for linting
COMPDB=$(TEMP)/compile_commands.json

# Linter options
LINT=clang-tidy
LINT_FLAGS=-checks=-*,bugprone-*,readability-*,modernize-* -p $(COMPDB)

# Build all target
.PHONY: all
all: clean $(TEMP) $(PROG_OUT)

# Create build directory
$(TEMP):
	mkdir -p $(TEMP)

# Generate compile commands database for clang-tidy using bear
$(COMPDB): makefile $(TEMP)
	bear --output $(COMPDB) -- make build

.PHONY: bear
bear: $(COMPDB)

.PHONY: lint
lint: $(SRC) $(HDR)
	$(LINT) $(LINT_FLAGS) $(SRC) $(HDR)

# Compile and link cpp to an executable
$(PROG_OUT): $(TEMP) $(SRC) $(HDR) makefile
	$(CXX) $(CFLAGS) $(SRC) -I $(HDR) -o $(PROG_OUT) $(LDFLAGS)

.PHONY: build
build: $(PROG_OUT) lint

# Launch the executable
.PHONY: run
run: $(PROG_OUT)
	$(PROG_OUT)

# Delete build directory
.PHONY: clean
clean:
	rm -rf $(TEMP)
