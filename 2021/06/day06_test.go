package main

import (
	"log"
	"testing"
)

var input = []int{3, 4, 3, 1, 2}

func TestPart1(T *testing.T) {
	expected := 26
	days := 26
	actual := Part1(input, days)

	if actual != expected {
		log.Fatalln("FAIL! Expected", expected, "got", actual)
	}
}
