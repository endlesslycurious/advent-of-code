package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func ReadPowerBits(filename string) []string {
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

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Loaded ", len(data), " readings from ", filename)

	return data
}

func main() {
	input := ReadPowerBits("./2021/03/input.txt")

	answer := Part1(input)
	fmt.Println("Part 1 answer: ", answer)
}

func Part1(input []string) int {
	oneFreq := make(map[int]int, 5)
	total := len(input)

	// Generate freqency map of 1s in input
	for _, data := range input {
		for i, bit := range data {
			if bit == '1' {
				oneFreq[i]++
			}
		}
	}

	var gamma, epsilon int

	// Use 1s frequency map to work out most common value per column/bit
	// to assemble gamma & epsilon
	for i, count := range oneFreq {
		// five bits but indexed from zero
		val := 1 << (4 - i)

		if count > total/2 {
			gamma = gamma | val
		} else {
			epsilon = epsilon | val
		}
	}

	return gamma * epsilon
}
