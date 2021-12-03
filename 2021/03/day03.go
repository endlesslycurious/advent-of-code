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
	return 0
}
