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
		values := make([]int, 0, len(line))

		for _, str := range strings.Split(line, "") {
			val, err := strconv.Atoi(str)

			if err != nil {
				log.Fatalln(err)
			}

			values = append(values, val)
		}

		data = append(data, values)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	return data
}

func main() {
	filename := "./2021/09/day09_input.txt"
	input := ReadInput(filename)
	fmt.Println("Loaded", len(input), "lines from", filename)

	answer := Part1(input)
	fmt.Println("Answer Part 1", answer)
}

func Part1([][]int) int {
	return 0
}
