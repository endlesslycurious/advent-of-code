#include <cassert>
#include <iostream>
#include <string>
#include <utility>
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

// Make double digit int using first and last digits found in source string
auto findCalibrationValue(std::string& input) noexcept -> unsigned int{
    int first{-1};
    int last{-1};

     for(char& chr :input)
     {
        if (chr >= '0' && chr <= '9')
        {
            int digit = atoi(&chr);
            if( first == -1 && last == -1)
            {
                first = last = digit;
            }else
            {
                last = digit;
            }
        }
     }

     std::string res_str = std::to_string(first) + std::to_string(last);
     unsigned int res = stoi(res_str);

    return res;
}

auto main() -> int
{
    std::cout << "-- Beginning testing! --" << std::endl;

    // test finding calibration values using supplied inputs & outputs
    std::cout << " - Testing findCalibrationValue! -" << std::endl;
    for( auto pair : test_find_cal_vals)
    {
        std::cout << "   " << pair.second << " => " << pair.first << std::endl;
        assert(pair.first == findCalibrationValue(pair.second));
    }

    // load inputs from text file
    std::vector<std::string> inputs;
    auto success = readInput(inputs);
    assert(success);
    assert(inputs.size() == read_input_test.size());

    // process the inputs
    auto count = process(inputs);
    assert(count == inputs.size());

    std::cout << "-- Testing passed! --" << std::endl;

    return 0;
}