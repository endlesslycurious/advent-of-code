package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadInput(filename string) []int {
	data := make([]int, 0)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	line := strings.TrimSpace(scanner.Text())
	for _, str := range strings.Split(line, ",") {
		fish, err := strconv.Atoi(str)

		if err != nil {
			log.Fatalln(err)
		}

		data = append(data, fish)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	return data
}

func main() {
	filename := "./2021/07/day07_input.txt"
	input := ReadInput(filename)
	fmt.Println("Loaded ", len(input), " crabs from ", filename)

	answer := Part1(input)
	fmt.Println("Part 1 Answer: ", answer)

	answer = Part2(input)
	fmt.Println("Part 2 Answer: ", answer)
}

func FindMax(in []int) int {
	var min, max int
	for _, pos := range in {
		if pos > max {
			max = pos
		} else if pos < min {
			min = pos
		}
	}

	if min != 0 {
		panic("expect min to be zero")
	}

	return max
}

func MapLocations(crabs []int) []int {
	max := FindMax(crabs)
	locations := make([]int, max+1)

	for _, pos := range crabs {
		locations[pos]++
	}

	return locations
}

type CostFunc func(int) int

// cost is simply the distance
func SimpleCost(dist int) int {
	return dist
}

// cost is sum of 0 to dist
func SumCost(dist int) int {
	return (dist * (dist + 1)) / 2
}

func FindOptimalCost(crabs []int, costFunc CostFunc) int {
	// first work out how many crabs at each location
	locations := MapLocations(crabs)
	var optimal int

	// work out costs to move all crabs to each location
	for x := range locations {
		var cost int

		for pos, count := range locations {
			dist := pos - x
			if dist < 0 {
				dist = -dist
			}
			cost += costFunc(dist) * count
		}

		// track the optimal cost
		if x == 0 {
			optimal = cost
		} else if cost < optimal {
			optimal = cost
		}
	}

	return optimal
}

func Part1(crabs []int) int {
	return FindOptimalCost(crabs, SimpleCost)
}

func Part2(crabs []int) int {
	return FindOptimalCost(crabs, SumCost)
}
