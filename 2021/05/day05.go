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

	answer = Part2(input)
	fmt.Println("Part 2 Answer: ", answer)
}

type Point struct {
	x int
	y int
}

type Line struct {
	start Point
	end   Point
}

func (l Line) Vertical() bool {
	return l.start.x == l.end.x
}

func (l Line) Horizontal() bool {
	return l.start.y == l.end.y
}

type Grid struct {
	data map[int]map[int]int
}

func CreateGrid() Grid {
	var grid Grid

	grid.data = make(map[int]map[int]int)

	return grid
}

// increment count of point on the grid
func (g *Grid) Increment(x, y int) int {
	_, exists := g.data[x]

	if !exists {
		g.data[x] = make(map[int]int)
	}

	g.data[x][y]++

	return g.data[x][y]
}

// count number of points on grid with count greater than 1
func (g *Grid) Intersections() int {
	var intersections int

	if g.data == nil {
		return intersections
	}

	for x := range g.data {
		for y := range g.data[x] {
			count := g.data[x][y]

			if count > 1 {
				intersections++
			}
		}
	}

	return intersections
}

func (g *Grid) AddVerticalLine(begin, end, x int) {
	// line direction may not be positive along y!
	if begin > end {
		begin, end = end, begin
	}

	for y := begin; y <= end; y++ {
		g.Increment(x, y)
	}
}

func (g *Grid) AddHorizontalLine(begin, end, y int) {
	// line direction may not be positive along x!
	if begin > end {
		begin, end = end, begin
	}

	for x := begin; x <= end; x++ {
		g.Increment(x, y)
	}
}

func (g *Grid) AddDiagonalLine(begin, end Point) {
	// work out direction of the diagonal line in both axis
	xDelta := 1
	if begin.x > end.x {
		xDelta = -1
	}

	yDelta := 1
	if begin.y > end.y {
		yDelta = -1
	}

	// walk the line, incrementing the counts but break out seperatly
	for x, y := begin.x, begin.y; ; x, y = x+xDelta, y+yDelta {
		g.Increment(x, y)

		if x == end.x && y == end.y {
			break
		}
	}
}

func Part1(input []Line, topLeft, bottomRight Point) int {
	grid := CreateGrid()

	for _, line := range input {
		if line.Vertical() {
			grid.AddVerticalLine(line.start.y, line.end.y, line.start.x)
		} else if line.Horizontal() {
			grid.AddHorizontalLine(line.start.x, line.end.x, line.start.y)
		}
	}

	overlap := grid.Intersections()

	return overlap
}

func Part2(input []Line) int {
	grid := CreateGrid()

	for _, line := range input {
		if line.Vertical() {
			grid.AddVerticalLine(line.start.y, line.end.y, line.start.x)
		} else if line.Horizontal() {
			grid.AddHorizontalLine(line.start.x, line.end.x, line.start.y)
		} else {
			grid.AddDiagonalLine(line.start, line.end)
		}
	}

	overlap := grid.Intersections()

	return overlap
}
