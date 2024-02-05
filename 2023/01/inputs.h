#ifndef INPUTS_H
#define INPUTS_H

#include <fstream>
#include <iostream>
#include <string>
#include <vector>

// Input for function, would be read from file if was allot more data
inline std::vector<std::string> read_input_test {"Hello", "C++", "Advent of Code", "2023!"};

// Inputs to test findCalibrationValue NOLINTBEGIN
inline std::vector<std::pair<unsigned int, std::string>> test_find_cal_vals= {
    {12, "1abc2"},
    {38, "pqr3stu8vwx"},
    {15, "a1b2c3d4e5f"},
    {77, "treb7uchet"}
    }; // NOLINTEND

// Read inputs from text file, use exception handling to prevent throws
inline auto readInput(std::vector<std::string>& output) noexcept -> bool
{
    try
    {
        std::ifstream file("input.txt");

        if( file.good() && file.is_open())
        {
            while(!file.eof())
            {
                std::string line;
                getline(file, line);
                output.push_back(line);
            }
        }
    }
    catch(const std::exception& e)
    {
        std::cerr << e.what() << '\n';
        return false;
    }
    return true;
}

#endif