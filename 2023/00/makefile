# First attempt at a makefile, refrences:
# - https://devblogs.microsoft.com/cppblog/now-announcing-makefile-support-in-visual-studio-code/
# - https://makefiletutorial.com/#makefile-cookbook
# - https://blog.djnavarro.net/posts/2023-06-30_makefiles/

# Insert homebrew installed LLVM at front of path, mainly to get bear intercepts to work
export PATH := /opt/homebrew/opt/llvm/bin/:$(PATH)

# Build directory; intermediate files, symbols and binary
TEMP=bin

# Name executable after parent directory e.g AoC-00, AoC-01, etc
PROG=AoC-$(notdir $(CURDIR))
PROG_OUT=$(TEMP)/$(PROG)

# Compiler options
CXX=clang++
CFLAGS=--std=c++20 -Wall -Wextra -Wpedantic -Wshadow -Werror -g

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
lint: main.cpp inputs.h
	$(LINT) $(LINT_FLAGS) main.cpp inputs.h

# Compile and link cpp to an executable
$(PROG_OUT): $(TEMP) main.cpp inputs.h
	$(CXX) $(CFLAGS) main.cpp -I inputs.h -o $(PROG_OUT)

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
