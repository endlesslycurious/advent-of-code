#ifndef INPUTS_H
#define INPUTS_H

#include <fstream>
#include <iostream>
#include <string>
#include <vector>

// Input for function, would be read from file if was allot more data
inline std::vector<std::string> read_input_test {"Hello", "C++", "Advent of Code", "2023!"};

// Inputs to test findCalibrationValueP1 NOLINTBEGIN
inline std::vector<std::pair<unsigned int, std::string>> test_find_cal_vals_p1= {
    {12, "1abc2"},
    {38, "pqr3stu8vwx"},
    {15, "a1b2c3d4e5f"},
    {77, "treb7uchet"}
    }; // NOLINTEND

#define P1_SUM 142

// Inputs to test findCalibrationValueP2 NOLINTBEGIN
inline std::vector<std::pair<unsigned int, std::string>> test_find_cal_vals_p2= {
    {29, "two1nine"},
    {83, "eightwothree"},
    {13, "abcone2threexyz"},
    {24, "xtwone3four"},
    {42, "4nineeightseven2"},
    {14, "zoneight234"},
    {76, "7pqrstsixteen"}
    }; // NOLINTEND

#define P2_SUM 281

// digit-words NOLINTBEGIN
inline std::vector<std::pair<const char *, int>> digit_words = {
    {"zero",0},
    {"one",1},
    {"two",2},
    {"three",3},
    {"four",4},
    {"five",5},
    {"six",6},
    {"seven",7},
    {"eight",8},
    {"nine",9}
    };  // NOLINTEND

// Lines of input from inputs.txt
#define INPUTS_COUNT 1000
#define INPUTS_FILE "inputs.txt"

// Read inputs from text file, use exception handling to prevent throws
inline auto readInput(std::vector<std::string>& output) noexcept -> bool
{
    try
    {
        std::ifstream file(INPUTS_FILE);

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
