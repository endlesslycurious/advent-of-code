#!/bin/bash
# Make executable with chmod u+x <filename>.sh

# Check parameter count
if [ "$#" -ne 2 ]; then
    echo "Requires two parameters; year then day!"
else
    year=$1
    day=$2
fi

# Create year folder, if they don't exist
if [ ! -d "$year" ]
then
    echo "Year ${year} folder doesn't exist, creating.."
    mkdir ./$year
    echo "Year ${year} folder created!"
else
    echo "Year ${year} folder already exists, skipping."
fi

# Create day folder and files, if they don't exist
if [ ! -d "$year/$day" ]
then
    echo "Day ${day} doesn't exist, creating.."
    mkdir ./$year/$day

    touch "./$year/$day/day${day}.txt"
    touch "./$year/$day/day${day}.go"
    touch "./$year/$day/day${day}_test.go"

    echo "Day ${day} folder & files created!"
else
    echo "Day ${day} folder already exists, skipping."
fi