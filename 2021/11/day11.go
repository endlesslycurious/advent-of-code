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
		vals := make([]int, 0, len(line))

		for _, chr := range strings.Split(line, "") {
			val, err := strconv.Atoi(chr)
			if err != nil {
				log.Fatalln(err)
			}

			vals = append(vals, val)
		}

		data = append(data, vals)
	}

	return data
}

func main() {
	filename := "./2021/11/day11_input.txt"
	input := ReadInput(filename)
	fmt.Println("Read", len(input), "lines from", filename)

	answer := Part1(input, Steps)
	fmt.Println("Part 1 Answer", answer)
}

const (
	Steps         int = 100 // steps in simulation
	EnergyInitial int = 0   // energy of an octopus after flashing
	EnergyFlash   int = 9   // enery level at which an octopus flashes
)

// Convience struct represents location on grid
type Point struct {
	x, y int
}

type Grid struct {
	data   [][]int
	width  int
	height int
}

// Create a grid, copying the input data
func CreateGrid(input [][]int) Grid {
	width := len(input[0])
	height := len(input)

	data := make([][]int, len(input))

	for idx := range input {
		data[idx] = make([]int, len(input[idx]))
		copy(data[idx], input[idx])
	}

	return Grid{data, width, height}
}

// Get the energet level of octopus in cell x,y
func (g Grid) GetEnergy(x, y int) int {
	if x < 0 || x >= g.width {
		log.Fatalln("Get", x, y, "x out of range 0-", g.width)
	} else if y < 0 || y >= g.height {
		log.Fatalln("Get", x, y, "y out of range 0-", g.height)
	}

	return g.data[y][x]
}

// Set the energet level of octopus in cell x,y
func (g *Grid) SetEnergy(x, y, energy int) {
	if x < 0 || x >= g.width {
		log.Fatalln("Get", x, y, "x out of range 0-", g.width)
	} else if y < 0 || y >= g.height {
		log.Fatalln("Get", x, y, "y out of range 0-", g.height)
	}

	g.data[y][x] = energy
}

// True if octopus should flash this step
func (g Grid) ShouldFlash(x, y int) bool {
	return g.GetEnergy(x, y) > EnergyFlash
}

// Increase the energy level of octopus in cell x,y
func (g *Grid) IncEnergy(x, y int) {
	g.SetEnergy(x, y, g.GetEnergy(x, y)+1)
}

// Reset energy level of octopus in cell x,y to initial value
func (g *Grid) ResetEnergy(x, y int) {
	g.SetEnergy(x, y, EnergyInitial)
}

// Return list of neighbours for a cell
func (g Grid) GetNeighbours(x, y int) []Point {
	neighbours := make([]Point, 0, 9)

	for cx := -1; cx <= 1; cx++ {
		for cy := -1; cy <= 1; cy++ {
			// don't check original point
			if cx == 0 && cy == 0 {
				continue
			}

			candidate := Point{x + cx, y + cy}

			if (candidate.x >= 0 && candidate.x < (g.width)) && (candidate.y >= 0 && candidate.y < (g.width)) {
				neighbours = append(neighbours, candidate)
			}
		}
	}

	return neighbours
}

// conveince method for debugging
func (g Grid) Print() {
	for y := range g.data {
		fmt.Println("    ", g.data[y])
	}
}

// Increment octopus energy then flash and update neighbors if enough energy and haven't flashed already
func (g *Grid) UpdateEnergy(x, y int, flashed *map[Point]bool) {
	current := Point{x, y}

	if _, visited := (*flashed)[current]; !visited {
		g.IncEnergy(x, y)

		if g.ShouldFlash(x, y) {
			g.ResetEnergy(x, y)
			(*flashed)[current] = true

			// bump neighbouring octopusses energy too
			for _, neighbour := range g.GetNeighbours(x, y) {
				g.UpdateEnergy(neighbour.x, neighbour.y, flashed)
			}
		}
	}
}

// Update the grid i.e a single step
func (g *Grid) Update() int {
	flashed := make(map[Point]bool)

	// increment the energy of all octopuses in grid, recording those ready to flash
	for x := 0; x < g.width; x++ {
		for y := 0; y < g.height; y++ {
			g.UpdateEnergy(x, y, &flashed)
		}
	}

	return len(flashed)
}

func Part1(input [][]int, steps int) int {
	grid := CreateGrid(input)
	var total int

	for step := 0; step < steps; step++ {
		total += grid.Update()
	}

	return total
}
