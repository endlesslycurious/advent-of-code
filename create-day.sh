#!/bin/bash

# Make executable with chmod u+x <filename>.sh

echo "Enter year e.g 2021"
read year

if [ ! -d "$year" ]
then
    echo "Year doesn't exist. Creating now"
    mkdir ./$year
    echo "Year created"
fi

echo "Enter day e.g 01"
read day

if [ ! -d "$yeear/$day" ]
then
    echo "Day doesn't exist. Creating now.."
    mkdir ./$year/$day

    touch "./$year/$day/day${day}.txt"
    touch "./$year/$day/day${day}.go"
    touch "./$year/$day/day${day}_test.go"

    echo "Day folder & files created!"
fi