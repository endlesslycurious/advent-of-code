package main

import "fmt"

func ReadInput(filename string) []Digit {
	return []Digit{}
}

func main() {
	input := ReadInput("./2021/08/input.txt")

	answer := Part1(input)
	fmt.Println("Part 1 Answer:", answer)
}

type Digit struct {
	input  []string
	output []string
}

func Part1(digits []Digit) int {
	return 0
}
