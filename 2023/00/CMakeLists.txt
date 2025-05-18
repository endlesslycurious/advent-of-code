cmake_minimum_required(VERSION 3.16)
project(AoC VERSION 0.1 LANGUAGES CXX)

# Set C++ standard and compile options
set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)
add_compile_options(-Wall -Wextra -Wpedantic -Werror -Wshadow -g)

# Automatically export compile_commands.json in the build dir
set(CMAKE_EXPORT_COMPILE_COMMANDS ON)

# macOS: Capture system SDK path for clang-tidy
if(APPLE)
  execute_process(
    COMMAND xcrun --show-sdk-path
    OUTPUT_VARIABLE MACOS_SDK
    OUTPUT_STRIP_TRAILING_WHITESPACE
  )
endif()

# Enable clang-tidy during the build
# Either tweak checks here or move to a .clang-tidy file
if(APPLE)
  set(CMAKE_CXX_CLANG_TIDY
    clang-tidy;
    -extra-arg=-isysroot;
    -extra-arg=${MACOS_SDK};
    -checks=-*,bugprone-*,readability-*,modernize-*;
    -warnings-as-errors=*;
  )
else()
  set(CMAKE_CXX_CLANG_TIDY
    clang-tidy;
    -checks=-*,bugprone-*,readability-*,modernize-*;
    -warnings-as-errors=*;
  )
endif()

# Name binary based on current folder, e.g. AoC-00
get_filename_component(PWD_NAME ${CMAKE_CURRENT_SOURCE_DIR} NAME)
set(EXEC_NAME "AoC-${PWD_NAME}")

# Gather all .cpp files
file(GLOB SRC CONFIGURE_DEPENDS "*.cpp")

# Create binary
add_executable(${EXEC_NAME} ${SRC})
