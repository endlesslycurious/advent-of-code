package main

import "fmt"

const (
	boardSide    = 5
	boardNumbers = boardSide * boardSide
)

// Represents a 5x5 bingo board
type Board struct {
	Numbers [boardNumbers]int
	Score   int
}

// Update board with latest number and return score (non-zero) if bingo
func (b *Board) Update(num int) int {
	return 0
}

// read bingo numbers and boards from text file
func ReadInputs(filename string) ([]int, []Board) {
	numbers := []int{}
	boards := []Board{}

	return numbers, boards
}

func main() {
	numbers, boards := ReadInputs("./2021/04/input.txt")

	answer := Part1(numbers, boards)
	fmt.Println("Part1 Answer: ", answer)
}

func Part1(numbers []int, boards []Board) int {
	return 0
}
