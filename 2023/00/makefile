# First attempt at a makefile, refrences:
# - https://devblogs.microsoft.com/cppblog/now-announcing-makefile-support-in-visual-studio-code/
# - https://makefiletutorial.com/#makefile-cookbook
# - https://blog.djnavarro.net/posts/2023-06-30_makefiles/

# Variables
TEMP=bin

# Build all target
.PHONY: all
all: clean dir build run

# Create TEMP directory
.PHONY: dir
dir:
	mkdir -p $(TEMP)

# Compile and links cpp to an executable with debug (-g) information
.PHONY: build
build: dir day00.cpp inputs.h
	clang++ --std=c++20 day00.cpp -I inputs.h -g -o $(TEMP)/day00

# Launch the executable
.PHONY: run
run: build
	$(TEMP)/day00

# Delete TEMP directory
.PHONY: clean
clean:
	rm -rf $(TEMP)