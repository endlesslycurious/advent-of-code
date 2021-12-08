package main

import (
	"testing"
)

var numbers = []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}

var boards = []Board{
	{Numbers: []int{22, 13, 17, 11, 0, 8, 2, 23, 4, 24, 21, 9, 14, 16, 7, 6, 10, 3, 18, 5, 1, 12, 20, 15, 19}},
	{Numbers: []int{3, 15, 0, 2, 22, 9, 18, 13, 17, 5, 19, 8, 7, 25, 23, 20, 11, 10, 24, 4, 14, 21, 16, 12, 6}},
	{Numbers: []int{14, 21, 17, 24, 4, 10, 16, 15, 9, 19, 18, 8, 23, 26, 20, 22, 11, 13, 6, 5, 2, 0, 12, 3, 7}},
}

func TestPart1(t *testing.T) {
	expected := 4512
	actual := Part1(numbers, boards)

	if expected != actual {
		t.Error("Expected ", expected, " got ", actual)
	}
}

func TestPart2(t *testing.T) {
	expected := 1924
	actual := Part2(numbers, boards)

	if expected != actual {
		t.Error("Expected ", expected, " got ", actual)
	}
}

func TestCheckColumns(t *testing.T) {
	// test third column of first board
	called := []int{17, 23, 14, 3, 20}
	board := boards[0]

	for _, num := range called {
		board.Update(num)
	}

	if !board.CheckColumns() {
		t.Error("No bingo found!")
	}
}

func TestCheckRows(t *testing.T) {
	// test third row of first board
	called := []int{21, 9, 14, 16, 7}
	board := boards[0]

	for _, num := range called {
		board.Update(num)
	}

	if !board.CheckRows() {
		t.Error("No bingo found!")
	}
}
