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

	fmt.Println("Loaded ", len(data), " fish from ", filename)

	return data
}

func main() {
	input := ReadInput("./2021/06/input.txt")

	answer := Part1(input, 80)
	fmt.Println("Answer Part 1:", answer)

	answer = Part2(input, 256)
	fmt.Println("Answer Part 2:", answer)
}

const (
	InitialTimer int = 8 // newly created fish
	ResetTimer   int = 6 // fish that has just spawned
	SpawnTimer   int = 0 // spawning threshold
)

func Part1(fishes []int, days int) int {

	for day := 0; day < days; day++ {
		newFish := make([]int, 0)

		for idx := range fishes {
			if fishes[idx] == SpawnTimer {
				fishes[idx] = ResetTimer
				newFish = append(newFish, InitialTimer)
			} else {
				fishes[idx]--
			}
		}

		fishes = append(fishes, newFish...)
	}

	return len(fishes)
}

func Part2(fishes []int, days int) int {
	return 0
}
