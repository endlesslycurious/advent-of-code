#include <iostream>
#include <vector>
#include <string>

#include "input.h"

using namespace std;

int main()
{
    for (const string& word : inputs)
    {
        cout << word << " ";
    }
    cout << endl;
}