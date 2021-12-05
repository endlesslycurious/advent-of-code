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

func TestPart1(T *testing.T) {
	expected := 5
	actual := Part1(input)

	if expected != actual {
		log.Fatalln("FAIL! Expected ", expected, " got ", actual)
	}
}
