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

func TestUpdateExtents(T *testing.T) {
	var topLeft, bottomRight Point

	for _, line := range input {
		UpdateExtents(line.start, &bottomRight, &topLeft)
		UpdateExtents(line.end, &bottomRight, &topLeft)
	}

	if topLeft.x != 0 || topLeft.y != 0 {
		log.Fatalln("FAIL! Expected ", Point{0, 0}, "got", topLeft)
	}
	if bottomRight.x != 9 || bottomRight.y != 9 {
		log.Fatalln("FAIL! Expected ", Point{9, 9}, "got", bottomRight)
	}
}

func TestPart1(T *testing.T) {
	expected := 5
	topLeft := Point{0, 0}
	bottomRight := Point{9, 9}

	actual := Part1(input, topLeft, bottomRight)

	if expected != actual {
		log.Fatalln("FAIL! Expected ", expected, " got ", actual)
	}
}
