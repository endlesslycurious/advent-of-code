#ifndef INPUTS_H
#define INPUTS_H

#include <fstream>
#include <iostream>
#include <string>
#include <vector>

// Input for function, would be read from file if was allot more data
inline std::vector<std::string> read_input_test {"Hello", "C++", "Advent of Code", "2023!"};

// Read inputs from text file, use exception handling to prevent throws
inline auto readInput(std::vector<std::string>& output) noexcept -> bool
{
    try
    {
        std::ifstream file("inputs.txt");

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