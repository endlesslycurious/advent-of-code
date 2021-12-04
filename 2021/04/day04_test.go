package main

import (
	"log"
	"testing"
)

var numbers = []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}

var boards = []*Board{
	{Numbers: []int{22, 13, 17, 11, 0, 8, 2, 23, 4, 24, 21, 9, 14, 16, 7, 6, 10, 3, 18, 5, 1, 12, 20, 15, 19}},
	{Numbers: []int{3, 15, 0, 2, 22, 9, 18, 13, 17, 5, 19, 8, 7, 25, 23, 20, 11, 10, 24, 4, 14, 21, 16, 12, 6}},
	{Numbers: []int{14, 21, 17, 24, 4, 10, 16, 15, 9, 19, 18, 8, 23, 26, 20, 22, 11, 13, 6, 5, 2, 0, 12, 3, 7}},
}

func TestPart1(T *testing.T) {
	expected := 4512
	actual := Part1(numbers, boards)

	if expected != actual {
		log.Fatalln("FAIL! Expected ", expected, " got ", actual)
	}
}

func TestCheckColumns(T *testing.T) {
	// test third column
	called := []int{17, 23, 14, 3, 20}
	board := Board{Numbers: []int{22, 13, 17, 11, 0, 8, 2, 23, 4, 24, 21, 9, 14, 16, 7, 6, 10, 3, 18, 5, 1, 12, 20, 15, 19}}
	expected := 4460

	var actual int
	for _, num := range called {
		actual = board.Update(num)
	}

	if actual == 0 {
		log.Fatalln("FAIL! No bingo found!")
	}

	if actual != expected {
		log.Fatalln("FAIL! Expected ", expected, " got ", actual)
	}
}

func TestCheckRows(T *testing.T) {
	// test third row
	called := []int{21, 9, 14, 16, 7}
	board := Board{Numbers: []int{22, 13, 17, 11, 0, 8, 2, 23, 4, 24, 21, 9, 14, 16, 7, 6, 10, 3, 18, 5, 1, 12, 20, 15, 19}}
	expected := 1631

	var actual int
	for _, num := range called {
		actual = board.Update(num)
	}

	if actual == 0 {
		log.Fatalln("FAIL! No bingo found!")
	}

	if actual != expected {
		log.Fatalln("FAIL! Expected ", expected, " got ", actual)
	}
}
