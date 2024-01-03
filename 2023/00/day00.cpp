#include <iostream>
#include <vector>
#include <string>

using namespace std;

int main()
{
    vector<string> msg {"Hello", "C++", "Advent of Code", "2023!"};

    for (const string& word : msg)
    {
        cout << word << " ";
    }
    cout << endl;
}