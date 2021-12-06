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

func Part1(input []Line, topLeft, bottomRight Point) int {
	grid := CreateGrid()

	for _, line := range input {
		if line.Vertical() {
			// line direction may not be positive along y!
			if line.start.y <= line.end.y {
				for y := line.start.y; y <= line.end.y; y++ {
					grid.Increment(line.start.x, y)
				}
			} else {
				for y := line.end.y; y <= line.start.y; y++ {
					grid.Increment(line.start.x, y)
				}
			}
		} else if line.Horizontal() {
			// line direction may not be positive along x!
			if line.start.x <= line.end.x {
				for x := line.start.x; x <= line.end.x; x++ {
					grid.Increment(x, line.start.y)
				}
			} else {
				for x := line.end.x; x <= line.start.x; x++ {
					grid.Increment(x, line.start.y)
				}
			}
		}
	}

	overlap := grid.Intersections()

	return overlap
}
