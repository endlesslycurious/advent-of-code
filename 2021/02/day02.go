package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Represents directions sub can move
type Direction string

const (
	Forward Direction = "forward"
	Up      Direction = "up"
	Down    Direction = "down"
)

type Magnitude int

// Represents sub direction with magnitude, a line from input file
type Instruction struct {
	dir Direction
	mag Magnitude
}

// Load instructions from input file
func LoadInstructions(filename string) []Instruction {
	data := make([]Instruction, 0)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		tokens := strings.Split(line, " ")

		var dir Direction
		switch Direction(tokens[0]) {
		case Forward:
			dir = Forward
		case Up:
			dir = Up
		case Down:
			dir = Down
		default:
			panic("Unexpected direction")
		}

		mag, err := strconv.Atoi(tokens[1])
		if err != nil {
			log.Fatalln(err)
		}

		data = append(data, Instruction{dir, Magnitude(mag)})
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	return data
}

func main() {
	filename := "./2021/02/day02_input.txt"
	input := LoadInstructions(filename)
	fmt.Println("Loaded ", len(input), " instructions from ", filename)

	answer := Part1(input)
	fmt.Println("Part 1 answer: ", answer)

	answer = Part2(input)
	fmt.Println("Part 2 answer: ", answer)
}

// Solve first problem for day 2
func Part1(instructions []Instruction) int {
	var horz, vert int

	for _, ins := range instructions {

		switch ins.dir {
		case Up:
			vert -= int(ins.mag)
		case Down:
			vert += int(ins.mag)
		case Forward:
			horz += int(ins.mag)
		}
	}

	return horz * vert
}

// Solve second problem for day 2
func Part2(Instructions []Instruction) int {
	var aim, depth, horz int

	for _, ins := range Instructions {

		switch ins.dir {
		case Up:
			aim -= int(ins.mag)
		case Down:
			aim += int(ins.mag)
		case Forward:
			horz += int(ins.mag)
			depth += aim * int(ins.mag)
		}
	}

	return depth * horz
}
