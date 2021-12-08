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

	fmt.Println("Loaded ", len(data), " crabs from ", filename)

	return data
}

func main() {
	input := ReadInput("./2021/07/input.txt")

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

func Part1(crabs []int) int {
	// first work out how many crabs at each location
	locations := MapLocations(crabs)

	// work out costs to move all crabs to each location
	costs := make([]int, len(locations))
	for x := range locations {
		var cost int

		for pos, count := range locations {
			dist := pos - x
			if dist < 0 {
				dist = -dist
			}
			cost += dist * count
		}

		costs[x] = cost
	}

	// find the lowest cost to move all crabs to one location
	var optimal int
	for i, cost := range costs {
		if i == 0 {
			optimal = cost
		} else if cost < optimal {
			optimal = cost
		}
	}

	return optimal
}

func Part2(crabs []int) int {
	locations := MapLocations(crabs)

	// work out costs to move all crabs to each location
	costs := make([]int, len(locations))
	for x := range locations {
		var cost int

		for pos, count := range locations {
			dist := pos - x
			if dist < 0 {
				dist = -dist
			}

			// cost is now sum of 0 to dist
			sum := (dist * (dist + 1)) / 2

			cost += sum * count
		}

		costs[x] = cost
	}

	// find the lowest cost to move all crabs to one location
	var optimal int
	for i, cost := range costs {
		if i == 0 {
			optimal = cost
		} else if cost < optimal {
			optimal = cost
		}
	}

	return optimal
}
