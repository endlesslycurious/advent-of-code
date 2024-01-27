#include <cassert>
#include <iostream>
#include <vector>
#include <string>

#include "inputs.h"

// Print words to console and return word count
unsigned long func(const std::vector<std::string>& words)
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

int main()
{
    std::cout << "-- Beginning testing! --" << std::endl;

    auto count = func(inputs);
    assert(count == inputs.size());

    std::cout << "-- Testing passed! --" << std::endl;
}