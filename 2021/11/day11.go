package main

import "fmt"

const (
	Steps int = 100
)

func ReadInput(filename string) [][]int {
	return [][]int{}
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
