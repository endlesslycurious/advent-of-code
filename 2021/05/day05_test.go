package main

import (
	"log"
	"testing"
)

var input = []Line{
	{Point{x: 0, y: 9}, Point{x: 5, y: 9}},
	{Point{x: 8, y: 0}, Point{x: 0, y: 8}},
	{Point{x: 9, y: 4}, Point{x: 3, y: 4}},
	{Point{x: 2, y: 2}, Point{x: 2, y: 1}},
	{Point{x: 7, y: 0}, Point{x: 7, y: 4}},
	{Point{x: 6, y: 4}, Point{x: 2, y: 0}},
	{Point{x: 0, y: 9}, Point{x: 2, y: 9}},
	{Point{x: 3, y: 4}, Point{x: 1, y: 4}},
	{Point{x: 0, y: 0}, Point{x: 8, y: 8}},
	{Point{x: 5, y: 5}, Point{x: 8, y: 2}},
}

func TestGridIncrement(T *testing.T) {
	grid := CreateGrid()

	expected := []int{1, 2, 3, 1, 4}
	inputs := []Point{{0, 0}, {0, 0}, {0, 0}, {1, 1}, {0, 0}}

	for i := range expected {
		actual := grid.Increment(inputs[i].x, inputs[i].y)

		if expected[i] != actual {
			log.Fatalln("FAIL! Expected", expected[i], "got", actual)
		}
	}
}

func TestGridIntersections(T *testing.T) {
	grid := CreateGrid()
	inputs := []Point{{0, 0}, {0, 0}, {0, 0}, {1, 1}, {0, 0}}
	expected := 1

	for i := range inputs {
		grid.Increment(inputs[i].x, inputs[i].y)
	}

	actual := grid.Intersections()

	if actual != expected {
		log.Fatalln("FAIL! Expected", expected, "got", actual)
	}
}

func TestPart1(T *testing.T) {
	expected := 5

	actual := Part1(input)

	if expected != actual {
		log.Fatalln("FAIL! Expected", expected, "got", actual)
	}
}

func TestPart2(T *testing.T) {
	expected := 12

	actual := Part2(input)

	if expected != actual {
		log.Fatalln("FAIL! Expected", expected, "got", actual)
	}
}
