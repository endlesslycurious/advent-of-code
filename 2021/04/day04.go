package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	boardSide     = 5
	boardNumCount = boardSide * boardSide
)

// Represents a 5x5 bingo board
type Board struct {
	Numbers []int
	Score   int
}

// Update board with latest number and return score (non-zero) if bingo
func (b *Board) Update(num int) int {
	return 0
}

func ParseInt(in string) int {
	num, err := strconv.Atoi(in)

	if err != nil {
		log.Fatalln(err)
	}

	return num
}

// read bingo numbers and boards from text file
func ReadInputs(filename string) ([]int, []Board) {
	numbers := make([]int, 0)
	boards := make([]Board, 0)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	// first line is the bingo numbers
	read := scanner.Scan()
	if !read {
		panic("Problem reading numbers line")
	}

	numLine := scanner.Text()
	for _, numStr := range strings.Split(numLine, ",") {
		num := ParseInt(numStr)
		numbers = append(numbers, num)
	}

	// rest of file is board definitions
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// new line marks start of a board definition, next X lines are the rows
		if len(line) == 0 {
			boardNumbers := make([]int, 0, boardNumCount)

			for i := 0; i < boardSide; i++ {
				read := scanner.Scan()
				if !read {
					panic("Problem reading board line")
				}

				rowStr := scanner.Text()
				for _, numStr := range strings.Split(rowStr, " ") {
					// space padded input numbers
					if len(numStr) > 0 {
						num := ParseInt(numStr)
						boardNumbers = append(boardNumbers, num)
					}
				}
			}

			board := Board{boardNumbers, 0}
			boards = append(boards, board)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Loaded ", len(numbers), " numbers and ", len(boards), " boards from ", filename)

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
