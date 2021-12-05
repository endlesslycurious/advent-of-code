package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Parse int from string, handling errors
func ParseInt(in string) int {
	val, err := strconv.Atoi(in)

	if err != nil {
		log.Fatalln(err)
	}

	return val
}

// Split string with seperator, panic if result is longer than count
func SplitString(in string, sep string, count int) []string {
	result := strings.Split(in, sep)

	if len(result) != count {
		panic("unexpected number of sub-strings")
	}

	return result
}

// Parse a Point struct from string
func ParsePoint(in string) Point {
	ints := SplitString(in, ",", 2)

	x := ParseInt(ints[0])
	y := ParseInt(ints[1])

	return Point{x, y}
}

func LoadLines(filename string) []Line {
	data := make([]Line, 0)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		pointsStr := SplitString(line, " -> ", 2)

		start := ParsePoint(pointsStr[0])
		end := ParsePoint(pointsStr[1])

		data = append(data, Line{start, end})
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Loaded ", len(data), " lines from ", filename)

	return data
}

func main() {
	input := LoadLines("./2021/05/input.txt")

	answer := Part1(input)
	fmt.Println("Part 1 Answer: ", answer)
}

type Point struct {
	x int
	y int
}

type Line struct {
	start Point
	end   Point
}

func Part1(input []Line) int {
	return 0
}
