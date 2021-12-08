package main

import (
	"log"
	"testing"
)

var input = []int{3, 4, 3, 1, 2}

func TestPart1Short(T *testing.T) {
	expected := 26
	days := 18
	actual := Part1(input, days)

	if actual != expected {
		log.Fatalln("FAIL! Expected", expected, "got", actual)
	}
}

func TestPart1Long(T *testing.T) {
	expected := 5934
	days := 80
	actual := Part1(input, days)

	if actual != expected {
		log.Fatalln("FAIL! Expected", expected, "got", actual)
	}
}
