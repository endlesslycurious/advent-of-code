#ifndef INPUTS_H
#define INPUTS_H

#include <algorithm>
#include <fstream>
#include <iostream>
#include <string>
#include <vector>

// Full input
#define INPUT_FILE "inputs.txt"
#define INPUT_COUNT 100
#define INPUT_ANS 2716

// Testing
// #define INPUT_FILE "test.txt"
// #define INPUT_COUNT 5
// #define INPUT_ANS 8

// One handful of cubes from the bag
struct Round {
    int blue = 0;
    int green = 0;
    int red = 0;
};

// Game is a series of rounds with an identity
struct Game {
    int id = 0;
    std::vector<Round> rounds;

    auto Totals() -> Round
    {
        Round totals;

        for (const auto& round : rounds)
        {
            totals.blue += round.blue;
            totals.green += round.green;
            totals.red += round.red;
        }

        return totals;
    }

    auto Max() -> Round
    {
        Round totals;

        for (const auto& round : rounds)
        {
            totals.blue = std::max(round.blue,totals.blue);
            totals.green = std::max(round.green,totals.green);
            totals.red = std::max(round.red,totals.red);
        }

        return totals;
    }
};

// Read inputs from text file, use exception handling to prevent throws
inline auto readInput(std::vector<std::string>& output) noexcept -> bool
{
    try
    {
        std::ifstream file(INPUT_FILE);

        if( file.good() && file.is_open())
        {
            while(!file.eof())
            {
                std::string line;
                getline(file, line);
                if (!line.empty())
                {
                    output.push_back(line);
                }
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
