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

	answer = Part2(input)
	fmt.Println("Part 2 Answer:", answer)
}

// scores of incorect closing brackets
var scores = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

// Map closing to opening brackets
var matching = map[rune]rune{
	')': '(', // 41 : 40
	']': '[', // 93 : 91
	'}': '{', // 125 : 121
	'>': '<', // 62 : 60
}

// check line returns true and problem bracket if error otherwise false, 0
func CheckLine(line string) (bool, rune) {
	open := make([]rune, 0, len(line))

	for _, chr := range line {
		if chr == ')' || chr == ']' || chr == '}' || chr == '>' {
			prev := open[len(open)-1]

			if prev != matching[chr] {
				return true, chr
			}

			open = open[:len(open)-1]
		} else {
			open = append(open, chr)
		}

	}

	return false, 0
}

// Generate map of errors by bracket to occurence count
func FindErrors(input []string) map[rune]int {
	errors := make(map[rune]int)

	for _, line := range input {
		error, chr := CheckLine(line)

		if error {
			errors[chr]++
		}
	}

	return errors
}

func Part1(input []string) int {
	var score int
	errors := FindErrors(input)

	for bracket, count := range errors {
		subscore := count * scores[bracket]
		score += subscore
	}

	return score
}

func Part2(input []string) int {
	return 0
}
