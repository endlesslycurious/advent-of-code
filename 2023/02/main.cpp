#include <cassert>
#include <cstring>
#include <iostream>
#include <string>
#include <sstream>
#include <vector>

#include "inputs.h"

auto readGame(const std::string& input) -> Game
{
    std::istringstream stream(input);
    std::string token;
    Game game;
    Round round = {};
    int word_count = 0;
    int last_number = 0;

    while( std::getline(stream, token, ' '))
    {
        if( word_count == 1)
        {
            game.id = atoi(token.c_str());
        }else if (word_count > 1) {
            if ( (word_count & 1) == 1)
            {
                // odd = color
                size_t token_length = token.size();
                char last_ch = token[token_length - 1];
                if( last_ch == ',' || last_ch ==';')
                {
                    --token_length;
                }

                if( strncmp("blue", token.c_str(), token_length) == 0)
                {
                    assert(round.blue == 0);
                    round.blue = last_number;
                }
                else if( strncmp("green", token.c_str(), token_length) == 0)
                {
                    assert(round.green == 0);
                    round.green = last_number;
                }
                else if( strncmp("red", token.c_str(), token_length) == 0)
                {
                    assert(round.red == 0);
                    round.red = last_number;
                }

                if( last_ch != ',')
                {
                    game.rounds.push_back(round);
                    round = {};
                }
            }else{
                // even = number
                last_number = atoi(token.c_str());
            }
        }

        word_count++;
    }

    return game;
}

// Print words to console and return word count
auto process(const std::vector<std::string>& games, const Round& limits) -> unsigned
{
    unsigned sum = 0;
    for (const std::string& input : games)
    {
        Game game = readGame(input);
        Round total = game.Max();

        if((total.red<=limits.red) && (total.green<=limits.green) && (total.blue<=limits.blue))
        {
            sum += game.id;
        }
    }

    return sum;
}

auto main() -> int
{
    try
    {
        // load inputs from text file
        std::vector<std::string> inputs;
        auto success = readInput(inputs);
        assert(success);
        std::cout << "Read " << inputs.size() << " inputs from " << INPUT_FILE << "\n";
        assert(inputs.size() == INPUT_COUNT);

        std::cout << "-- Beginning testing! --" << "\n";

        // process the inputs
        Round limits = {14,13,12};
        [[maybe_unused]] auto sum = process(inputs,limits);
        std::cout << "Sum: " << sum << '\n';
        assert(sum == INPUT_ANS);

        std::cout << "-- Testing passed! --" << "\n";
    }
    catch(const std::exception& e)
    {
        std::cerr << e.what() << std::endl;
        __builtin_trap();
        
        return 1;
    }

    return 0;
}
