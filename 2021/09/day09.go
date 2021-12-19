package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func ReadInput(filename string) [][]int {
	data := make([][]int, 0)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		values := make([]int, 0, len(line))

		for _, str := range strings.Split(line, "") {
			val, err := strconv.Atoi(str)

			if err != nil {
				log.Fatalln(err)
			}

			values = append(values, val)
		}

		data = append(data, values)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	return data
}

func main() {
	filename := "./2021/09/day09_input.txt"
	input := ReadInput(filename)
	fmt.Println("Loaded", len(input), "lines from", filename)

	answer := Part1(input)
	fmt.Println("Answer Part 1:", answer)

	answer = Part2(input)
	fmt.Println("Answer Part 2:", answer)
}

const (
	BarrierHeight int = 9  // Height of barrier/borders between basins
	Visited       int = -1 // Inidicates already visited cell
)

type Grid struct {
	data   [][]int // points arranged [y][x], indexed from 0
	width  int
	height int
}

// Create grid from data
func CreateGrid(data [][]int) Grid {
	width := len(data[0])
	height := len(data)

	return Grid{data, width, height}
}

// True if height of x,y is not a barrier e.g. 9
func (g Grid) GetNavigable(x, y int) bool {
	return g.GetHeight(x, y) != BarrierHeight
}

// Get height at point x,y
func (g Grid) GetHeight(x, y int) int {
	if x < 0 || x >= g.width {
		log.Fatalln("Get", x, y, "x out of range 0-", g.width)
	} else if y < 0 || y >= g.height {
		log.Fatalln("Get", x, y, "y out of range 0-", g.height)
	}

	return g.data[y][x]
}

// Set height at point x,y to visited
func (g *Grid) SetVisited(x, y int) {
	if x < 0 || x >= g.width {
		log.Fatalln("Set", x, y, "x out of range 0-", g.width)
	} else if y < 0 || y >= g.height {
		log.Fatalln("Set", x, y, "y out of range 0-", g.height)
	}

	g.data[y][x] = Visited
}

// Return true if cell has been visited
func (g Grid) GetVisited(x, y int) bool {
	return g.GetHeight(x, y) == Visited
}

// Get heights of neighbours of point x,y
func (g Grid) GetNeighboursHeights(x, y int) []int {
	if x < 0 || x >= g.width {
		log.Fatalln("Get", x, y, "x out of range 0-", g.width)
	} else if y < 0 || y >= g.height {
		log.Fatalln("Get", x, y, "y out of range 0-", g.height)
	}

	heights := make([]int, 0, 4)

	if x == 0 {
		heights = append(heights, g.GetHeight(x+1, y))
	} else if x == (g.width - 1) {
		heights = append(heights, g.GetHeight(x-1, y))
	} else {
		heights = append(heights, g.GetHeight(x+1, y))
		heights = append(heights, g.GetHeight(x-1, y))
	}

	if y == 0 {
		heights = append(heights, g.GetHeight(x, y+1))
	} else if y == (g.height - 1) {
		heights = append(heights, g.GetHeight(x, y-1))
	} else {
		heights = append(heights, g.GetHeight(x, y+1))
		heights = append(heights, g.GetHeight(x, y-1))
	}

	return heights
}

// Visit neighbours in each direction until grid edge or unnavigable cell encountered
func (g *Grid) VisitNeighbours(x, y, basin int) int {
	size := 0

	if !g.GetVisited(x, y) && g.GetNavigable(x, y) {
		size++

		g.SetVisited(x, y)
	} else {
		return 0
	}

	if x == 0 {
		size += g.VisitNeighbours(x+1, y, basin)
	} else if x == (g.width - 1) {
		size += g.VisitNeighbours(x-1, y, basin)
	} else {
		size += g.VisitNeighbours(x+1, y, basin)
		size += g.VisitNeighbours(x-1, y, basin)
	}

	if y == 0 {
		size += g.VisitNeighbours(x, y+1, basin)
	} else if y == (g.height - 1) {
		size += g.VisitNeighbours(x, y-1, basin)
	} else {
		size += g.VisitNeighbours(x, y+1, basin)
		size += g.VisitNeighbours(x, y-1, basin)
	}

	return size
}

func Part1(input [][]int) int {
	grid := CreateGrid(input)
	var risk int

	for x := 0; x < grid.width; x++ {
		for y := 0; y < grid.height; y++ {
			pointHeight := grid.GetHeight(x, y)
			neighbours := grid.GetNeighboursHeights(x, y)

			lowest := true
			for _, height := range neighbours {
				if height <= pointHeight {
					lowest = false
					break
				}

			}
			if lowest {
				risk += pointHeight + 1
			}
		}
	}

	return risk
}

func Part2(input [][]int) int {
	grid := CreateGrid(input)
	basins := make(map[int]int)
	basin := 0

	// visit nodes to calculate basin sizes
	for x := 0; x < grid.width; x++ {
		for y := 0; y < grid.height; y++ {
			if grid.GetHeight(x, y) != Visited && grid.GetNavigable(x, y) {
				basins[basin] += grid.VisitNeighbours(x, y, basin)

				basin++
			}
		}
	}

	// sort basins sizes
	var result int
	sizes := make([]int, 0, len(basins))
	for _, size := range basins {
		sizes = append(sizes, size)
	}
	sort.Ints(sizes)

	// multiply three largest basin sizes together
	for i := len(sizes) - 3; i < len(sizes); i++ {
		if result > 0 {
			result *= sizes[i]
		} else {
			result = sizes[i]
		}
	}

	return result
}
