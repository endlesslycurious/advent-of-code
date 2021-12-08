package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := []int{3, 4, 3, 1, 2}
	expected := 26
	days := 18

	actual := Part1(input, days)

	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestPart2(t *testing.T) {
	input := []int{3, 4, 3, 1, 2}
	expected := 26984457539
	days := 256

	actual := Part2(input, days)

	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}
