#include <cassert>
#include <iostream>
#include <string>
#include <utility>
#include <vector>

#include "inputs.h"

// Make double digit int using first and last digits found in source string
auto findCalibrationValueP1(const std::string& input) -> unsigned int{
    int first{-1};
    int last{-1};

     for(const char& chr :input)
     {
        if (chr >= '0' && chr <= '9')
        {
            // need null terminated string or atoi will read too far!
            std::string temp(&chr, 1);
            int digit = atoi(temp.c_str());

            if( first == -1 && last == -1)
            {
                first = digit;
            }

            last = digit;
        }
     }

     std::string res_str = std::to_string(first) + std::to_string(last);
     unsigned int res = stoi(res_str);

    return res;
}

void testFindCalibrationValueP1()
{
    std::cout << " - Testing findCalibrationValueP1! -" << std::endl;

    // test finding calibration values using supplied inputs & outputs
    unsigned int test_sum {0};
    for( auto pair : test_find_cal_vals)
    {
        std::cout << "   " << pair.second << " => " << pair.first << std::endl;
        auto val = findCalibrationValueP1(pair.second);
        assert(pair.first == val);

        test_sum += val;
    }

    std::cout << "   " << "Sum = " << test_sum << std::endl;
    assert(test_sum == 142);

    std::cout << " - Testing passed! -" << std::endl;
}

auto main() -> int
{
    try
    {
        testFindCalibrationValueP1();

        std::cout << " - Beginning V1 calculation from inputs! -" << std::endl;
        // load inputs from text file
        std::vector<const std::string> inputs;
        auto success = readInput(inputs);
        assert(success);
        std::cout << "   " << inputs.size() << " inputs read" << std::endl;
        assert(inputs.size() == INPUTS_COUNT);

        // process the inputs into values and generate the sum
        unsigned int sum {0};
        for( const std::string& line : inputs)
        {
            sum += findCalibrationValueP1(line);
        }
        std::cout << "   Sum = " << sum << std::endl;
        assert (sum == 55621);

    }
    catch(const std::exception& e)
    {
        std::cerr << e.what() << '\n';
        __builtin_debugtrap();

        return 1;
    }

    return 0;
}