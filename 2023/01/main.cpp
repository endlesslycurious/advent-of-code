#include <cassert>
#include <iostream>
#include <string>
#include <utility>
#include <vector>

#include "inputs.h"

// Make double digit int using first and last digits found in source string
auto findCalibrationValueP1(const std::string& input) -> unsigned int{
    std::string first, last;

     for(const char& chr :input)
     {
        if (chr >= '0' && chr <= '9')
        {
            std::string digit(&chr, 1);

            if( first.empty() && last.empty())
            {
                first = digit;
            }

            last = digit;
        }
     }

     std::string res_str = first + last;
     unsigned int res = stoi(res_str);

    return res;
}

void testFindCalibrationValueP1()
{
    std::cout << " - Testing findCalibrationValueP1! -" << std::endl;

    // test finding calibration values using supplied inputs & outputs
    unsigned int test_sum {0};
    for( auto pair : test_find_cal_vals_p1)
    {
        std::cout << "   " << pair.second << " => " << pair.first << std::endl;
        auto val = findCalibrationValueP1(pair.second);
        assert(pair.first == val);

        test_sum += val;
    }

    std::cout << "   " << "Sum = " << test_sum << std::endl;
    assert(test_sum == P1_SUM);

    std::cout << " - Testing passed! -" << std::endl;
}

void testFindCalibrationValueP2()
{
    std::cout << " - Testing findCalibrationValueP2! -" << std::endl;

    // test finding calibration values using supplied inputs & outputs
    unsigned int test_sum {0};
    for( auto pair : test_find_cal_vals_p2)
    {
        std::cout << "   " << pair.second << " => " << pair.first << std::endl;
        auto val = findCalibrationValueP1(pair.second);
        assert(pair.first == val);

        test_sum += val;
    }

    std::cout << "   " << "Sum = " << test_sum << std::endl;
    assert(test_sum == P2_SUM);

    std::cout << " - Testing passed! -" << std::endl;
}

auto main() -> int
{
    try
    {
        testFindCalibrationValueP1();
        testFindCalibrationValueP2();

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