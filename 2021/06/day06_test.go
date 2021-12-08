package main

import (
	"log"
	"testing"
)

var input = []int{3, 4, 3, 1, 2}

func TestPart1(T *testing.T) {
	expected := 26
	days := 18

	actual := Part1(input, days)

	if actual != expected {
		log.Fatalln("FAIL! Expected", expected, "got", actual)
	}
}

func TestPart2(T *testing.T) {
	expected := 26984457539
	days := 256

	actual := Part2(input, days)

	if actual != expected {
		log.Fatalln("FAIL! Expected", expected, "got", actual)
	}
}
