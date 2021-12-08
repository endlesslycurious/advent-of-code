package main

import (
	"log"
	"testing"
)

var input = []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}

func TestPart1(T *testing.T) {
	expected := 37

	actual := Part1(input)

	if actual != expected {
		log.Fatalln("FAIL! Expected", expected, "got", actual)
	}
}
