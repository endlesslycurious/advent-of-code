package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func ReadInput(filename string) []Digit {
	data := make([]Digit, 0)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		digit := DigitFromLine(line)

		data = append(data, digit)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Loaded ", len(data), " digits from ", filename)

	return data
}

func main() {
	input := ReadInput("./2021/08/input.txt")

	answer := Part1(input)
	fmt.Println("Part 1 Answer:", answer)
}

type Digit struct {
	input  []string // input values
	output []string // output values
}

func DigitFromLine(line string) Digit {
	sections := strings.Split(line, "|")

	if len(sections) != 2 {
		panic("Should only be two sections per line")
	}

	input := strings.Split(sections[0], " ")
	output := strings.Split(sections[1], " ")

	return Digit{input, output}
}

func Part1(digits []Digit) int {
	return 0
}
