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
func readfile(filename string) []int {
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
	data := readfile("./input.txt")
	answer := Part1(data)
	fmt.Println(answer)
}

// Part1 solves the first problem
func Part1(input []int) int {
	var ans int

	for _, val := range input {
		fmt.Println(val)
	}

	return ans
}
