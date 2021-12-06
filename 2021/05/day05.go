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

// Update topLeft, bottomRight extents from point
func UpdateExtents(point Point, bottomRight, topLeft *Point) {
	if point.x < topLeft.x {
		topLeft.x = point.x
	} else if point.x > bottomRight.x {
		bottomRight.x = point.x
	}

	if point.y < topLeft.y {
		topLeft.y = point.y
	} else if point.y > bottomRight.y {
		bottomRight.y = point.y
	}
}

func LoadLines(filename string) ([]Line, Point, Point) {
	data := make([]Line, 0)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()

	var topLeft, bottomRight Point

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		pointsStr := SplitString(line, " -> ", 2)

		start := ParsePoint(pointsStr[0])
		end := ParsePoint(pointsStr[1])

		UpdateExtents(start, &bottomRight, &topLeft)
		UpdateExtents(end, &bottomRight, &topLeft)

		data = append(data, Line{start, end})
	}

	fmt.Println(topLeft, bottomRight)

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Loaded", len(data), "lines from", filename, "with extents", topLeft, "to", bottomRight)

	return data, topLeft, bottomRight
}

func main() {
	input, topLeft, bottomRight := LoadLines("./2021/05/input.txt")

	answer := Part1(input, topLeft, bottomRight)
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

func (l Line) Horizontal() bool {
	return l.start.x == l.end.x
}

func (l Line) Vertical() bool {
	return l.start.y == l.end.y
}

func Part1(input []Line, topLeft, bottomRight Point) int {
	var overlap int

	for _, line := range input {
		if line.Horizontal() || line.Vertical() {
			// process line
		}
	}

	return overlap
}
