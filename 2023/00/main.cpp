#include <cassert>
#include <iostream>
#include <vector>
#include <string>

#include "inputs.h"

// Print words to console and return word count
int func(const std::vector<std::string>& inputs)
{
    int count = 0;
    for (const std::string& word : inputs)
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

    int count = func(inputs);
    assert(count == inputs.size());

    std::cout << "-- Testing passed! --" << std::endl;
}