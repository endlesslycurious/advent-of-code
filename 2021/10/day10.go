package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func ReadInput(filename string) []string {
	data := make([]string, 0)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		data = append(data, line)
	}

	return data
}

func main() {
	filename := "./2021/10/day10_input.txt"
	input := ReadInput(filename)
	fmt.Println("Loaded", len(input), "lines from", filename)

	answer := Part1(input)
	fmt.Println("Part 1 Answer:", answer)
}

func Part1(input []string) int {
	var score int
	return score
}
