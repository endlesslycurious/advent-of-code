# Advent of Code 2023 in C++ 

Trying to keep things simple, see (day) `00` folder for simplest example:
- Make file for building with C++ 20.
- Assert macro for testing.
- Header file for inputs.

## Makefile Targets
A few (phony) makefile targets to make the workflow easier:
- `clean` to delete the temp files and output.
- `lint` to lint the project using clang-tidy.
- `build` to compile (with debug symbols) & lint the project.
- `run` to run the program.

### Prerequisites
This code is expected to be run on a Apple silicon Mac and is not tested on other platforms:
- [LLVM / Clang](https://clang.llvm.org) installed by [Homebrew](https://formulae.brew.sh/formula/llvm).
- [Clang-tidy](https://clang.llvm.org/extra/clang-tidy/) for linting, installed by Homebrew as part of LLVM.
- [Bear](https://github.com/rizsotto/Bear) to generate the compile database used by clang-tidy.
