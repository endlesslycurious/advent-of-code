#include <cassert>
#include <iostream>
#include <vector>
#include <string>

#include "inputs.h"

// Print words to console and return word count
auto func(const std::vector<std::string>& words) -> unsigned long
{
    unsigned long count = 0;
    for (const std::string& word : words)
    {
        std::cout << word << " ";
        count++;
    }
    std::cout << std::endl;

    return count;
}

auto main() -> int
{
    std::cout << "-- Beginning testing! --" << std::endl;

    auto count = func(inputs);
    assert(count == inputs.size());

    std::cout << "-- Testing passed! --" << std::endl;
}