package main

import (
	"log"
	"testing"
)

var numbers = []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}

var boards = []Board{}

func TestPart1(T *testing.T) {
	expected := 4512
	actual := Part1(numbers, boards)

	if expected != actual {
		log.Fatalln("FAIL! Expected ", expected, " got ", actual)
	}
}
