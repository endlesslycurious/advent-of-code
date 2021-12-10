package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//  Read the input file and return slice of ints
func ReadIntsFile(filename string) []int {
	data := make([]int, 0)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))

		if err != nil {
			log.Fatalln(err)
		}

		data = append(data, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	return data
}

func main() {
	filename := "./2021/01/day01_input.txt"
	data := ReadIntsFile(filename)
	fmt.Println("Loaded ", len(data), " instructions from ", filename)

	answer := Part1(data)
	fmt.Println("Part One: ", answer)

	answer = Part2(data)
	fmt.Println("Part Two: ", answer)
}

// Part1 solves the first problem
func Part1(input []int) int {
	var ans int

	if len(input) < 2 {
		return 0
	}

	for i, j := 0, 1; j < len(input); i, j = i+1, j+1 {
		if input[i] < input[j] {
			ans++
		}
	}

	return ans
}

// Part2 sovle the second problem
func Part2(input []int) int {
	var ans int

	if len(input) < 3 {
		return 0
	}

	prev := -1
	for i, j, k := 0, 1, 2; k < len(input); i, j, k = i+1, j+1, k+1 {
		current := input[i] + input[j] + input[k]

		if prev > 0 && prev < current {
			ans++
		}

		prev = current
	}

	return ans
}
