package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

type Grid struct {
	data   [][]int // points arranged [y][x], indexed from 0
	width  int
	height int
}

func CreateGrid(data [][]int) Grid {
	width := len(data[0])
	height := len(data)
	return Grid{data, width, height}
}

// Get hieght at point x,y
func (g Grid) GetHeight(x, y int) int {
	if x < 0 || x >= g.width {
		log.Fatalln("Get", x, y, "x out of range 0-", g.width)
	} else if y < 0 || y >= g.height {
		log.Fatalln("Get", x, y, "y out of range 0-", g.height)
	}

	return g.data[y][x]
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
	return 0
}
