package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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

			boards = append(boards, Board{Numbers: boardNumbers})
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	return numbers, boards
}

const (
	boardSide     = 5
	boardNumCount = boardSide * boardSide
	calledMarker  = -1
)

// Represents a 5x5 bingo board
type Board struct {
	Numbers []int
	Marked  []int
}

func (b *Board) CheckRows() bool {
	for i := 0; i < boardNumCount; i += boardSide {
		var markers int

		for j := 0; j < boardSide; j++ {
			idx := i + j

			if b.Numbers[idx] == calledMarker {
				markers++
			}
		}

		if markers == boardSide {
			return true
		}
	}
	return false
}

func (b *Board) CheckColumns() bool {
	for i := 0; i < boardSide; i++ {
		var markers int

		for j := 0; j < boardNumCount; j += boardSide {
			idx := i + j

			if b.Numbers[idx] == calledMarker {
				markers++
			}
		}

		if markers == boardSide {
			return true
		}
	}
	return false
}

// Update board with latest number and return score (non-zero) if bingo
func (b *Board) Update(calledNum int) int {
	// mark the called number if it exists in the board
	for i, num := range b.Numbers {
		if num == calledNum {
			b.Numbers[i] = calledMarker
			b.Marked = append(b.Marked, calledNum)
			break
		}
	}

	// check for potential Bingo if marked > boardSide
	var score int
	if len(b.Marked) >= boardSide {
		if b.CheckRows() || b.CheckColumns() {
			for _, num := range b.Numbers {
				if num != calledMarker {
					score += num
				}
			}

			if calledNum != 0 {
				score *= calledNum
			}
		}
	}

	return score
}

func ParseInt(in string) int {
	num, err := strconv.Atoi(in)

	if err != nil {
		log.Fatalln(err)
	}

	return num
}

func main() {
	filename := "./2021/04/day04_input.txt"
	numbers, boards := ReadInputs(filename)
	fmt.Println("Loaded ", len(numbers), " numbers and ", len(boards), " boards from ", filename)

	answer := Part1(numbers, boards)
	fmt.Println("Part1 Answer: ", answer)

	// reload inputs as Part1 will have modified the boards
	_, boards = ReadInputs(filename)

	answer = Part2(numbers, boards)
	fmt.Println("Part2 Answer: ", answer)
}

func Part1(numbers []int, boards []Board) int {
	for _, num := range numbers {

		// use index to access each board in array to avoid modifying local copies generated by range
		for idx := range boards {
			score := boards[idx].Update(num)

			// winner!
			if score != 0 {
				return score
			}
		}
	}

	return 0
}

func Part2(numbers []int, boards []Board) int {
	finished := make(map[*Board]bool, len(boards))
	var lastScore int

	for _, num := range numbers {
		for idx := range boards {
			board := &boards[idx]

			// only check unfinished boards
			if !finished[board] {
				score := board.Update(num)

				// mark winner as finished and save latest score
				if score != 0 {
					finished[board] = true
					lastScore = score
				}
			}
		}
	}

	return lastScore
}
