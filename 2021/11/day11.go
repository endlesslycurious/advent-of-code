package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadInput(filename string) [][]int {
	data := make([][]int, 0)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		vals := make([]int, 0, len(line))

		for _, chr := range strings.Split(line, "") {
			val, err := strconv.Atoi(chr)
			if err != nil {
				log.Fatalln(err)
			}

			vals = append(vals, val)
		}

		data = append(data, vals)
	}

	return data
}

func main() {
	filename := "./2021/11/day11_input.txt"
	input := ReadInput(filename)
	fmt.Println("Read", len(input), "lines from", filename)

	answer := Part1(input)
	fmt.Println("Part 1 Answer", answer)
}

func Part1(input [][]int) int {
	return 0
}
