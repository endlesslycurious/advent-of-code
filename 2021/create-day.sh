#!/bin/bash
# Make executable with chmod u+x <filename>.sh

# Check parameter count
if [ "$#" -ne 1 ]; then
    echo "Requires day parameter!"
else
    day=$1
fi

# Create day folder and files, if they don't exist
if [ ! -d "$day" ]
then
    echo "Day ${day} doesn't exist, creating.."
    mkdir ./$day

    touch "./$day/day${day}.txt"
    touch "./$day/day${day}.go"
    touch "./$day/day${day}_test.go"
    touch "./$day/day${day}_input.txt"

    echo "Day ${day} folder & files created!"
else
    echo "Day ${day} folder already exists, skipping."
fi