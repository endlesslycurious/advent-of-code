#include <cassert>
#include <iostream>
#include <vector>
#include <string>

#include "input.h"

using namespace std;

int func(vector<string> inputs)
{
    int count = 0;
    for (const string& word : inputs)
    {
        cout << word << " ";
        count++;
    }
    cout << endl;

    return count;
}

int main()
{
    int count = func(inputs);
    assert(count == inputs.size());
}