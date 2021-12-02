package main

import "fmt"

// Represents directions sub can move
type Direction string

const (
	Forward Direction = "forward"
	Up      Direction = "up"
	Down    Direction = "down"
)

// Represents sub direction with magnitude, a line from input file
type Instruction struct {
	direction Direction
	magnitude int
}

// Load instructions from input file
func LoadInstructions(filename string) []Instruction {
	return []Instruction{}
}

func main() {
	input := LoadInstructions("./2021/02/input.txt")
	answer := Part1(input)
	fmt.Println("Part 1 answer: ", answer)
}

// Solve first problem for day 2
func Part1(instructions []Instruction) int {
	return 0
}
