#include <cassert>
#include <iostream>
#include <string>
#include <vector>

#include "inputs.h"

// Print words to console and return word count
auto process(const std::vector<std::string>& words) -> unsigned long
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
    try
    {
        // load inputs from text file
        std::vector<std::string> inputs;
        auto success = readInput(inputs);
        assert(success);
        assert(inputs.size() == INPUT_COUNT);
        std::cout << "Read " << inputs.size() << " inputs from " << INPUT_FILE << std::endl;

        std::cout << "-- Beginning testing! --" << std::endl;

        // process the inputs
        auto count = process(inputs);
        assert(count == inputs.size());

        std::cout << "-- Testing passed! --" << std::endl;
    }
    catch(const std::exception& e)
    {
        std::cerr << e.what() << '\n';
        __builtin_debugtrap();
        
        return 1;
    }

    return 0;
}