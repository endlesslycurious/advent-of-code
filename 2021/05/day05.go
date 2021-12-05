package main

import "fmt"

func LoadData(filename string) []Line {
	return []Line{}
}

func main() {
	input := LoadData("./2021/04/input.txt")

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
